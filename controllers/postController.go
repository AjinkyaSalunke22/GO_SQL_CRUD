package controllers

import (
	"github.com/AjinkyaSalunke22/GO_SQL_CRUD.git/initializers"
	"github.com/AjinkyaSalunke22/GO_SQL_CRUD.git/models"
	"github.com/gin-gonic/gin"
)

func CreateEvent(c *gin.Context) {
	var eventBody struct{
		Name string
		Place string
	}

	c.Bind(&eventBody)

	event := models.Event{Name: eventBody.Name, Place: eventBody.Place}

	result := initializers.DB.Create(&event)

	if result.Error != nil{
		c.Status(400)
		c.JSON(400, gin.H{
			"message": "Server error in creating event",
		})
		return
	}

	if eventBody.Name == "" || eventBody.Place == "" {
		c.Status(400)
		c.JSON(400, gin.H{
			"message": "Event fields cannot be empty",
		})
		return
	}

	c.JSON(200, gin.H{
		"event" : event,
		"message": "Event creation successfull",
	})
}

func GetEvents(c *gin.Context) {
	
	var events []models.Event
	initializers.DB.Find(&events)

	if events == nil{
		c.Status(404)
		c.JSON(404, gin.H{
			"message": "No event is registered at this time",
		})
		return
	}

	c.JSON(200, gin.H{
		"event" : events,
		"message": "Events fetched successfully",
	})

}

func GetEvent(c *gin.Context) {
	
	id := c.Param("id")

	var event []models.Event
	initializers.DB.Find(&event, id)

	if len(event) == 0{
		c.Status(404)
		c.JSON(404, gin.H{
			"message": "No event is registered with this ID",
		})
		return
	}

	c.JSON(200, gin.H{
		"event" : event,
		"message": "Event fetched successfully",
	})
}

func UpdateEvent(c *gin.Context){

	id := c.Param("id")

	var event struct{
		Name string
		Place string
	}

	c.Bind(&event)

	var db_event []models.Event
	initializers.DB.First(&db_event, id)

	if len(db_event) == 0{
		c.Status(404)
		c.JSON(404, gin.H{
			"message": "No event is registered with this ID",
		})
		return
	}

	initializers.DB.Model(&db_event).Updates(models.Event{Name: event.Name, Place: event.Place})

	c.JSON(200, gin.H{
		"event" : db_event,
		"message": "Event updated successfully",
	})

}

func DeleteEvent(c *gin.Context){
	id := c.Param("id")

	var db_event []models.Event
	initializers.DB.First(&db_event, id)

	if len(db_event) == 0{
		c.Status(404)
		c.JSON(404, gin.H{
			"message": "No event is registered with this ID",
		})
		return
	}

	initializers.DB.Delete(&models.Event{}, id)

	c.JSON(200, gin.H{
		"event" : "",
		"message": "Event deleted successfully",
	})

}