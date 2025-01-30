package main

import (
	"github.com/AjinkyaSalunke22/GO_SQL_CRUD.git/controllers"
	"github.com/AjinkyaSalunke22/GO_SQL_CRUD.git/initializers"
	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadEnvVar()
	initializers.ConnectToDb()
}

func main() {
	r := gin.Default()
	r.POST("/event", controllers.CreateEvent)
	r.GET("/events", controllers.GetEvents)
	r.GET("/event/:id", controllers.GetEvent)
	r.PUT("/event/:id", controllers.UpdateEvent)
	r.DELETE("/event/:id", controllers.DeleteEvent)
	r.Run() 
}