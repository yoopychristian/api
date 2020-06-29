package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"challenge6.2/response"
	"challenge6.2/service"
)

func TeacherData(w http.ResponseWriter, r *http.Request) {
	// var message = "User Page"

	dataguru := service.Teacher()
	// json.NewEncoder(w).Encode(user) // cara1
	var pesan response.Response
	pesan.Messages = "Get Data from Teacher Database"
	pesan.Data = dataguru

	byteOfUser2, err := json.Marshal(pesan) //cara2

	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(byteOfUser2)
	fmt.Println("Endpoint Hit: TeacherPage")
}
func UserTeacherHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		createNewTeacher(w, r)
	case "GET":
		findTeacherData(w, r)
	case "PUT":
		updateDataTeacher(w, r)
	case "DELETE":
		deleteDataTeacher(w, r)
	}
}

func findTeacherData(w http.ResponseWriter, r *http.Request) {
	// var message = "User Page"
	var response response.Response
	w.Header().Set("Content-Type", "application/json")
	var id = r.URL.Query()
	idTeacher := id["id"]
	idIndex := idTeacher[0]
	result, err := service.FindTeacher(idIndex)
	if err != nil {
		log.Fatal(err)
	}
	response.Messages = "Find Data for Teacher Database"
	response.Data = (result)[0]
	//mengubah data struct menjadi JSON
	byteOfTeacher, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(byteOfTeacher)
	http.Error(w, "", http.StatusBadRequest)
	fmt.Println("Endpoint Hit: FindTeacherPage")
}
func createNewTeacher(w http.ResponseWriter, r *http.Request) {
	var teacher service.DataGuru

	// Cara Ke 1 Pake New Decoder
	_ = json.NewDecoder(r.Body).Decode(&teacher) // json ke struct
	service.CreateTeacher(teacher.Id, teacher.Firstname, teacher.Lastname, teacher.Email)
	var pesan response.Response
	pesan.Messages = "Create Data for Teacher Database"
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAlreadyReported)
	fmt.Println("Endpoint Hit: CreateTeacherPage")
}

func updateDataTeacher(w http.ResponseWriter, r *http.Request) {
	var teacher service.DataGuru

	// Cara Ke 1 Pake New Decoder
	_ = json.NewDecoder(r.Body).Decode(&teacher) // json ke struct
	service.UpdateTeacher(teacher.Id, teacher.Firstname, teacher.Lastname, teacher.Email)
	var pesan response.Response
	pesan.Messages = "Update Data for Teacher Database"
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAlreadyReported)
	fmt.Println("Endpoint Hit: UpdateTeacherPage")
}

func deleteDataTeacher(w http.ResponseWriter, r *http.Request) {
	ID := r.FormValue("id")
	service.DeleteTeacher(ID)
	var pesan response.Response
	pesan.Messages = "Delete Data for Teacher Database"
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAlreadyReported)
	fmt.Println("Endpoint Hit: DeleteTeacherPage")
}
