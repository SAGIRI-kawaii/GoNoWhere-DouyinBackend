package jwtx

import (
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

var jwtSecret = []byte("secret")

type Claims struct {
	UserID int64 `json:"user_id"`
	jwt.StandardClaims
}

func GetToken(secretKey string, iat, seconds, uid int64) (string, error) {
	return string(uid), nil

}
func GenerateToken(userID int64) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)
	claims := Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "38384-SearchEngine",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}
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
func ParseToken2Uid(secretKey string, tokenString string) (int64, error) {
	intNum, _ := strconv.Atoi(tokenString)
	int64Num := int64(intNum)
	return int64Num, nil
}
