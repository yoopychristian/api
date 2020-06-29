package service

import (
	"log"

	"challenge6.2/utils"
)

type Subject struct {
	Id     string `json:"id"`
	Matpel string `json:"matpel"`
}

func Subjectsekolah() []Subject {
	db, err := utils.Connect()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	data, err := db.Query("select * from subject")

	if err != nil {
		log.Fatal(err)
	}

	defer data.Close()

	var result3 []Subject

	for data.Next() {
		var matpelbaru = Subject{}
		var err = data.Scan(&matpelbaru.Id, &matpelbaru.Matpel)

		if err != nil {
			log.Fatal(err)
		}

		result3 = append(result3, matpelbaru)
	}

	if err = data.Err(); err != nil {
		log.Fatal(err)
	}

	return result3

}

func FindSubject(Id string) ([]Subject, error) {
	db, err := utils.Connect()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	data, err := db.Query("SELECT * FROM subject WHERE id = ?", Id)

	if err != nil {
		log.Fatal(err)
	}

	defer data.Close()

	var result2 []Subject

	for data.Next() {
		var subject = Subject{}
		var err = data.Scan(&subject.Id, &subject.Matpel)

		if err != nil {
			log.Fatal(err)
		}

		result2 = append(result2, subject)
	}

	if err = data.Err(); err != nil {
		log.Fatal(err)
	}

	return result2, nil

}

func CreateSubject(Id, Matpel string) []Subject {
	db, err := utils.Connect()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	data, err := db.Query("INSERT INTO subject VALUES (?,?)", Id, Matpel)

	if err != nil {
		log.Fatal(err)
	}

	defer data.Close()

	var result2 []Subject

	for data.Next() {
		var siswa = Subject{}
		var err = data.Scan(&siswa.Id, &siswa.Matpel)

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

func UpdateSubject(Id, Matpel string) []Subject {
	db, err := utils.Connect()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	data, err := db.Query("UPDATE subject SET subject_name=? WHERE id=?", Matpel, Id)

	if err != nil {
		log.Fatal(err)
	}

	defer data.Close()

	var result2 []Subject

	for data.Next() {
		var siswa = Subject{}
		var err = data.Scan(&siswa.Id, &siswa.Matpel)

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

func DeleteSubject(id string) {
	db, err := utils.Connect()

	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	tx, err := db.Begin()
	stmt, err := tx.Exec("DELETE FROM subject WHERE id=?", id)
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
