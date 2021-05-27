package router

import (
	"api_crud/app/controllers"
	"api_crud/app/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRouter(){
	r := gin.Default()
	api_v1 := r.Group("/v1")
	api_v1.POST("/auth/login" , controllers.Login)
	api_v1.POST("/auth/register" , controllers.Register)

	api_v1.Use(middlewares.AuthHandler())
	{
		api_v1.GET("/books" , controllers.GetBooks)
		api_v1.POST("/books" , controllers.Create)
		api_v1.PUT("/books/:id" , controllers.UpdateBook)
		api_v1.DELETE("/books/:id" , controllers.DeleteBook)
		api_v1.POST("/auth/logout" , controllers.Logout)
		api_v1.GET("/user/profile" , controllers.GetProfile)
	}
	r.Run(":8080")
}
