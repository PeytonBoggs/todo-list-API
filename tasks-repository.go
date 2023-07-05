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

// Returns all tasks in SQL database with specified Complete value
func getTasksByComplete_sql(complete bool) ([]Task, error) {
	var taskList []Task

	completeTasks, err := db.Query("SELECT * FROM tasks where COMPLETE = ?", complete)
	if err != nil {
		return taskList, err
	}

	for completeTasks.Next() {
		var tsk Task

		if err := completeTasks.Scan(&tsk.ID, &tsk.Title, &tsk.Complete); err != nil {
			return taskList, err
		}

		taskList = append(taskList, tsk)
	}

	return taskList, nil
}

// Adds task to the end of SQL database
func postTask_sql(tsk TaskPayload) (int64, error) {
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

func deleteTaskByID_sql(id int) (int64, error) {
	result, err := db.Exec("DELETE FROM tasks WHERE id=(?)", id)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}
