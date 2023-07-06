package main

import (
	"fmt"
	"strconv"
)

// Returns all tasks in SQL database that match input filter
func getTasksByFilter_sql(searchedID string, searchedTitle string, searchedComplete string) ([]Task, error) {
	var taskList []Task

	query := "SELECT * FROM tasks WHERE 1=1"

	if searchedID != "" {
		id, err := strconv.Atoi(searchedID)
		if err != nil {
			return taskList, err
		}
		query += " AND id = " + strconv.Itoa(id)
	}

	if searchedTitle != "" {
		query += " AND title LIKE '%" + searchedTitle + "%'"
	}

	if searchedComplete != "" {
		complete, err := strconv.ParseBool(searchedComplete)
		if err != nil {
			return taskList, err
		}
		query += " AND complete = " + strconv.FormatBool(complete)
	}

	tasks, err := db.Query(query)
	if err != nil {
		return taskList, err
	}

	for tasks.Next() {
		var tsk Task

		if err := tasks.Scan(&tsk.ID, &tsk.Title, &tsk.Complete); err != nil {
			return taskList, err
		}

		taskList = append(taskList, tsk)
	}

	if err := tasks.Err(); err != nil {
		return taskList, err
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

// Deletes all tasks in SQL database
func deleteTasks_sql() (int64, error) {
	result, err := db.Exec("DELETE FROM tasks")
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	_, err = db.Exec("TRUNCATE TABLE tasks")
	if err != nil {
		return 0, err
	}

	_, err = db.Exec("ALTER TABLE tasks AUTO_INCREMENT = 1")
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

// Deletes task in SQL database with specified ID
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
