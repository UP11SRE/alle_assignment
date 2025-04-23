package routes

import (
	"alle_assignment/app/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	
	
	router.GET("/", controllers.Default)
	router.POST("/tasks", controllers.CreateTask)
	router.GET("/tasks", controllers.GetAllTasks)
	router.PATCH("/tasks/:id", controllers.UpdateTask)
	router.DELETE("/tasks/:id", controllers.DeleteTask)
	router.GET("/tasks/filter", controllers.FilterTasks)
	router.GET("/tasks/:id", controllers.GetTaskByID)

	
}
