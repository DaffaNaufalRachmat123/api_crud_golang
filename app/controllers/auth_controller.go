package controllers

import (
	"api_crud/app/database"
	"api_crud/app/middlewares"
	"api_crud/app/models"
	"api_crud/app/response"
	"api_crud/validator"
	"github.com/gin-gonic/gin"
)


func Login(c *gin.Context){
	var req models.LoginValidate
	err := validator.Validate(c , &req)
	if err != nil {
		response.BadValidator(c , err.Error())
		return
	}
	var user models.User
	database := database.GetDatabase()
	c.BindJSON(&req)
	err = database.Where("username = ?" , req.Username).Find(&user).Error
	if err != nil {
		response.NotFound(c , "This user")
		return
	}
	valid := middlewares.CheckPasswordHash(req.Password , user.Password)
	if valid == false {
		response.BadCredentials(c)
		return
	}
	user.Unique_Key = middlewares.RandomString(10)
	database.Save(&user)
	token , exp , err := middlewares.GenerateToken(user.Username , user.Unique_Key)
	if err != nil {
		response.InternalError(c)
		return
	}
	response.Send(c , 200 , "Login Successfully" , gin.H {
		"status" : 200,
		"server_message" : "Success",
		"token" : "Bearer " + token,
		"exp" : exp,
		"data" : gin.H {
			"id" : user.ID,
			"email" : user.Email,
		},
	})
}

func Register(c *gin.Context){
	var req models.RegisterValidate
	err := validator.Validate(c , &req)
	if err != nil {
		response.BadValidator(c , err.Error())
		return
	}
	database := database.GetDatabase()
	c.BindJSON(&req)
	password , _ := middlewares.HashPassword(req.Password)
	unique_key := middlewares.RandomString(10)
	user := models.User {
		Username : req.Username,
		Password : password,
		Email : req.Email,
		Unique_Key: unique_key,
		Profile : models.Profile {
			Name : req.Name,
			Phone : req.Phone,
			City : req.City,
		},
	}
	err = database.Where("username = ?" , req.Username).Find(&user).Error
	if err == nil {
		response.Exists(c , "This users")
		return
	}
	database.Create(&user)
	response.Send(c , 201 , "Successfully registered" , &user)
	return
}

func Logout(c *gin.Context){
	username := middlewares.GetUsername()
	database := database.GetDatabase()
	var user models.User
	err := database.Where("username = ?" , username).Find(&user).Error
	if err != nil {
		response.NotFound(c , "This user")
		return
	}
	user.Unique_Key = middlewares.RandomString(10)
	database.Save(&user)
	response.Send(c , 200 , "Successfully logout" , nil)
}