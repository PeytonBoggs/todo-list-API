package main

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// TODO: Implement getHealth
func getHealth(c *gin.Context) {
}

// getTasks godoc
// @Summary getTasks
// @Description Gets all tasks in database
// @Tags tasks
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
// @Tags tasks
// @Param id path int true "ID to get"
// @Accept */*
// @Produce json
// @Router /tasks/id/{id} [get]
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

// getTasksByComplete godoc
// @Summary getTasksByComplete
// @Description Gets all tasks with specified "complete" value
// @Tags tasks
// @Param complete path boolean true "Complete? true or false""
// @Accept */*
// @Produce json
// @Router /tasks/complete/{complete} [GET]
func getTasksByComplete(c *gin.Context) {
	complete, err := strconv.ParseBool(strings.ToUpper(c.Param("complete")))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Error: value can either be true or false")
		return
	}

	taskList, err := getTasksByComplete_sql(complete)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err)
	}

	c.IndentedJSON(http.StatusOK, taskList)
}

// postTask godoc
// @Summary postTasks
// @Description Adds new task at the end of database
// @Tags tasks
// @RequestBody required
// @Param Task body TaskPayload true "Task to add"
// @Accept */*
// @Produce json
// @Router /tasks [POST]
func postTask(c *gin.Context) {
	var newTask TaskPayload

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

// deleteTasks godoc
// @Summary deleteTasks
// @Description Deletes all tasks
// @Tags tasks
// @Accept */*
// @Produce json
// @Router /tasks [DELETE]
func deleteTasks(c *gin.Context) {
	rowsAffected, err := deleteTasks_sql()
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err)
		return
	}

	message := strconv.Itoa(int(rowsAffected)) + " tasks deleted"
	c.IndentedJSON(http.StatusOK, message)
}

// deleteTaskByID godoc
// @Summary deleteTaskByID
// @Description Deletes tast at specified ID
// @Tags tasks
// @Param id path int true "The specified ID"
// @Accept */*
// @Produce json
// @Router /tasks/id/{id} [DELETE]
func deleteTaskByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	rowsAffected, err := deleteTaskByID_sql(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err)
		return
	}

	message := strconv.Itoa(int(rowsAffected)) + " task deleted"
	c.IndentedJSON(http.StatusOK, message)
}

// getTasksByTitle godoc
// @Summary getTasksByTitle
// @Description Gets all tasks whose title includes the specified string
// @Tags tasks
// @Param title path string true "The specified string"
// @Accept */*
// @Produce json
// @Router /tasks/title/{title} [get]
func getTasksByTitle(c *gin.Context) {
	title := c.Param("title")

	taskList, err := getTasksByTitle_sql(title)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err)
		return
	}
	c.IndentedJSON(http.StatusOK, taskList)
}

// patchCompleteByID godoc
// @Summary patchCompleteByID
// @Description Toggles complete at specified ID
// @Tags tasks
// @Param id path int true "The specified ID"
// @Accept */*
// @Produce json
// @Router /tasks/id/{id} [PATCH]
func patchCompleteByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	rowsAffected, err := patchCompleteByID_sql(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err)
		return
	}

	message := strconv.Itoa(int(rowsAffected)) + " task toggled"
	c.IndentedJSON(http.StatusOK, message)
}
