package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// TODO: Implement getHealth
func getHealth(c *gin.Context) {
}

// TODO: Implement getTasks
func getTasks(c *gin.Context) {
}

// TODO: Implement getTasksByID
func getTaskByID(c *gin.Context) {
}

// TODO: add func getTasksByComplete

// postTask godoc
// @Summary postTasks
// @Description Adds new task at the end of database
// @Tags root
// @RequestBody required
// @Param task body Task true "The task to add"
// @Accept */*
// @Produce json
// @Router /tasks [POST]
func postTask(c *gin.Context) {
	var newTask Task

	if err := c.BindJSON(&newTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	newID, err := postTask_sql(newTask)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	result := "Task created: " + newTask.Title + " at ID " + strconv.Itoa(int(newID))
	c.IndentedJSON(http.StatusCreated, result)
}

// TODO: Implement putTasks
func putTasks(c *gin.Context) {
}

// TODO: Implement deleteTask
func deleteTask(c *gin.Context) {
}
