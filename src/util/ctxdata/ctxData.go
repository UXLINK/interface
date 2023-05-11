package ctxdata

import (
	"context"
)

func GetDappIdFromCtx(ctx context.Context) string {
	var dappId string
	if jsonUid, ok := ctx.Value(CtxKeyJwtDappId).(string); ok {
		dappId = jsonUid
	}

	return dappId
}

func SetDappIToCtx(ctx context.Context, dappId string) context.Context {
	ctx = context.WithValue(ctx, CtxKeyJwtDappId, dappId)
	return ctx
}
