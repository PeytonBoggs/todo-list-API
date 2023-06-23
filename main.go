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
	router.GET("/tasks/:id", getTaskByID)
	router.POST("/tasks", postTasks)
	router.DELETE("/tasks/:id", deleteTask)

	router.Run("localhost:8080")
}

func getHealth(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "OK")
}

func getTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tasks)
}

func getTaskByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range tasks {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found:"})
}

func postTasks(c *gin.Context) {
	var newTask task

	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	tasks = append(tasks, newTask)
	c.IndentedJSON(http.StatusCreated, newTask)
}

func deleteTask(c *gin.Context) {
	var id string = c.Param("id")

	for i := 0; i < len(tasks); i++ {
		if tasks[i].ID == id {
			firstHalf := tasks[0:i]
			secondHalf := tasks[i+1:]
			tasks = append(firstHalf, secondHalf...)
			c.IndentedJSON(http.StatusOK, tasks)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"Error": "ID not found"})
}
