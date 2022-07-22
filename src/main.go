package main

import (
	
	"net/http"
	"fmt"
	"log"
	"encoding/json"
	"strconv"
	"math/rand"
	"time"
	"github.com/gorilla/mux"

)

type Tasks struct{
	ID          string `json:"id"`
	TaskName    string `json:"task_name"`
	TaskDetails string `json:"task_details"`
	Date        string `json:"date"`
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
		params := mux.Vars(r)
		
		fmt.Println(params["id"])
		flag := false
		for i := 0; i < len(tasks); i++ {
			if params["id"] == tasks[i].ID {
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
		currentTime := time.Now().Format("01-02-2006")
		task.Date = currentTime
		tasks = append(tasks, task)
		json.NewEncoder(w).Encode(task)
	}
	
	func deleteTask (w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		flag := false
		for index, item := range tasks {
			if item.ID == params["id"]{	
				tasks = append(tasks[:index], tasks[index+1:]...)
				flag = true
				json.NewEncoder(w).Encode(map[string]string{"status": "Success"})
				return
			}
		}
		if flag == false{
			json.NewEncoder(w).Encode(map[string]string{"status": "Error"})
		}
	}
	
	func updateTask (w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		flag := false
		for index, item := range tasks {
			if item.ID == params["id"]{
				
				tasks = append(tasks[:index], tasks[index+1:]...)
				var task Tasks
				_= json.NewDecoder(r.Body).Decode(&task)
				task.ID = params["id"]
				currentTime := time.Now().Format("01-02-2006")
		task.Date = currentTime
				tasks = append(tasks, task)
				flag = true
				json.NewEncoder(w).Encode(task)
				return
			}
		}
		if flag == false{
			json.NewEncoder(w).Encode(map[string]string{"status": "Error"})
		}
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

	taskss := Tasks{
		ID: "1",
		TaskName: "New Flutter",
		TaskDetails: "You must lead the project and finish it",
		Date: "2022-01-22",
	}

	tasks = append(tasks, taskss)
	fmt.Println("your tasks are", tasks)
}


	


func handleRoute(){
	route := mux.NewRouter()
	route.HandleFunc("/", homePage).Methods("GET")
	route.HandleFunc("/gettasks", getTasks).Methods("GET")
	route.HandleFunc("/gettask/", getTask).Queries("id", "{id}").Methods("GET")
	route.HandleFunc("/create", createTask).Methods("POST")
	route.HandleFunc("/delete/", deleteTask).Queries("id", "{id}").Methods("DELETE")
	route.HandleFunc("/update/", updateTask).Queries("id", "{id}").Methods("PUT")

	log.Fatal(http.ListenAndServe(":8082", route))
}
func main()  {
	
	allTasks()
	fmt.Println("Hello Flutter boys")
	handleRoute()
}
