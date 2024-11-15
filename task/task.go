package task

import (
	"fmt"
	"net/http"
)

func SetTask(w http.ResponseWriter, r *http.Request) {
	message := "Create Task"
	fmt.Fprintln(w, message)
	// con := context.Background()

	// client := redis.NewClient(&redis.Options{
	// 	Addr:     "localhost:6379",
	// 	Password: "",
	// 	DB:       0,
	// })

	// type Task struct {
	// 	ID      string
	// 	Name    string `json:"name"`
	// 	Content string `json:"Content"`
	// 	Time    string `json:"Time"`
	// }

	// akID := uuid.NewString()
	// jsonString, err := json.Marshal(Task{
	// 	ID:      akID,
	// 	Name:    "Task1",
	// 	Content: "Do the dishes",
	// 	Time:    time.DateTime,
	// })
	// if err != nil {
	// 	fmt.Printf("Failed to marshal: %s", err.Error())
	// 	return
	// }

	// akKey := fmt.Sprintf("Task:%s", akID)
	// err = client.Set(con, akKey, jsonString, 0).Err()
	// if err != nil {
	// 	fmt.Printf("Failed to set the value in the redis instance: %s", err.Error())
	// 	return
	// }

	// val, err := client.Get(con, akKey).Result()
	// if err != nil {
	// 	fmt.Printf("Failed to set the value in the redis instance: %s", err.Error())
	// 	return
	// }

	// fmt.Printf("Value retrieved: %s\n", val)

}

func GetTask(w http.ResponseWriter, r *http.Request) {
	message := "Read Task"
	fmt.Fprintln(w, message)
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	message := "Read all tasks"
	fmt.Fprintln(w, message)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	message := "Update Task"
	fmt.Fprintln(w, message)
}
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	message := "Delete Task"
	fmt.Fprintln(w, message)
}

func DeleteTasks(w http.ResponseWriter, r *http.Request) {
	message := "Delete all tasks"
	fmt.Fprintln(w, message)
}
