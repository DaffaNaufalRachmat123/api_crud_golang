package controllers

import (
	"api_crud/app/database"
	"api_crud/app/models"
	"api_crud/app/response"
	"api_crud/config"
	"api_crud/validator"
	"github.com/gin-gonic/gin"
	"github.com/ulule/deepcopier"
)

func Create(c *gin.Context){
	var req models.BooksValidate
	err := validator.Validate(c , &req)
	if err != nil {
		response.BadValidator(c , err.Error())
		return
	}
	database := database.GetDatabase()
	var books models.Books
	c.BindJSON(&req)
	deepcopier.Copy(req).To(&books)
	err = database.Where("title = ?" , req.Title).Find(&books).Error
	if err == nil {
		response.Exists(c , "This book")
		return
	}
	database.Create(&books)
	response.Send(c , 200 , "Successfully create book" , books)
}

func GetBooks(c *gin.Context){
	var books []models.Books
	database := database.GetDatabase()
	p := config.Page(c)
	err := database.Limit(config.PaginationLimit).Offset(p * config.PaginationLimit).Find(&books).Error
	if err != nil {
		response.NotFound(c , "This books")
		return
	} else {
		response.Send(c , 200 , "Successfully get books" , books)
		return
	}
}

func DeleteBook(c *gin.Context){
	id := c.Params.ByName("id")
	var books models.Books
	database := database.GetDatabase()

	err := database.Where("id = ?" , id).Find(&books).Error
	if err != nil {
		response.NotFound(c , "This books")
		return
	}
	database.Delete(&books)
	response.Send(c , 200 , "Succesfully delete book" , nil)
}

func UpdateBook(c *gin.Context){
	var req models.BooksValidate
	id := c.Params.ByName("id")
	err := validator.Validate(c , &req)
	if err != nil {
		response.BadValidator(c , err.Error())
		return
	}
	database := database.GetDatabase()
	var books models.Books
	err = database.Where("id = ?" , id).Find(&books).Error
	if err != nil {
		response.NotFound(c , "This books")
		return
	}
	c.BindJSON(&req)
	deepcopier.Copy(req).To(&books)
	database.Save(&books)
	response.Send(c , 200 , "Successfully update books" , books)
}