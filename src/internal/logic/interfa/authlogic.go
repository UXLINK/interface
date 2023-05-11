package interfa

import (
	"context"
	"uxuy/src/util/ctxdata"
	bizresponse "uxuy/src/util/response"

	"uxuy/src/internal/svc"
	"uxuy/src/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthLogic {
	return &AuthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuthLogic) Auth() (resp *types.AuthResp, err error) {
	dappId := ctxdata.GetDappIdFromCtx(l.ctx)

	if dappId == "" {
		l.Errorf("Auth failed. dappId: %v", dappId)
		return nil, bizresponse.ErrClientCancel
	}

	// jwtId
	tokenResp, err := l.svcCtx.JwtClient.GenerateToken(dappId)
	if err != nil {
		l.Errorf("Auth GenerateToken failed. %v", err)
		return nil, bizresponse.ErrInternalFailed
	}

	return &types.AuthResp{
		Token: tokenResp.AccessToken,
	}, nil
}
