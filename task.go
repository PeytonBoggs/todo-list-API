package main

type Task struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Complete bool   `json:"complete"`
}

type shortTask struct {
	Title    string `json:"title"`
	Complete bool   `json:"complete"`
}
