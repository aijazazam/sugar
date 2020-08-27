package handler

import (
	"fmt"
	"time"
	"errors"
	"strings"
	"net/http"
	"github.com/dgrijalva/jwt-go"
)

const (
  SECRET_KEY = "my_secret_key"
)

func GetJwt(r *http.Request) string {

	bearToken := r.Header.Get("Authorization")

	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")

	if len(strArr) == 2 {
		return strArr[1]
	}

	return ""
}

func CreateToken( tUserName string, isAdmin bool ) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
    "username": tUserName,
    "loginAt" : time.Now().Unix(),
    "expireAt": time.Now().Add(time.Hour * time.Duration(1)).Unix(),
    "isAdmin" : isAdmin,
  })

	tokenString, err := token.SignedString( []byte(SECRET_KEY) )

	return tokenString, err
}

func DecodeToken(r *http.Request) (username string, isAdmin bool, err error) {

	tokenString := GetJwt(r)
	if len(tokenString) == 0 {
		err = errors.New("Empty jwt Token: Only Loged in users allowed to rate/comment movies and admin allowed to add new movies")
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return "", false, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username	= claims["username"].(string)
		isAdmin		= claims["isAdmin"].(bool)
	}

	return
}

