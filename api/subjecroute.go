package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"challenge6.2/response"
	"challenge6.2/service"
)

func SubjectData(w http.ResponseWriter, r *http.Request) {
	// var message = "User Page"

	datasubject := service.Subjectsekolah()
	// json.NewEncoder(w).Encode(user) // cara1
	var pesan response.Response
	pesan.Messages = "Get Data from Subject Database"
	pesan.Data = datasubject

	byteOfUser3, err := json.Marshal(pesan) //cara2

	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(byteOfUser3)
	fmt.Println("Endpoint Hit: userPage")
}
func UserSubjectHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		createNewSubject(w, r)
	case "GET":
		findSubjectData(w, r)
	case "PUT":
		updateDataSubject(w, r)
	case "DELETE":
		deleteDataSubject(w, r)
	}
}

func findSubjectData(w http.ResponseWriter, r *http.Request) {
	var response response.Response
	w.Header().Set("Content-Type", "application/json")
	var id = r.URL.Query()
	idSubject := id["id"]
	idIndex := idSubject[0]
	result, err := service.FindSubject(idIndex)
	if err != nil {
		log.Fatal(err)
	}
	response.Messages = "Find Data for Subject Database"
	response.Data = (result)[0]
	//mengubah data struct menjadi JSON
	byteOfSubject, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(byteOfSubject)
	http.Error(w, "", http.StatusBadRequest)
	fmt.Println("Endpoint Hit: FindSubjectPage")
}

func createNewSubject(w http.ResponseWriter, r *http.Request) {
	var subject service.Subject

	// Cara Ke 1 Pake New Decoder
	_ = json.NewDecoder(r.Body).Decode(&subject) // json ke struct
	service.CreateSubject(subject.Id, subject.Matpel)
	var pesan response.Response
	pesan.Messages = "Create Data for Subject Database"
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAlreadyReported)
	fmt.Println("Endpoint Hit: CreateSubjectPage")
}

func updateDataSubject(w http.ResponseWriter, r *http.Request) {
	var subject service.Subject

	// Cara Ke 1 Pake New Decoder
	_ = json.NewDecoder(r.Body).Decode(&subject) // json ke struct
	service.UpdateSubject(subject.Id, subject.Matpel)
	var pesan response.Response
	pesan.Messages = "Update Data for Subject Database"
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAlreadyReported)
	fmt.Println("Endpoint Hit: UpdateSubjectPage")
}

func deleteDataSubject(w http.ResponseWriter, r *http.Request) {
	ID := r.FormValue("id")
	service.DeleteSubject(ID)
	var pesan response.Response
	pesan.Messages = "Delete Data for Subject Database"
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAlreadyReported)
	fmt.Println("Endpoint Hit: DeleteSubjectPage")
}
