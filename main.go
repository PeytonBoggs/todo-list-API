package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	initSQL()

	tskID, err := addTask(Task{
		ID:       0,
		Title:    "Go to the gym",
		Complete: "false",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of added task: %v\n", tskID)

	tasks, err := tasksByComplete("false")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Tasks found: %v\n", tasks)

	tsk, err := taskByID(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Task found: %v\n", tsk)

	router := gin.Default()

	router.GET("/health", getHealth)
	router.GET("/tasks", getTasks)
	router.GET("/tasks/:id", getTaskByID)
	router.POST("/tasks", postTasks)
	router.PUT("/tasks/:id", putTasks)
	router.DELETE("/tasks/:id", deleteTask)

	router.Run("localhost:8080")
}
