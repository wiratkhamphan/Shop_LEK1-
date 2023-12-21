package PostRegister

import (
	dbtypeiosperson "SHOP_LEK/API/Shop/Shop_LEK/Login/DbtypeIosPerson"
	condb "SHOP_LEK/database/condb"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var db *sql.DB

func PostAlbums(c *gin.Context) {
	var err error

	// Use the existing db variable from the package scope
	db, err = condb.NewConndb_mysql()
	if err != nil {
		log.Fatal(err)
	}

	var newAlbum dbtypeiosperson.IOSPerson

	if err := c.BindJSON(&newAlbum); err != nil {
		fmt.Println("Error binding JSON:", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	fmt.Println("Received Data:", newAlbum)
	_, err = db.Exec("INSERT INTO ios_person (iso_p_code, p_code, email, password, p_name) VALUES (?, ?,?, ?,?)",
		newAlbum.Iso_p_code, newAlbum.P_code, newAlbum.Email, newAlbum.Password, newAlbum.P_name)
	if err != nil {
		fmt.Println("Error inserting into database:", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert data into the database"})
		return
	}

	c.IndentedJSON(http.StatusCreated, newAlbum)
}
