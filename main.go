package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func main() {

	fmt.Println("GO REDIS")
	con := context.Background()

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	type Task struct {
		ID      string
		Name    string `json:"name"`
		Content string `json:"Content"`
		Date    string `json:"Date"`
	}

	akID := uuid.NewString()
	jsonString, err := json.Marshal(Task{
		ID:      akID,
		Name:    "Task1",
		Content: "Do the dishes",
		Date:    "Today",
	})
	if err != nil {
		fmt.Printf("Failed to marshal: %s", err.Error())
		return
	}

	err = client.Set(con, "Task", jsonString, 0).Err()
	if err != nil {
		fmt.Printf("Failed to set the value in the redis instance: %s", err.Error())
		return
	}

	val, err := client.Get(con, "Task").Result()
	if err != nil {
		fmt.Printf("Failed to set the value in the redis instance: %s", err.Error())
		return
	}

	ping, err := client.Ping(con).Result()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(ping)

	fmt.Printf("Value retrieved: %s\n", val)
}
