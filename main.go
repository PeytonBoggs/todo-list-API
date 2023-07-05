package main

import (
	"todo-list/web-service-gin/docs"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Gin Swagger todo-list API
// @version 1.0
// @description This is a todo-list server.

// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	initSQL()

	router := gin.Default()

	docs.SwaggerInfo.BasePath = ""
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/health", getHealth)
	router.GET("/tasks", getTasks)
	router.GET("/tasks/id/:id", getTaskByID)
	router.GET("/tasks/complete/:complete", getTasksByComplete)
	router.POST("/tasks", postTask)
	router.PATCH("/tasks/id/:id", patchCompleteByID)
	router.PUT("/tasks/id/:id", putTasks)
	router.DELETE("/tasks/id/:id", deleteTaskByID)
	router.DELETE("/tasks", deleteTasks)

	router.Run("localhost:8080")
}
