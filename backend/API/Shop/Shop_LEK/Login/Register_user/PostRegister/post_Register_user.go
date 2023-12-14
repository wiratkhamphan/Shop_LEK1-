// PostRegister/post_Register_user.go
package PostRegister

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IOSPerson struct {
	Id       int    `json:"id"`
	IsoPCode string `json:"iso_p_code"`
	PCode    string `json:"p_code"`
	Email    string `json:"email"`
	Password string `json:"password"`
	PName    string `json:"p_name"`
}

var DB *sql.DB
var persons []IOSPerson

// RegisterUser should now be a gin.HandlerFunc
func RegisterUser(c *gin.Context) {
	var P IOSPerson
	if err := c.BindJSON(&P); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash the password before storing it in the database
	hashedPassword := hashPassword(P.Password)

	// Assuming you have a variable DB of type *sql.DB representing your database connection
	_, err := DB.Exec("INSERT INTO ios_person (iso_p_code, p_code, email, password, p_name) VALUES (?, ?, ?, ?, ?)",
		P.IsoPCode, P.PCode, P.Email, hashedPassword, P.PName)

	if err != nil {
		log.Println("Failed to insert data into the database:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert data into the database"})
		return
	}
	// Append the new user to the slice
	persons = append(persons, P)

	// Respond with the created user
	c.JSON(http.StatusCreated, P)
}

// Replace this with your actual password hashing logic
func hashPassword(password string) string {
	// Implement password hashing logic here
	// For example, you can use a library like bcrypt
	// Don't store passwords as plain text in the database
	return password
}
