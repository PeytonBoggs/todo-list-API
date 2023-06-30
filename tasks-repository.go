package main

import (
	"fmt"
)

// TODO: add func getTasks_sql

// TODO: add func getTaskByID_sql

// TODO: add func getTasksByComplete_sql

// Adds task to the end of SQL database
func postTask_sql(tsk Task) (int64, error) {
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
