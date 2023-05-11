package bizresponse

type CodeSuccess struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
	Code    int    `json:"code"`
	Data    any    `json:"data"`
}

func NewSuccessResp(data any) *CodeSuccess {
	ret := &CodeSuccess{
		Success: true,
		Msg:     "ok",
		Code:    200,
		Data:    data,
	}
	if data == nil {
		ret.Data = struct{}{}
	}

	return ret
}

// WithData 用于从已定义的ErrResponse重载生成包含指定返回的数据
func (b *CodeSuccess) WithData(data any) *CodeSuccess {
	resp := NewSuccessResp(data)
	return resp
}
