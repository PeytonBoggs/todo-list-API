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

// Toggles the complete boolean of the task with input id in SQL database
func patchCompleteByID_sql(id int) (int64, error) {
	result, err := db.Exec("UPDATE tasks SET complete = CASE WHEN complete = true THEN false WHEN complete = false THEN true END WHERE id = (?)", id)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}
