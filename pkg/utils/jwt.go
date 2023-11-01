package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("hfbpw")

type Claims struct {
	ID        uint   `json:"id"`
	UserName  string `json:"username"`  // 注意区分：UserName ❌
	Authority int    `json:"authority"` // 权限
	jwt.StandardClaims
}

// GenerateToken 签发用户Token
func GenerateToken(id uint, username string, authority int) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)
	claims := Claims{
		ID:        id,
		UserName:  username,
		Authority: authority,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "x",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // ❌ ES256
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// ParseToken 解析用户token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

type EmailClaims struct {
	UserID        uint   `json:"user_id"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	OperationType uint   `json:"operation_type"`
	jwt.StandardClaims
}

// GenerateEmailToken 签发 Email Token
func GenerateEmailToken(userId, Operation uint, email, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)
	claims := EmailClaims{
		UserID:        userId,
		Email:         email,
		Password:      password,
		OperationType: Operation,

		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "x",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // ❌ ES256
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// ParseEmailToken 解析 Email Token
func ParseEmailToken(token string) (*EmailClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &EmailClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*EmailClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
