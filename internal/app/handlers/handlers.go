package handlers

import (
	"encoding/json"
	"hello/internal/app/models"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var person = make(map[uuid.UUID]models.Task)

func PostTask(c *gin.Context) {
	var task models.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	task.ID = uuid.New()

	person[task.ID] = task

	c.JSON(http.StatusCreated, task.ID)
}
func GetTask(c *gin.Context) {
	statusStr := c.Query("status")
	priorityStr := c.Query("priority")
	c.Header("Cache-Control", "public, max-age=3600")
	if statusStr == "" && priorityStr == "" {
		c.JSON(http.StatusOK, person)
		return
	}

	var tasks []models.Task

	for _, task := range person {
		if statusStr != "" {
			status, err := strconv.ParseBool(statusStr)
			if err != nil || task.Status != status {
				continue
			}
		}
		if priorityStr != "" {
			priority, err := strconv.Atoi(priorityStr)
			if err != nil || task.Priority != priority {
				continue
			}
		}

		tasks = append(tasks, task)
	}

	c.JSON(http.StatusOK, tasks)
}

func PutTask(c *gin.Context) {

	id := c.Param("id")

	parsedID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	var task models.Task

	task, exists := person[parsedID]
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Task not found"})
		return
	}

	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	person[parsedID] = task

	c.JSON(http.StatusCreated, parsedID)

}

func DelTask(c *gin.Context) {

	id := c.Param("id")

	parsedID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	delete(person, parsedID)

}

func SaveTasks(c *gin.Context) {
	data, err := json.MarshalIndent(person, "", "\t")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = os.WriteFile("person.json", data, 0644)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

func LoadTasks(c *gin.Context) {
	data, err := os.ReadFile("person.json")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = json.Unmarshal(data, &person)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}
