package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

// Initializes SQL database
func initSQL() {
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
}

// Adds task to the end of SQL database
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

// Returns all tasks in SQL database with specified 'complete' state
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

// Returns task in SQL database with specified ID
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
