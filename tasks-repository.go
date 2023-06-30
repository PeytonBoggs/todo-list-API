package main

import (
	// "database/sql"
	"fmt"
)

// Returns task in SQL database with specified ID
// func getTaskByID_sql(id int64) (Task, error) {
// 	var tsk Task

// 	row := db.QueryRow("SELECT * FROM tasks WHERE id = ?", id)
// 	if err := row.Scan(&tsk.Title, &tsk.Complete); err != nil {
// 		if err == sql.ErrNoRows {
// 			return tsk, fmt.Errorf("taskByID %d: no such task", id)
// 		}
// 		return tsk, fmt.Errorf("taskByID %d: %v", id, err)
// 	}
// 	return tsk, nil
// }

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

// Returns all tasks in SQL database with specified 'complete' state
// func tasksByComplete(complete string) ([]Task, error) {
// var tasks []Task

// rows, err := db.Query("SELECT * FROM tasks WHERE complete = ?", complete)
// if err != nil {
// 	return nil, fmt.Errorf("tasksByComplete %q: %v", complete, err)
// }
// defer rows.Close()

// for rows.Next() {
// 	var tsk Task
// 	if err := rows.Scan(&tsk.ID, &tsk.Title, &tsk.Complete); err != nil {
// 		return nil, fmt.Errorf("tasksByComplete %q: %v", complete, err)
// 	}
// 	tasks = append(tasks, tsk)
// }

// if err := rows.Err(); err != nil {
// 	return nil, fmt.Errorf("albumsByComplete %q: %v", complete, err)
// }
// return tasks, nil
// }
