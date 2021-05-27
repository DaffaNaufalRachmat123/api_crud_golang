package middlewares

import (
	"time"
	"api_crud/config"
	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(username string , unique_key string) (string , interface{} , error){
	k := []byte(config.LoadEnv("SIGNED_KEY"))
	tokens := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	claims["unique_key"] = unique_key
	tokens.Claims = claims
	tokenString , err := tokens.SignedString(k)
	return tokenString , claims["exp"] , err
}

func ValidateToken(t string , k string)(*jwt.Token , error){
	token , err := jwt.Parse(t , func(token *jwt.Token)(interface{} , error){
		return []byte(k) , nil
	})
	return token , err
}