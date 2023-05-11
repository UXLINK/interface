package interfa

import (
	"context"
	bizresponse "uxuy/src/util/response"

	"uxuy/src/internal/svc"
	"uxuy/src/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAddressLogic {
	return &GetAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAddressLogic) GetAddress(req *types.AddressReq) (resp *types.AddressResp, err error) {
	user, err := l.svcCtx.UserModel.FindUserByDid(req.Did)
	if err != nil {
		l.Errorf("Recommend FindUsers err. err: %+v", err)
		return nil, bizresponse.ErrInvalidArgs
	}

	return &types.AddressResp{
		Address: user.Address,
	}, nil
}
