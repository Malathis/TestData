package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Order struct {
	Id   string `json:"id"`
	Name string `json:"Order"`
}

type Schedule struct {
	Id   int    `json:"id"`
	Name string `json:"schedule"`
}

type Screen struct {
	Id   string `json:"code"`
	Name string `json:"screen"`
}

type Show struct {
	Id   int    `json:"id"`
	Name string `json:"show"`
}

var Orders []Order
var Schedules []Schedule
var Screens []Screen
var Shows []Show

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "TestData homepage endpoint hit")
}

func getOrders(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getOrders")

	Orders = []Order{
		Order{Id: "Order1", Name: "Order Name 1"},
		Order{Id: "Order2", Name: "Order Name 2"},
	}

	json.NewEncoder(w).Encode(Orders)
}

func getSchedules(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getSchedules")

	Schedules = []Schedule{
		Schedule{Id: 1, Name: "Schedule Name 1"},
		Schedule{Id: 2, Name: "Schedule Name 2"},
	}

	json.NewEncoder(w).Encode(Schedules)
}

func getScreens(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getScreens")

	Screens = []Screen{
		Screen{Id: "ScreenCode1", Name: "Screen Name 1"},
		Screen{Id: "ScreenCode2", Name: "Screen Name 2"},
	}

	json.NewEncoder(w).Encode(Screens)
}

func getShows(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getShows")

	Shows = []Show{
		Show{Id: 1, Name: "Morning"},
		Show{Id: 2, Name: "Afterrnoon"},
		Show{Id: 3, Name: "Evening"},
		Show{Id: 3, Name: "Night"},
	}

	json.NewEncoder(w).Encode(Shows)
}

func handleRequests() {
	http.HandleFunc("/", homePage)

	http.HandleFunc("/getOrders", getOrders)

	http.HandleFunc("/getSchedules", getSchedules)

	http.HandleFunc("/getScreens", getScreens)

	http.HandleFunc("/getShows", getShows)

	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}
