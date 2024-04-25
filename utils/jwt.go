package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"oms-fiber/global"
	"oms-fiber/model/system/request"
	"time"
)

type JWT struct {
	SigningKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.Config.JWT.SigningKey),
	}
}

func (j *JWT) CreateClaims(baseClaims request.BaseClaims) request.CustomClaims {
	//ep, _ := ParseDuration(global.Config.JWT.ExpiresTime)
	ep := time.Hour
	claims := request.CustomClaims{
		BaseClaims: baseClaims,
		MapClaims: jwt.MapClaims{
			"exp": time.Now().Add(ep).Unix(),           //Expiration Time：过期时间。过期时间是 UNIX 时间戳(自 1970 年 1 月 1 日以来的秒数)。
			"nbf": time.Now().Add(-time.Minute).Unix(), //Not Before：生效时间。在此时间之前，JWT 将被认为是无效的。
			"iat": time.Now().Add(-time.Minute).Unix(), //Issued At：签发时间。 JWT何时被创建，通常是 UNIX 时间戳。
			"aud": "",                                  //Audience：接收者。JWT 的预期的接收者，可以是单个字符串或字符串数组
			"iss": global.Config.JWT.Issuer,            //Issuer：签发者。
			"sub": "",                                  //Subject：它指示了 JWT 所代表的主体，任意搞。
		},
	}
	return claims
}

// CreateToken 创建一个token
func (j *JWT) CreateToken(claims request.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// CreateTokenByOldToken 旧token 换新token 使用归并回源避免并发问题
func (j *JWT) CreateTokenByOldToken(oldToken string, claims request.CustomClaims) (string, error) {
	v, err, _ := global.ConcurrencyControl.Do("JWT:"+oldToken, func() (interface{}, error) {
		return j.CreateToken(claims)
	})
	return v.(string), err
}

// 解析 token
func (j *JWT) ParseToken(tokenString string) (*request.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &request.CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return nil, err
	}
	if token != nil {
		if claims, ok := token.Claims.(*request.CustomClaims); ok && token.Valid {
			return claims, nil
		}
	}
	return nil, errors.New("token is invalid")
}
