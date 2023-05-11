package interfa

import (
	"context"
	bizresponse "uxuy/src/util/response"

	"uxuy/src/internal/svc"
	"uxuy/src/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RecommendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRecommendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecommendLogic {
	return &RecommendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RecommendLogic) Recommend(req *types.RecommendReq) (resp *types.RecommendListResp, err error) {
	if req.Cursor != "" {
		return &types.RecommendListResp{
			Cursor: "a8728c5910a04c94b70e18694d72cbb0",
			List:   make([]types.RecommendUser, 0),
		}, nil
	}

	users, err := l.svcCtx.UserModel.FindUsers()
	if err != nil {
		l.Errorf("Recommend FindUsers err. err: %+v", err)
		return nil, bizresponse.ErrInvalidArgs
	}
	list := make([]types.RecommendUser, len(users))
	for i, user := range users {
		recommendUser := types.RecommendUser{
			UxuyId:          user.UxuyId,
			Name:            user.Name,
			Avatar:          user.Avatar,
			RecommendReason: "tags: Seaman NFT",
		}

		list[i] = recommendUser
	}

	return &types.RecommendListResp{
		Cursor: "a8728c5910a04c94b70e18694d72cbb0",
		List:   list,
	}, nil
}
