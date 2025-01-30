package main

import (
	"github.com/AjinkyaSalunke22/GO_SQL_CRUD.git/initializers"
	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadEnvVar()
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}