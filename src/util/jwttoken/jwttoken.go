package jwttoken

import (
	"time"

	"uxuy/src/util/ctxdata"
	bizresponse "uxuy/src/util/response"

	"github.com/golang-jwt/jwt/v4"
)

type JwtClient struct {
	accessSecret  string
	accessExpire  int64
	refreshExpire int64
}

type JwtToken struct {
	AccessToken  string
	RefreshToken string
	AccessExpire int64
	RefreshAfter int64
}

type UserClaims struct {
	jwt.RegisteredClaims
	JwtUserId int64 `json:"jwtUserId"`
}

func NewJwtClient(accessSecret string, accessExpire int64, refreshExpire int64) *JwtClient {
	return &JwtClient{
		accessSecret:  accessSecret,
		accessExpire:  accessExpire,
		refreshExpire: refreshExpire,
	}
}

func (j *JwtClient) GenerateToken(dappId string) (*JwtToken, error) {
	now := time.Now().Unix()
	accessToken, err := j.getJwtToken(now, dappId)
	if err != nil {
		return nil, bizresponse.ErrInvalidArgs
	}

	refreshToken, err := j.getJwtRefreshToken(now, dappId)
	if err != nil {
		return nil, bizresponse.ErrInvalidArgs
	}

	return &JwtToken{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		AccessExpire: now + j.accessExpire,
		RefreshAfter: now + j.accessExpire/2, // 有效时间过一半，希望刷新token
	}, nil
}

func (j *JwtClient) getJwtToken(iat int64, dappId string) (string, error) {

	claims := make(jwt.MapClaims)
	claims["exp"] = iat + j.accessExpire
	claims["iat"] = iat
	claims[ctxdata.CtxKeyJwtDappId] = dappId
	// 你还可以埋更多的信息在这里 比如用户名字啥的 回头再ctxData中写个获取方法就行
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(j.accessSecret))
}

func (j *JwtClient) getJwtRefreshToken(iat int64, dappId string) (string, error) {

	claims := make(jwt.MapClaims)
	claims["exp"] = iat + j.refreshExpire
	claims["iat"] = iat
	claims[ctxdata.CtxKeyJwtDappId] = dappId
	// 你还可以埋更多的信息在这里 比如用户名字啥的 回头再ctxData中写个获取方法就行
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(j.accessSecret))
}

func (j *JwtClient) ParseToken(tokenString string) (int64, error) {
	return ParseWithClaims(j.accessSecret, tokenString)
}

func ParseWithClaims(accessSecret, tokenString string) (int64, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&UserClaims{},
		func(token *jwt.Token) (any, error) {
			return []byte(accessSecret), nil
		},
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}),
	)
	if err != nil {
		return 0, err
	}
	if !token.Valid {
		return 0, bizresponse.ErrInvalidArgs
	}

	userInfo, ok := token.Claims.(*UserClaims)
	if !ok {
		return 0, bizresponse.ErrInvalidArgs
	}

	return userInfo.JwtUserId, nil
}
