package interfa

import (
	"context"
	bizresponse "uxuy/src/util/response"

	"uxuy/src/internal/svc"
	"uxuy/src/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNameLogic {
	return &GetNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetNameLogic) GetName(req *types.NameReq) (resp *types.NameResp, err error) {
	user, err := l.svcCtx.UserModel.FindUserByAddress(req.Address)
	if err != nil {
		l.Errorf("Recommend FindUsers err. err: %+v", err)
		return nil, bizresponse.ErrInvalidArgs
	}

	return &types.NameResp{
		Did: user.Did,
	}, nil
}
