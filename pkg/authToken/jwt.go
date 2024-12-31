package authToken

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWTSecret 是用于签名和验证 JWT 的密钥
var JWTSecret string

func init() {
	// 生成一个 32 字节的随机密钥
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}
	// 将随机字节转换为 base64 编码的字符串
	JWTSecret = base64.URLEncoding.EncodeToString(key)
}

// Claims 是自定义的 JWT 声明结构体，它嵌入了标准的 Claims 结构体
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateToken 生成 JWT 令牌
func GenerateToken(username string) (string, error) {
	// 过期时间设置为 1 小时后
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Subject:   "user",
		},
	}

	// 创建 JWT 令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用密钥签名令牌
	tokenString, err := token.SignedString(JWTSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ValidateToken 验证 JWT 令牌
func ValidateToken(tokenString string) (*Claims, error) {
	// 解析 JWT 令牌
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// 确保使用的签名算法是 HS256
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Method.Alg())
		}
		return JWTSecret, nil
	})
	if err != nil {
		return nil, err
	}

	// 验证令牌并提取声明
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// RefreshToken 刷新 JWT 令牌
func RefreshToken(tokenString string) (string, error) {
	// 首先验证原令牌
	claims, err := ValidateToken(tokenString)
	if err != nil {
		return "", err
	}

	// 创建新的过期时间，这里设置为 1 小时后
	newExpirationTime := time.Now().Add(1 * time.Hour)
	claims.ExpiresAt = jwt.NewNumericDate(newExpirationTime)

	// 生成新的令牌
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	newTokenString, err := newToken.SignedString(JWTSecret)
	if err != nil {
		return "", err
	}
	return newTokenString, nil
}
