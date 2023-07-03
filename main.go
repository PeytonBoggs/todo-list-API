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
	router.GET("/tasks/:id", getTaskByID)
	router.POST("/tasks", postTask)
	router.PUT("/tasks/:id", putTasks)
	router.DELETE("/tasks/:id", deleteTaskByID)

	router.Run("localhost:8080")
}
