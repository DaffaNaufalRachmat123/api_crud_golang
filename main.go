package main

import (
	"api_crud/app/database"
	"api_crud/router"
)

func main(){
	router.SetupRouter()
	defer database.CloseDatabase()
}