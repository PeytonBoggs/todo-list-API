package main

import (
	"database/sql"
	"fmt"
)

// TODO: add func getTasks_sql

// Returns task in SQL database with specified ID
func getTaskByID_sql(id int) (Task, error) {
	var foundTask Task

	row := db.QueryRow("SELECT * FROM tasks WHERE id = ?", id)
	if err := row.Scan(&foundTask.ID, &foundTask.Title, &foundTask.Complete); err != nil {
		if err == sql.ErrNoRows {
			return foundTask, fmt.Errorf("taskByID %d: no such task", id)
		}
		return foundTask, fmt.Errorf("taskByID %d: %v", id, err)
	}
	return foundTask, nil
}

// TODO: add func getTasksByComplete_sql

// Adds task to the end of SQL database
func postTask_sql(tsk shortTask) (int64, error) {
	result, err := db.Exec("INSERT INTO tasks (Title, Complete) VALUES (?, ?)", tsk.Title, tsk.Complete)
	if err != nil {
		return 0, fmt.Errorf("error: %v", err)
	}

	newID, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error: %v", err)
	}
	return newID, nil
}

// TODO: add func putTasks_sql

// TODO: add func deleteTask_sql
