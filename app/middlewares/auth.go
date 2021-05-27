package middlewares

import (
	"fmt"
	"strings"
	"api_crud/config"
	"api_crud/app/database"
	"api_crud/app/models"
	"api_crud/app/response"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
)

var username string


func AuthHandler() gin.HandlerFunc {
	return func(c *gin.Context){
		token := c.Request.Header.Get("Authorization")
		b := "Bearer "
		if !strings.Contains(token , b) {
			response.Unauthorized(c)
			c.Abort()
			return
		}
		t := strings.Split(token , b)
		if len(t) < 2 {
			response.InvalidToken(c)
			c.Abort()
			return
		}
		valid , err := ValidateToken(t[1] , config.LoadEnv("SIGNED_KEY"))
		if err != nil {
			response.InvalidToken(c)
			c.Abort()
			return
		}
		db := database.GetDatabase()
		var user models.User
		username = fmt.Sprintf("%v" , valid.Claims.(jwt.MapClaims)["username"])
		err = db.Where("username = ?" , username).Find(&user).Error
		if err != nil {
			response.Unauthorized(c)
			return
		}
		unique_key := valid.Claims.(jwt.MapClaims)["unique_key"]
		if user.Unique_Key != unique_key {
			response.Unauthorized(c)
			return
		}
		c.Set("username" , username)
		c.Next()
	}
}

func GetUsername() string {
	return username
}

func HashPassword(password string)(string , error){
	bytes , err := bcrypt.GenerateFromPassword([]byte(password) , 14)
	return string(bytes) , err
}

func CheckPasswordHash(password , hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash) , []byte(password))
	return err == nil
}