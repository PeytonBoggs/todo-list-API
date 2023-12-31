package main

import (
	"todo-list/web-service-gin/docs"

	"github.com/gin-contrib/cors"
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

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Type"}
	config.AllowMethods = []string{"GET", "POST", "PATCH", "DELETE"}
	router.Use(cors.New(config))

	docs.SwaggerInfo.BasePath = ""
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/tasks", getTasksByFilter)
	router.POST("/tasks", postTask)
	router.PATCH("/tasks/id/:id", patchCompleteByID)
	router.DELETE("/tasks/id/:id", deleteTaskByID)
	router.DELETE("/tasks", deleteTasks)

	router.Run("localhost:8080")
}
