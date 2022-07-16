package main

import (
	
	"net/http"
	"fmt"
	"log"
	"encoding/json"
	"strconv"
	"math/rand"
	"github.com/gorilla/mux"

)

type Tasks struct{
	ID string `json:"id"`
	TaskName string `json:"task_name"`
	TaskDetails string `json:"task_details"`
	Date string `json:"date"`
}
var tasks[]Tasks
func homePage (w http.ResponseWriter, r *http.Request){
	fmt.Println("I am home page")	
	}
	
	func getTasks(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(tasks)
	}
	
	func getTask(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		taskId := mux.Vars(r)
		flag := false
		for i := 0; i < len(tasks); i++ {
			if taskId["id"] == tasks[i].ID {
				json.NewEncoder(w).Encode(tasks[i])
				flag = true
				break
			}
		}
		if !flag {
			json.NewEncoder(w).Encode(map[string]string{"status": "Error"})
		}
	}
	
	func createTask(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var task Tasks
		_ = json.NewDecoder(r.Body).Decode(&task)
		task.ID = strconv.Itoa(rand.Intn(1000))
		tasks = append(tasks, task)
		json.NewEncoder(w).Encode(task)
	}
	
	func deleteTask (w http.ResponseWriter, r *http.Request){
		
	}
	
	func updateTask (w http.ResponseWriter, r *http.Request){
		
	}
		
		

func allTasks(){
	task := Tasks{
		ID: "1",
		TaskName: "New Projects",
		TaskDetails: "You must lead the project and finish it",
		Date: "2022-01-22",
	}

	tasks = append(tasks, task)
	fmt.Println("your tasks are", tasks)

	task1 := Tasks{
		ID: "2",
		TaskName: "Power project",
		TaskDetails: "We need to hire staffs before deadline",
		Date: "2022-01-22",
	}

	tasks = append(tasks, task1)
	fmt.Println("your tasks are", tasks)
}


	


func handleRoute(){
	route := mux.NewRouter()
	route.HandleFunc("/", homePage).Methods("GET")
	route.HandleFunc("/gettasks", getTasks).Methods("GET")
	route.HandleFunc("/gettask/{id}", getTask).Methods("GET")
	route.HandleFunc("/create", createTask).Methods("POST")
	route.HandleFunc("/delete/{id}", deleteTask).Methods("DELETE")
	route.HandleFunc("/update/{id}", updateTask).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8082", route))
}
func main()  {
	
	allTasks()
	fmt.Println("Hello Flutter boys")
	handleRoute()
}
