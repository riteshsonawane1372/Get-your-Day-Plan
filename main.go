package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ListInfo struct {
	ID      string 	`json:"id"`
	Title   string `json:"title"`
	Date    string    `json:"date"`
	Time    string    `json:"Time"`
	Day     string `json:"day"`
	Message string `json:"message"`
}

// initally the list will be empty

var list []ListInfo

func main() {

	fmt.Println("Welcome to the Golang todo")

	r:= mux.NewRouter()

	list= append(list,
		ListInfo{
			ID:"1120",
			Title: "First Todo",
			Date: "07",
			Time: "19",
			Day: "Monday",
			Message: "First Todo",		
	})

	r.HandleFunc("/",AllTodo).Methods("GET")
	r.HandleFunc("/{id}",TodoByID).Methods("GET")
	r.HandleFunc("/{id}",EditTodo).Methods("PUT")
	r.HandleFunc("/{id}",DeleteTodo).Methods("DELETE")

	fmt.Println("Server is running on port 3000")
	log.Fatal(http.ListenAndServe("127.0.0.1:3000",r))
	

}

func AllTodo(w http.ResponseWriter,r *http.Request){

	w.Header().Set("Content-type","application.json")
	json.NewEncoder(w).Encode(list)

}
func TodoByID(w http.ResponseWriter, r*http.Request){
	w.Header().Set("Content-Type","application/json")

	paramter:= mux.Vars(r)

	for _,item:=range list{
		if paramter["id"] == item.ID{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode("Not Found")
	return
}

func EditTodo(w http.ResponseWriter, r*http.Request){

	w.Header().Set("Content-Type","application/json")

	parameter := mux.Vars(r)

	for _,item := range list{

		if parameter["id"]==item.ID{
			item.Title=parameter["title"]
			item.Date=parameter["date"]
			item.Day=parameter["day"]
			item.Message=parameter["message"]
			item.Time=parameter["time"]

			json.NewEncoder(w).Encode(item)
			return
		}


	}
	json.NewEncoder(w).Encode("Enter a Valid ID")

}

func DeleteTodo(w http.ResponseWriter, r*http.Request){

	w.Header().Set("Content-Type","application/json")

	params := mux.Vars(r)

	for index,item:= range list{
		if params["id"] == item.ID{
			list=append(list[:index],list[index+1:]... )
		}
	}
	json.NewEncoder(w).Encode(list)

}