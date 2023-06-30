package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// getHealth godoc
// @Summary getHealth
// @Description Returns the health of the server - currently hardcoded to "OK"
// @Tags root
// @Accept */*
// @Produce json
// @Router /health [get]
func getHealth(c *gin.Context) {
	// c.IndentedJSON(http.StatusOK, "OK")
}

// getTasks godoc
// @Summary getTasks
// @Description Gets all tasks in database
// @Tags root
// @Accept */*
// @Produce json
// @Router /tasks [get]
func getTasks(c *gin.Context) {
	// c.IndentedJSON(http.StatusOK, tasks)
}

// getTaskByID godoc
// @Summary getTaskByID
// @Description Gets all tasks with specified ID
// @Tags root
// @Param id path int true "The specified ID"
// @Accept */*
// @Produce json
// @Router /tasks/{id} [get]
func getTaskByID(c *gin.Context) {
	// 	var taskToSearch Task

	// 	if err := c.BindJSON(&taskToSearch); err != nil {
	// 		c.IndentedJSON(http.StatusBadRequest, err)
	// 		return
	// 	}

	// foundTask, foundErr := getTaskByID_sql(taskToSearch)
	//
	//	if foundErr != nil {
	//		c.IndentedJSON(http.StatusNotFound, foundErr)
	//		return
	//	}
	//
	// c.IndentedJSON(http.StatusOK, foundTask)
}

// postTasks godoc
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

// putTasks godoc
// @Summary putTasks
// @Description Adds new task at the specified ID, or end of database if ID can't be found
// @Tags root
// @RequestBody required
// @Param task body Task true "The task to add"
// @Param id path int true "The specified ID"
// @Accept */*
// @Produce json
// @Router /tasks/{id} [PUT]
func putTasks(c *gin.Context) {
	// var newTask Task
	// id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	// if err != nil {
	// 	return
	// }

	// if err := c.BindJSON(&newTask); err != nil {
	// 	c.IndentedJSON(http.StatusBadRequest, gin.H{"Error": "could not bind JSON"})
	// }

	// for i := 0; i < len(tasks); i++ {
	// 	if tasks[i].ID == id {
	// 		tasks[i] = newTask
	// 		c.IndentedJSON(http.StatusOK, newTask)
	// 		return
	// 	}
	// }
	// tasks = append(tasks, newTask)
	// c.IndentedJSON(http.StatusCreated, newTask)
}

// deleteTask godoc
// @Summary deleteTask
// @Description Deletes tast at specified ID
// @Tags root
// @Param id path int true "The specified ID"
// @Accept */*
// @Produce json
// @Router /tasks/{id} [DELETE]
func deleteTask(c *gin.Context) {
	// id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	// if err != nil {
	// 	return
	// }

	// for i := 0; i < len(tasks); i++ {
	// 	if tasks[i].ID == id {
	// 		firstHalf := tasks[0:i]
	// 		secondHalf := tasks[i+1:]
	// 		tasks = append(firstHalf, secondHalf...)
	// 		c.IndentedJSON(http.StatusOK, tasks)
	// 		return
	// 	}
	// }
	// c.IndentedJSON(http.StatusNotFound, gin.H{"Error": "ID not found"})
}
