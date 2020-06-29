package service

import (
	"log"

	"challenge6.2/utils"
)

type Datastudent struct {
	Id        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
}

func Student() []Datastudent {
	db, err := utils.Connect()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	data, err := db.Query("SELECT * FROM student")

	if err != nil {
		log.Fatal(err)
	}

	defer data.Close()

	var result2 []Datastudent

	for data.Next() {
		var siswa = Datastudent{}
		var err = data.Scan(&siswa.Id, &siswa.Firstname, &siswa.Lastname, &siswa.Email)
		if err != nil {
			log.Fatal(err)
		}

		result2 = append(result2, siswa)
	}

	if err = data.Err(); err != nil {
		log.Fatal(err)
	}

	return result2

}

func FindStudent(id string) ([]Datastudent, error) {
	db, err := utils.Connect()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	data, err := db.Query("SELECT * FROM student WHERE id = ?", id)

	if err != nil {
		log.Fatal(err)
	}

	defer data.Close()

	var result2 []Datastudent

	for data.Next() {
		var siswa = Datastudent{}
		var err = data.Scan(&siswa.Id, &siswa.Firstname, &siswa.Lastname, &siswa.Email)

		if err != nil {
			log.Fatal(err)
		}

		result2 = append(result2, siswa)
	}

	if err = data.Err(); err != nil {
		log.Fatal(err)
	}

	return result2, nil

}

func CreateStudent(id, first_name, last_name, email string) []Datastudent {
	db, err := utils.Connect()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	data, err := db.Query("INSERT INTO student VALUES (?,?,?,?)", id, first_name, last_name, email)

	if err != nil {
		log.Fatal(err)
	}

	defer data.Close()

	var result2 []Datastudent

	for data.Next() {
		var siswa = Datastudent{}
		var err = data.Scan(&siswa.Id, &siswa.Firstname, &siswa.Lastname, &siswa.Email)

		if err != nil {
			log.Fatal(err)
		}

		result2 = append(result2, siswa)
	}

	if err = data.Err(); err != nil {
		log.Fatal(err)
	}

	return result2

}

func UpdateStudent(id, first_name, last_name, email string) []Datastudent {
	db, err := utils.Connect()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	data, err := db.Query("UPDATE student SET first_name=?, last_name=?, email=? WHERE id=?", first_name, last_name, email, id)

	if err != nil {
		log.Fatal(err)
	}

	defer data.Close()

	var result2 []Datastudent

	for data.Next() {
		var siswa = Datastudent{}
		var err = data.Scan(&siswa.Id, &siswa.Firstname, &siswa.Lastname, &siswa.Email)

		if err != nil {
			log.Fatal(err)
		}

		result2 = append(result2, siswa)
	}

	if err = data.Err(); err != nil {
		log.Fatal(err)
	}

	return result2

}

func DeleteStudent(id string) {
	db, err := utils.Connect()

	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	tx, err := db.Begin()
	stmt, err := tx.Exec("DELETE FROM student WHERE id=?", id)
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
