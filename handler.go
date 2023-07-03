package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// TODO: Implement getHealth
func getHealth(c *gin.Context) {
}

// getTasks godoc
// @Summary getTasks
// @Description Gets all tasks in database
// @Tags root
// @Accept */*
// @Produce json
// @Router /tasks [get]
func getTasks(c *gin.Context) {
	taskList, err := getTasks_sql()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	}

	c.IndentedJSON(http.StatusOK, taskList)
}

// getTaskByID godoc
// @Summary getTaskByID
// @Description Gets all tasks with specified ID
// @Tags root
// @Param id path int true "ID to get"
// @Accept */*
// @Produce json
// @Router /tasks/{id} [get]
func getTaskByID(c *gin.Context) {
	searchedID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
	}

	foundTask, err := getTaskByID_sql(searchedID)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err)
		return
	}

	c.IndentedJSON(http.StatusOK, foundTask)
}

// TODO: add func getTasksByComplete

// postTask godoc
// @Summary postTasks
// @Description Adds new task at the end of database
// @Tags root
// @RequestBody required
// @Param Task body shortTask true "Task to add"
// @Accept */*
// @Produce json
// @Router /tasks [POST]
func postTask(c *gin.Context) {
	var newTask shortTask

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
