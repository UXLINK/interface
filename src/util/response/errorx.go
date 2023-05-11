package bizresponse

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CodeError struct {
	Success  bool   `json:"success"`
	Code     int    `json:"code"`
	Msg      string `json:"msg"`
	EnMsg    string `json:"enMsg"`
	Language string `json:"language"`
}

type CodeErrorResponse struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
}

func NewCodeError(code int, msg string) *CodeError {
	return &CodeError{
		Success: false,
		Code:    code,
		Msg:     msg,
	}
}

// 初始化CodeError并塞进pool中
func NewCodeErrorWithInit(code int, msg string) *CodeError {
	cr := NewCodeError(code, msg)
	pool[cr.generateKey()] = cr
	return cr
}

// 为传递grpcerror的code和msg,实现GRPCStatus
func (e *CodeError) GRPCStatus() *status.Status {
	return status.New(codes.Code(e.Code), e.Msg)
}

// 实现error的msg方法
func (e *CodeError) Error() string {
	return e.Msg
}

// 生成Pool的key名
func (e *CodeError) generateKey() string {
	return fmt.Sprintf("%d%v", e.Code, e.Msg)
}

func (e *CodeError) Data() *CodeErrorResponse {
	return &CodeErrorResponse{
		Code: e.Code,
		Msg:  e.Msg,
	}
}

// WithMessage 用于从已定义的CodeError重载生成包含指定msg的BizResponse
func (e *CodeError) WithMessage(msg string) *CodeError {
	return NewCodeError(e.Code, msg)
}

// 比较bizresponse，注意普通的error会被转成ErrInternalFailed后进行判断
func (e *CodeError) EqualsTo(err error) bool {
	if err == nil {
		return false
	}
	c := FromError(err)
	if c == nil {
		return false
	}
	return (e == c) || (e.Code == c.Code)
}

// 通过error转成CodeError
func FromError(err error) (resp *CodeError) {
	if err == nil {
		return nil
	}

	switch err {
	case context.Canceled:
		return ErrClientCancel
	case context.DeadlineExceeded:
		return ErrDeadlineExceed
	default:

	}

	s, ok := status.FromError(err)
	if !ok {
		if biz, ok := err.(*CodeError); ok {
			return biz
		}
		return ErrInternalFailed.WithMessage(err.Error())
	}

	if br, ok := pool[fmt.Sprintf("%d%v", s.Code(), s.Message())]; ok {
		return br
	}

	// grpc 通常都是100以内
	if s.Code() < 100 {
		switch s.Code() {
		case codes.Canceled:
			return ErrClientCancel
		}

		return ErrInternalFailed.WithMessage(err.Error())
	}
	return NewCodeError(int(s.Code()), s.Message())
}
