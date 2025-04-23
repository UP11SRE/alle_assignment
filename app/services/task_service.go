package services

import (
	"alle_assignment/app/models"
	"alle_assignment/app/repositories"
	"errors"
)

func IsValidStatus(status string) bool {
	return status == "pending" || status == "completed"
}

func CreateTask(task *models.Task) error {

	if task.Status == "" {
		task.Status = "pending"
	}
   
	if !IsValidStatus(task.Status) {
		return errors.New("task status must be either 'pending' or 'completed'")
	}
	


	return repositories.CreateTask(task)
}

func GetAllTasks(page int) ([]models.Task, int, error) {
	return repositories.GetAllTasks(page)
}

func UpdateTask(id int, updates map[string]interface{}) error {

	if !IsValidStatus(updates["status"].(string)) {
		return errors.New("task status must be either 'pending' or 'completed'")
	}
	
	if _, exists := updates["user_name"]; exists {
		return errors.New("user_name can't be updated")
	}

	allowedFields := map[string]bool{
		"title":       true,
		"description": true,
		"status":      true,
	}

	filteredUpdates := make(map[string]interface{})
	for key, value := range updates {
		if allowedFields[key] {
			filteredUpdates[key] = value
		}
	}

	if len(filteredUpdates) == 0 {
		return errors.New("no valid fields provided for update")
	}

	return repositories.UpdateTask(id, filteredUpdates)
}

func DeleteTask(id int) error {
	return repositories.DeleteTask(id)
}
func FilterTasks( status *string, page int) ([]models.Task, int, error) {
	return repositories.FilterTasks( status, page)
}



func GetTaskByID(id string) (*models.Task, error) {
    return repositories.GetTaskByID(id)
}
