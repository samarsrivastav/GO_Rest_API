package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// model  for courses

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}
type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake  db
var courses []Course

// middleware , helper -file
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("-- API --")
	r := mux.NewRouter()

	courses = append(courses, Course{CourseId: "2", CourseName: "react", CoursePrice: 299, Author: &Author{Fullname: "samar k", Website: "react.dev"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "mern", CoursePrice: 199, Author: &Author{Fullname: "samar k", Website: "mern.dev"}})

	// routing
	r.HandleFunc("/",serveHome).Methods("GET")
	r.HandleFunc("/courses",getAllCourse).Methods("GET")
	r.HandleFunc("/course/{id}",getOneCourse).Methods("GET")
	r.HandleFunc("/course",createCourse).Methods("POST")
	r.HandleFunc("/course/{id}",updateCourse).Methods("PUT")
	r.HandleFunc("/course/{id}",deleteCourse).Methods("DELETE")
	// listen to a port

	log.Fatal(http.ListenAndServe(":4000", r))
}










// controller end points

// serve home
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1> welcome  to home server </h1>"))
}

func getAllCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All COurses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}
func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// grab one value
	params := mux.Vars(r)

	//loop and output
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode(" = Nothing found =")
	return
}
func createCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	// what if : body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	// what about -{}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Please send some data inside json")
		return
	}

	// generate unique id , string
	// apend couse  into Courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)

	json.NewEncoder(w).Encode(course)
	return
}
func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	// first - grab id from req
	params := mux.Vars(r)

	// loop, id, remove ,add new
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)

			json.NewEncoder(w).Encode(course)
			return
		}

	}
	// TODO  : send a response when id is not found
}
func deleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DELETE one course")
	w.Header().Set("Content-Type", "application/json")

	// first - grab id from req
	params := mux.Vars(r)

	// loop, id, remove
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}

	}
	// TODO  : send a response when id is not found
}
