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

// Model for course   - file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"-"` // - export default empty
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"` //export the json
	Email    string `json:"email"`
}

// Fake DB
var courses []Course

// middlewares , helper  -file

func (c *Course) IsEmpty() bool {
	return c.CourseName == ""

}

func main() {
	fmt.Println("API- Learning API")
	r := mux.NewRouter()

	// seeding
	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJs", CoursePrice: 299, Author: &Author{Fullname: "Nischay Jain", Email: "nischayjain4948@gmail.com"}})

	courses = append(courses, Course{CourseId: "3", CourseName: "ReactNative", CoursePrice: 399, Author: &Author{Fullname: "Tanishk Jain", Email: "tanishkjain4948@gmail.com"}})

	//   route         handlerMrthod   methods verb
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourse).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))

}

// controllers are used to manage the routes easily! -file

// serveHome
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API</h1>"))
}

// getAllCourse
func getAllCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all course!!")
	w.Header().Set("Content-Type", "applications/json")
	json.NewEncoder(w).Encode(courses)

}

// getOneCourse
func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// grab the unique Id which is comes with request  and  put in a variable
	params := mux.Vars(r)

	// loop through the course array
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)

			return
		}
	}
	json.NewEncoder(w).Encode("No course found with this id")
	return
}

// createOneCourse
func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	// what if : body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some json data")
	}

	// what if: data is like {}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	// what if : already exist coursename

	for _, c := range courses {
		if c.CourseName == course.CourseName {
			json.NewEncoder(w).Encode("courseName already exist please change")
			return
		}
	}

	// generate unique id, string
	// append course into course
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

// update one course

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This is an update one course method")
	w.Header().Set("Content-Type", "application/json")

	// first grab the id which comes with request
	params := mux.Vars(r)
	// loop through the slice and remove and add again

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}

	}

	// send the response when id is not found
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")
	// loop through the course slice.. find the doc and get the id from params and delete the particular course
	param := mux.Vars(r)
	for index, course := range courses {
		if course.CourseId == param["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Course has been removed from slices")
			break
		}
	}
	json.NewEncoder(w).Encode("No course found!")

}
