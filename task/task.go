package task

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-redis/redis"
	"github.com/google/uuid"
)

var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

// TODO FIX THIS TO CREATE MULTIPLE TASKS
func SetTask(w http.ResponseWriter, r *http.Request) {
	message := "Create Task"
	fmt.Fprintln(w, message)

	// con := context.Background()
	akID := uuid.NewString()
	akKey := fmt.Sprintf("Task:%s", akID)

	type Task struct {
		ID      string
		Name    string `json:"name"`
		Content string `json:"Content"`
		Time    string `json:"Time"`
	}

	jsonString, err := json.Marshal(Task{
		ID:      akID,
		Name:    "Task1",
		Content: "Do the dishes",
		Time:    time.DateTime,
	})
	if err != nil {
		fmt.Printf("Failed to marshal: %s", err.Error())
		return
	}

	err = rdb.Set(akKey, jsonString, 0).Err()
	if err != nil {
		fmt.Printf("Failed to set the value in the redis instance: %s", err.Error())
		return
	}

	val, err := rdb.Get(akKey).Result()
	if err != nil {
		fmt.Printf("Failed to get the value in the redis instance: %s", err.Error())
		return
	}

	fmt.Printf("Value set: %s\n", val)

}

func GetTask(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	key := fmt.Sprintf("Task:%s", id)

	val, err := rdb.Get(key).Result()
	if err != nil {
		fmt.Printf("Failed to set the value in the redis instance: %s", err.Error())
		return
	}

	fmt.Fprintf(w, "TASK: %s\n", val)

}

func GetTasks(w http.ResponseWriter, r *http.Request) {

	// return all keys in rdb
	// loop through them
	// return values of key

	iter := rdb.Scan(0, "Task:*", 0).Iterator()
	for iter.Next() {
		// fmt.Fprintln(w, iter.Val())
		val, err := rdb.Get(iter.Val()).Result()

		if err != nil {
			fmt.Printf("Failed to get value for key: %s", err.Error())
			return
		}
		fmt.Fprintf(w, "TASK: %s\n", val)
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}

}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	message := "Update Task"
	fmt.Fprintln(w, message)
}
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	message := "Delete Task"
	fmt.Fprintln(w, message)
}

// TODO Fix this method, currently broken
func DeleteTasks(w http.ResponseWriter, r *http.Request) {

	iter := rdb.Scan(0, "Task:*", 0).Iterator()
	for iter.Next() {
		key := iter.Val()
		d, err := rdb.TTL(key).Result()

		if err != nil {
			fmt.Printf("Failed to delete key: %s", err.Error())
			panic(err)
		}

		if d == -1 { // -1 means no TTL
			fmt.Printf("Deleting task: %s", key)
			if err := rdb.Del(key).Err(); err != nil {
				panic(err)
			}
			fmt.Printf("Deleted task: %s", key)
		}
	}

	if err := iter.Err(); err != nil {
		panic(err)
	}
}
