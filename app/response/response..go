package response
import (
	"github.com/gin-gonic/gin"
)

func NotFound(c *gin.Context , item string){
	c.AbortWithStatusJSON(404 , gin.H {
		"status_code" : 404,
		"status_msg" : item + " not found",
	})
}

func InternalError(c *gin.Context){
	c.AbortWithStatusJSON(500 , gin.H {
		"status_code" : 500,
		"status_msg" : "Internal Server Error",
	})
}

func Unauthorized(c *gin.Context){
	c.AbortWithStatusJSON(401 , gin.H {
		"status_code" : 401,
		"status_msg" : "Unauthorized",
	})
}

func Exists(c *gin.Context , item string){
	c.AbortWithStatusJSON(409 , gin.H {
		"status_code" : 409,
		"status_msg" : item + " already exists",
	})
}

func BadCredentials(c *gin.Context){
	c.AbortWithStatusJSON(401 , gin.H {
		"status_code" : 401,
		"status_msg" : "bad credentials",
	})
}

func InvalidToken(c *gin.Context){
	c.AbortWithStatusJSON(401 , gin.H {
		"status_code" : 401,
		"status_msg" : "Invalid Token",
	})
}

func BadValidator(c *gin.Context , err interface{}){
	c.AbortWithStatusJSON(422 , gin.H {
		"status_code" : 422,
		"status_msg" : err,
	})
}

func Send(c *gin.Context , statusCode int , msg string , data interface{}){
	c.JSON(statusCode ,gin.H {
		"status_code" : statusCode,
		"status_msg" : msg,
		"data" : data,
	})
}