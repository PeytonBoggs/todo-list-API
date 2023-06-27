package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Returns health, currently hardcoded to "OK"
func getHealth(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "OK")
}

// Returns all tasks
func getTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tasks)
}

// Returns tasks with specified ID
func getTaskByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return
	}

	for _, a := range tasks {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found:"})
}

// Adds task with specified params to the end of tasks
func postTasks(c *gin.Context) {
	var newTask Task

	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	tasks = append(tasks, newTask)
	c.IndentedJSON(http.StatusCreated, newTask)
}

// Adds task with specified params at the task's ID (or end if ID does not exist)
func putTasks(c *gin.Context) {
	var newTask Task
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return
	}

	if err := c.BindJSON(&newTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Error": "could not bind JSON"})
	}

	for i := 0; i < len(tasks); i++ {
		if tasks[i].ID == id {
			tasks[i] = newTask
			c.IndentedJSON(http.StatusOK, newTask)
			return
		}
	}
	tasks = append(tasks, newTask)
	c.IndentedJSON(http.StatusCreated, newTask)
}

// Deletes task with specifed ID
func deleteTask(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return
	}

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
