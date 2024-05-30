package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"oms-fiber/model/system/request"
	"testing"
	"time"
)

var jwtManager *JWT

func setup() {
	jwtManager = &JWT{
		[]byte("mdfk"),
	}
}

func TestMain(m *testing.M) {
	setup()
	m.Run()
}

func TestJWTInvalidToken(t *testing.T) {
	// 无效的 Token
	invalidToken := "invalid_token"

	// 解析  Token
	parsedClaims, err := jwtManager.ParseToken(invalidToken)

	assert.Error(t, err)
	assert.Nil(t, parsedClaims)
}

func TestJWTToken(t *testing.T) {
	// 创建 BaseClaims
	baseClaims := request.BaseClaims{
		Username: "testuser",
	}

	// 创建 Token 的 Claims
	claims := jwtManager.CreateClaims(baseClaims)
	claims.MapClaims["exp"] = time.Now().Add(time.Hour * 8888).Unix()
	// 测试创建 Token 的 Claims
	assert.NotNil(t, claims)
	assert.Equal(t, baseClaims.Username, claims.Username)

	// 创建 Token
	token, err := jwtManager.CreateToken(claims)
	if err == nil {
		fmt.Printf("Expired Token : Bearer %v\n", token)
		fmt.Printf("UserID: %s\n", claims.Username)
		fmt.Printf("ExpiresAt: %v\n", claims.MapClaims["exp"])
	}

	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	// 解析 Token
	parsedClaims, err := jwtManager.ParseToken(token)
	assert.NoError(t, err)
	assert.NotNil(t, parsedClaims)

	// 测试解析后的 Claims
	assert.Equal(t, claims.Username, parsedClaims.Username)
}

func TestJWTExpiredToken(t *testing.T) {
	// 创建过期时间为 1 秒的 Token
	baseClaims := request.BaseClaims{
		Username: "testuser",
	}
	expiredClaims := jwtManager.CreateClaims(baseClaims)
	expiredClaims.MapClaims["exp"] = time.Now().Add(-time.Second).Unix()
	token, err := jwtManager.CreateToken(expiredClaims)

	if err == nil {
		fmt.Printf("Expired Token : Bearer %v\n", token)
		fmt.Printf("UserID: %s\n", expiredClaims.Username)
		fmt.Printf("ExpiresAt: %v\n", expiredClaims.MapClaims["exp"])
	}

	// 解析过期的 Token
	parsedClaims, err := jwtManager.ParseToken(token)

	if err == nil {
		fmt.Printf("parsedClaims :  %v\n", parsedClaims)
	}

	assert.Error(t, err)
	assert.Nil(t, parsedClaims)
}
