package main

type Task struct {
	ID       int64  `json:"id"`
	Title    string `json:"title"`
	Complete string `json:"complete"`
}

var tasks = []Task{
	{ID: 1, Title: "Wake up", Complete: "false"},
	{ID: 2, Title: "Go to work", Complete: "false"},
	{ID: 3, Title: "Make dinner", Complete: "false"},
}
