package main

import (
	"fmt"
	"go_api/task"
	"net/http"
)

func main() {
	fmt.Println("REST API with Redis and Go!")
	mux := http.NewServeMux()

	// Define routes
	mux.HandleFunc("POST /api/v1/task", task.SetTask)
	mux.HandleFunc("GET /api/v1/task/{id}", task.GetTask)
	mux.HandleFunc("GET /api/v1/task", task.GetTasks)
	mux.HandleFunc("UPDATE /api/v1/task/{id}", task.UpdateTask)
	mux.HandleFunc("DELETE /api/v1/task/{id}", task.DeleteTask)
	mux.HandleFunc("DELETE /api/v1/task", task.DeleteTasks)

	// Start server on port 8080 and if err print error
	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		fmt.Println(err.Error())
	}

}
