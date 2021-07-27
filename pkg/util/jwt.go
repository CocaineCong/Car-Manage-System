package util

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
	UserID uint `json:"user_id"`
	UserName  string `json:"user_name"`
	OpenID  string `json:"open_id"`
	Authority int    `json:"authority"`
	jwt.StandardClaims
}

//GenerateToken 签发用户Token
func GenerateToken(userid uint,username, open_id string, authority int) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)
	claims := Claims{
		UserID : userid,
		UserName:  username,
		OpenID:  open_id,
		Authority: authority,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "KeyValueTeam",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

//ParseToken 验证用户token
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

//EmailClaims
type EmailClaims struct {
	UserID        uint   `json:"user_id"`
	Email         string `json:"email"`
	OpenID        string 	 `json:"open_id"`
	OperationType uint   `json:"operation_type"`
	jwt.StandardClaims
}

//GenerateEmailToken 签发邮箱验证Token
func GenerateEmailToken(userID, Operation uint, email, OpenID string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(15 * time.Minute)
	claims := EmailClaims{
		UserID:        userID,
		Email:         email,
		OpenID:      OpenID,
		OperationType: Operation,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "KeyValueTeam",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

//ParseEmailToken 验证邮箱验证token
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
