package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"challenge6.2/response"
	"challenge6.2/service"
)

func StudentData(w http.ResponseWriter, r *http.Request) {
	// var message = "User Page"

	datamurid := service.Student()
	// json.NewEncoder(w).Encode(user) // cara1
	var pesan response.Response
	pesan.Messages = "Get All Data from Student Database"
	pesan.Data = datamurid

	byteOfUser1, err := json.Marshal(pesan) //cara2

	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAlreadyReported)
	w.Write(byteOfUser1)
	fmt.Println("Endpoint Hit: AllStudentPage")
}

func UserStudentHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		createNewStudent(w, r)
	case "GET":
		findStudentData(w, r)
	case "PUT":
		updateDataStudent(w, r)
	case "DELETE":
		deleteDataStudent(w, r)
	}
}

func findStudentData(w http.ResponseWriter, r *http.Request) {
	// var message = "User Page"
	var response response.Response
	w.Header().Set("Content-Type", "application/json")
	var id = r.URL.Query()
	idStudent := id["id"]
	idIndex := idStudent[0]
	result, err := service.FindStudent(idIndex)
	if err != nil {
		log.Fatal(err)
	}
	response.Messages = "Find Data for Student Database"
	response.Data = (result)[0]
	//mengubah data struct menjadi JSON
	byteOfStudent, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(byteOfStudent)
	http.Error(w, "", http.StatusBadRequest)
	fmt.Println("Endpoint Hit: FindStudentPage")
}

func createNewStudent(w http.ResponseWriter, r *http.Request) {
	var student service.Datastudent

	// Cara Ke 1 Pake New Decoder
	_ = json.NewDecoder(r.Body).Decode(&student) // json ke struct
	service.CreateStudent(student.Id, student.Firstname, student.Lastname, student.Email)
	var pesan response.Response
	pesan.Messages = "Create Data for Student Database"
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAlreadyReported)
	fmt.Println("Endpoint Hit: CreateStudentPage")
}

func updateDataStudent(w http.ResponseWriter, r *http.Request) {
	var student service.Datastudent

	// Cara Ke 1 Pake New Decoder
	_ = json.NewDecoder(r.Body).Decode(&student) // json ke struct
	service.UpdateStudent(student.Id, student.Firstname, student.Lastname, student.Email)
	var pesan response.Response
	pesan.Messages = "Update Data for Student Database"
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAlreadyReported)
	fmt.Println("Endpoint Hit: UpdateStudentPage")
}

func deleteDataStudent(w http.ResponseWriter, r *http.Request) {
	ID := r.FormValue("id")
	service.DeleteStudent(ID)
	var pesan response.Response
	pesan.Messages = "Delete Data for Student Database"
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAlreadyReported)
	fmt.Println("Endpoint Hit: DeleteStudentPage")
}
