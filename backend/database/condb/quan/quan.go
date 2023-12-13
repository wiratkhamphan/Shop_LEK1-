package quan

import "database/sql"

type User struct {
	Id       int
	Username string
	Password string
}

func GetUser(db *sql.DB, id int) (User, error) {
	var U User
	row := db.QueryRow(`SELECT id, username, password FROM users WHERE id = $1;`, id)
	err := row.Scan(&U.Id, &U.Username, &U.Password)
	if err != nil {
		return User{}, err
	}
	return U, nil
}
