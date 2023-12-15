package shoplek

import (
	PostRegister "SHOP_LEK/API/Shop/Shop_LEK/Login/Register_user/PostRegister"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	condb "SHOP_LEK/database/condb"

	"github.com/gin-gonic/gin"
)

var db *sql.DB

type IOSPerson struct {
	ID         string `json:"id"`
	Iso_p_code string `json:"iso_p_code"`
	P_code     string `json:"p_code"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	P_name     string `json:"p_name"`
}

func RunLocalhost() {
	var err error
	dbInstance, err := condb.NewConndb_mysql()
	// fmt.Println(dbInstance)
	if err != nil {
		log.Fatal(err)
	}

	db = dbInstance

	router := gin.Default()
	// router.POST("/albums", postAlbums1)
	router.POST("/Register_user", PostRegister.PostAlbums)
	router.POST("/Register_user1", postAlbums)

	router.Run("localhost:8080")
}

func postAlbums(c *gin.Context) {
	var newAlbum IOSPerson

	if err := c.BindJSON(&newAlbum); err != nil {
		fmt.Println("Error binding JSON:", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	fmt.Println("Received Data:", newAlbum)

	_, err := db.Exec("INSERT INTO ios_person (iso_p_code, p_code, email, password, p_name) VALUES (?, ?,?, ?,?)",
		newAlbum.Iso_p_code, newAlbum.P_code, newAlbum.Email, newAlbum.Password, newAlbum.P_name)

	if err != nil {
		fmt.Println("Error inserting into database:", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert data into the database"})
		return
	}

	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// func postAlbums1(c *gin.Context) {
// 	var newAlbum IOSPerson

// 	if err := c.BindJSON(&newAlbum); err != nil {
// 		fmt.Println("Error binding JSON:", err)
// 		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
// 		return
// 	}

// 	fmt.Println("Received Data:", newAlbum)

// 	// Corrected SQL query, removed extra comma
// 	// INSERT INTO `ios_person` (`id`, `iso_p_code`, `p_code`, `email`, `password`, `p_name`) VALUES (NULL, '4', '44', '4', '44', '44');
// 	_, err := db.Exec("INSERT INTO ios_person (iso_p_code, p_code, email, password, p_name) VALUES (?, ?,?, ?,?)",
// 		newAlbum.Iso_p_code, newAlbum.P_code, newAlbum.Email, newAlbum.Password, newAlbum.P_name)

// 	if err != nil {
// 		fmt.Println("Error inserting into database:", err)
// 		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert data into the database"})
// 		return
// 	}

// 	c.IndentedJSON(http.StatusCreated, newAlbum)
// }
