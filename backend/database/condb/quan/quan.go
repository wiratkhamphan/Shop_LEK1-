package quan

import (
	"database/sql"
	"fmt"
	"log"
)

type ios_person struct {
	Id         int
	iso_p_code string
	p_code     string
	email      string
	password   string
	p_name     string
}

func GetUser(db *sql.DB, id int) (ios_person, error) {
	var U ios_person
	row := db.QueryRow(`SELECT id, username, password FROM users WHERE id = ?;`, id)
	err := row.Scan(&U.Id, &U.iso_p_code, &U.password, &U.p_code, U.p_name, U.email)
	if err != nil {
		return ios_person{}, err
	}
	return U, nil
}

func Shop(db *sql.DB) {
	ID, err := GetUser(db, 1)
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println("print ID", ID)
}
