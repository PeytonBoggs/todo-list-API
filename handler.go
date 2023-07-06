package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// TODO: Implement getHealth
func getHealth(c *gin.Context) {
}

// getTasksByFilter godoc
// @Summary getTasksByFilter
// @Description Gets all tasks in database that pass applied filter
// @Tags tasks
// @Param id query int false "id"
// @Param title query string false "title"
// @Param complete query boolean false "complete"
// @Accept */*
// @Produce json
// @Router /tasks [GET]
func getTasksByFilter(c *gin.Context) {
	searchedID := c.Query("id")

	searchedTitle := c.Query("title")

	searchedComplete := c.Query("complete")

	filteredList, err := getTasksByFilter_sql(searchedID, searchedTitle, searchedComplete)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	}

	c.IndentedJSON(http.StatusOK, filteredList)
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
