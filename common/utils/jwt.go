package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// CustomPayload 自定义载荷继承原有接口并附带自己的字段
type CustomPayload struct {
	UserId     uint64
	GrantScope string // role
	jwt.RegisteredClaims
}

// GenerateToken 生成Token uid 用户id subject 签发对象  secret 加盐
func GenerateToken(uid uint64, subject string, secret string) (string, error) {
	claim := CustomPayload{
		UserId:     uid,
		GrantScope: subject,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Auth_Server",                                   //签发者
			Subject:   subject,                                         //签发对象
			Audience:  jwt.ClaimStrings{"PC"},                          //签发受众
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),   //过期时间
			NotBefore: jwt.NewNumericDate(time.Now().Add(time.Second)), //最早使用时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                  //签发时间
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString([]byte(secret))
	return token, err
}

func ParseToken(token string, secret string) (*CustomPayload, error) {
	// 解析token
	parseToken, err := jwt.ParseWithClaims(token, &CustomPayload{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(secret), nil
	}, jwt.WithLeeway(5*time.Second))
	if err != nil {
		return nil, err
	}
	if claims, ok := parseToken.Claims.(*CustomPayload); ok && parseToken.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

func GetJwtToken(secretKey string, iat, seconds int64, uid int64, role string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["uid"] = uid
	claims["role"] = role
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

func ParseJwtToken(secertKey string, token string) (jwt.MapClaims, error) {
	parseToken, err := jwt.Parse(token, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(secertKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := parseToken.Claims.(jwt.MapClaims); ok && parseToken.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
