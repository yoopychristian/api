package main

import (
	"fmt"
	"log"
	"net/http"

	"challenge6.2/api"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//Endpoint
	http.HandleFunc("/students", api.StudentData)
	http.HandleFunc("/student", api.UserStudentHandler)
	http.HandleFunc("/teachers", api.TeacherData)
	http.HandleFunc("/teacher", api.UserTeacherHandler)
	http.HandleFunc("/subjects", api.SubjectData)
	http.HandleFunc("/subject", api.UserSubjectHandler)

	// http.HandleFunc("/productpage", productPage)

	//server
	fmt.Println("Running On Port 3000")
	err := http.ListenAndServe("localhost:3000", nil)

	if err != nil {
		log.Fatal(err)
	}
}

// func userTeacherHandler(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case "POST":
// 		FindTeacherByID(w, r)
// 	case "GET":
// 		CreateNewTeacher(w, r)
// 	case "PUT":
// 		UpdateTeacher(w, r)
// 	case "Delete":
// 		DeleteTeacher(w, r)
// 	}
// }

// func userSubjectHandler(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case "POST":
// 		FindSubjectByID(w, r)
// 	case "GET":
// 		CreateNewSubject(w, r)
// 	case "PUT":
// 		UpdateSubject(w, r)
// 	case "Delete":
// 		DeleteSubject(w, r)
// 	}
// }
