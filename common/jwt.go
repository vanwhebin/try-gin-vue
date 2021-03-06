package common

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"time"
	"vanwhebin/try-gin-vue/model"
)

var jwtKey = []byte(viper.GetString("jwt.key"))

// define jwt token struct

type JwtClaim struct {
	UserID uint
	jwt.StandardClaims
}

// release token
func ReleaseToken(user model.User) (string, error) {
	expire := time.Now().Add(7 * 24 * time.Hour)
	claim := JwtClaim{
		user.ID,
		jwt.StandardClaims{
			ExpiresAt: expire.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    viper.GetString("jwt.issuer"),
			Subject:   viper.GetString("jwt.subject"),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	ss, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	fmt.Printf("%v %v", ss, err)
	return ss, nil
}

// parse token get value
func ParseToken(tokenString string) (*jwt.Token, *JwtClaim, error) {
	claims := &JwtClaim{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{},err error) {
		return jwtKey, nil
	})
	return token, claims, err
}
