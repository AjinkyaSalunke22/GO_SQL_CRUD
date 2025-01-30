package main

import (
	"log"

	"github.com/AjinkyaSalunke22/GO_SQL_CRUD.git/initializers"
	"github.com/AjinkyaSalunke22/GO_SQL_CRUD.git/models"
)

func init() {
	initializers.LoadEnvVar()
	initializers.ConnectToDb()
}

func main() {
	initializers.DB.AutoMigrate(&models.Event{})
	log.Println("Model creation successfull")
}