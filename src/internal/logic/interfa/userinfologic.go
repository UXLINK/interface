package interfa

import (
	"context"
	bizresponse "uxuy/src/util/response"

	"uxuy/src/internal/svc"
	"uxuy/src/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	user, err := l.svcCtx.UserModel.FindUser(req.Address)
	if err != nil {
		l.Errorf("Recommend FindUsers err. err: %+v", err)
		return nil, bizresponse.ErrInvalidArgs
	}

	return &types.UserInfoResp{
		UxuyId:    user.UxuyId,
		Name:      user.Name,
		Avatar:    user.Avatar,
		Did:       user.Did,
		CreatedAt: user.CreatedAt,
	}, nil
}
