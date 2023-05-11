package interfa

import (
	"context"
	bizresponse "uxuy/src/util/response"

	"uxuy/src/internal/svc"
	"uxuy/src/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserMpcKeyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserMpcKeyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserMpcKeyLogic {
	return &UserMpcKeyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserMpcKeyLogic) UserMpcKey(req *types.UserMpcKeyReq) (resp *types.UserMpcKeyResp, err error) {
	user, err := l.svcCtx.UserModel.FindUser(req.Address)
	if err != nil {
		l.Errorf("Recommend FindUsers err. err: %+v", err)
		return nil, bizresponse.ErrInvalidArgs
	}

	return &types.UserMpcKeyResp{
		MpcKey: user.MpcKey,
	}, nil
}
