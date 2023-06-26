package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/go-sql-driver/mysql"
)

type Task struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Complete string `json:"complete"`
}

var tasks = []Task{
	{ID: "1", Title: "Wake up", Complete: "false"},
	{ID: "2", Title: "Go to work", Complete: "false"},
	{ID: "3", Title: "Make dinner", Complete: "false"},
}

var db *sql.DB

func main() {
	//Setting up database
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "tasks",
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	//Executing commands
	tskID, err := addTask(Task{
		ID:       "0",
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

	// router := gin.Default()
	// router.GET("/health", getHealth)
	// router.GET("/tasks", getTasks)
	// router.GET("/tasks/:id", getTaskByID)
	// router.POST("/tasks", postTasks)
	// router.PUT("/tasks/:id", putTasks)
	// router.DELETE("/tasks/:id", deleteTask)

	// router.Run("localhost:8080")
}

// Old CRUD commands
func getHealth(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "OK")
}

func getTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tasks)
}

func getTaskByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range tasks {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found:"})
}

func postTasks(c *gin.Context) {
	var newTask Task

	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	tasks = append(tasks, newTask)
	c.IndentedJSON(http.StatusCreated, newTask)
}

func putTasks(c *gin.Context) {
	var newTask Task
	var id string = c.Param("id")

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

func deleteTask(c *gin.Context) {
	var id string = c.Param("id")

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

// New SQL commands
func addTask(tsk Task) (int64, error) {
	result, err := db.Exec("INSERT INTO tasks (ID, Title, Complete) VALUES (?, ?, ?)", tsk.ID, tsk.Title, tsk.Complete)
	if err != nil {
		return 0, fmt.Errorf("addTask: %v", err)
	}
	id, err := result.LastInsertId()
	fmt.Println(id)
	if err != nil {
		return 0, fmt.Errorf("addTask: %v", err)
	}
	return id, nil
}

func tasksByComplete(complete string) ([]Task, error) {
	var tasks []Task

	rows, err := db.Query("SELECT * FROM tasks WHERE complete = ?", complete)
	if err != nil {
		return nil, fmt.Errorf("tasksByComplete %q: %v", complete, err)
	}
	defer rows.Close()

	for rows.Next() {
		var tsk Task
		if err := rows.Scan(&tsk.ID, &tsk.Title, &tsk.Complete); err != nil {
			return nil, fmt.Errorf("tasksByComplete %q: %v", complete, err)
		}
		tasks = append(tasks, tsk)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByComplete %q: %v", complete, err)
	}
	return tasks, nil
}

func taskByID(id int64) (Task, error) {
	var tsk Task

	row := db.QueryRow("SELECT * FROM tasks WHERE id = ?", id)
	if err := row.Scan(&tsk.ID, &tsk.Title, &tsk.Complete); err != nil {
		if err == sql.ErrNoRows {
			return tsk, fmt.Errorf("taskByID %d: no such task", id)
		}
		return tsk, fmt.Errorf("taskByID %d: %v", id, err)
	}
	return tsk, nil
}
