package controllers

import (
	"alle_assignment/app/models"
	"alle_assignment/app/services"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Default(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to the Task Management API!"})
}

func CreateTask(c *gin.Context) {
	var task models.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.CreateTask(&task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	c.JSON(http.StatusCreated, task)
}

func GetAllTasks(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	tasks, total, err := services.GetAllTasks(page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get tasks"})
		return
	}

	pageSize := 10
	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))
	hasNext := page < totalPages

	c.JSON(http.StatusOK, gin.H{
		"tasks": tasks,
		"meta": gin.H{
			"current_page": page,
			"page_size": pageSize,
			"total_items": total,
			"total_pages": totalPages,
			"has_next": hasNext,
		},
	})
}




func UpdateTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.UpdateTask(id, updates); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
}

func DeleteTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	if err := services.DeleteTask(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}

func FilterTasks(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	queryParams := c.Request.URL.Query()
	for key := range queryParams {
		if key != "status" && key != "page" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Invalid query parameters"})
			return
		}
	}

	status := c.Query("status")
	if status == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "status parameter is required"})
		return
	}

	tasks, total, err := services.FilterTasks(&status, page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to filter tasks"})
		return
	}

	pageSize := 10
	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))
	hasNext := page < totalPages

	c.JSON(http.StatusOK, gin.H{
		"tasks": tasks,
		"meta": gin.H{
			"current_page": page,
			"page_size":    pageSize,
			"total_items":  total,
			"total_pages":  totalPages,
			"has_next":     hasNext,
		},
	})
}





func GetTaskByID(c *gin.Context) {
    id := c.Param("id") 
    task, err := services.GetTaskByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
        return
    }
    c.JSON(http.StatusOK, task)
}
