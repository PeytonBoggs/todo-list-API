package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type task struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Complete bool   `json:"complete"`
}

var tasks = []task{
	{ID: "1", Name: "Wake up", Complete: false},
	{ID: "2", Name: "Go to work", Complete: false},
	{ID: "3", Name: "Make dinner", Complete: false},
}

func main() {
	router := gin.Default()
	router.GET("/health", getHealth)
	router.GET("/tasks", getTasks)
	router.POST("/tasks", postTasks)

	router.Run("localhost:8080")
}

func getHealth(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "OK")
}

func getTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tasks)
}

func postTasks(c *gin.Context) {
	var newTask task

	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	tasks = append(tasks, newTask)
	c.IndentedJSON(http.StatusCreated, newTask)
}
