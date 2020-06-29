package service

import (
	"log"

	"challenge6.2/utils"
)

type DataGuru struct {
	Id        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
}

func Teacher() []DataGuru {
	db, err := utils.Connect()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	data, err := db.Query("select * from teacher")
	if err != nil {
		log.Fatal(err)
	}

	defer data.Close()

	var result1 []DataGuru

	for data.Next() {
		var guru = DataGuru{}
		var err = data.Scan(&guru.Id, &guru.Firstname, &guru.Lastname, &guru.Email)

		if err != nil {
			log.Fatal(err)
		}

		result1 = append(result1, guru)
	}

	if err = data.Err(); err != nil {
		log.Fatal(err)
	}

	return result1

}

func FindTeacher(Id string) ([]DataGuru, error) {
	db, err := utils.Connect()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	data, err := db.Query("SELECT * FROM teacher WHERE id = ?", Id)

	if err != nil {
		log.Fatal(err)
	}

	defer data.Close()

	var result2 []DataGuru

	for data.Next() {
		var guru = DataGuru{}
		var err = data.Scan(&guru.Id, &guru.Firstname, &guru.Lastname, &guru.Email)

		if err != nil {
			log.Fatal(err)
		}

		result2 = append(result2, guru)
	}

	if err = data.Err(); err != nil {
		log.Fatal(err)
	}

	return result2, nil

}

func CreateTeacher(id, first_name, last_name, email string) []DataGuru {
	db, err := utils.Connect()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	data, err := db.Query("INSERT INTO teacher VALUES (?,?,?,?)", id, first_name, last_name, email)

	if err != nil {
		log.Fatal(err)
	}

	defer data.Close()

	var result2 []DataGuru

	for data.Next() {
		var guru = DataGuru{}
		var err = data.Scan(&guru.Id, &guru.Firstname, &guru.Lastname, &guru.Email)

		if err != nil {
			log.Fatal(err)
		}

		result2 = append(result2, guru)

		if err = data.Err(); err != nil {
			log.Fatal(err)
		}

	}
	return result2
}

func UpdateTeacher(id, first_name, last_name, email string) []DataGuru {
	db, err := utils.Connect()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	data, err := db.Query("UPDATE teacher SET first_name=?, last_name=?, email=? WHERE id=?", first_name, last_name, email, id)

	if err != nil {
		log.Fatal(err)
	}

	defer data.Close()

	var result2 []DataGuru

	for data.Next() {
		var guru = DataGuru{}
		var err = data.Scan(&guru.Id, &guru.Firstname, &guru.Lastname, &guru.Email)

		if err != nil {
			log.Fatal(err)
		}

		result2 = append(result2, guru)
	}

	if err = data.Err(); err != nil {
		log.Fatal(err)
	}

	return result2

}

func DeleteTeacher(id string) {
	db, err := utils.Connect()

	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	tx, err := db.Begin()
	stmt, err := tx.Exec("DELETE FROM teacher WHERE id=?", id)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	_, err = stmt.LastInsertId()

	tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

}
