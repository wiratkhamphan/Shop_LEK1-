package condb

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	DBDriver string
	DBUser   string
	DBPass   string
	DBName   string
}

func NewEnv() (*Env, error) {
	err := godotenv.Load("database/Env/.env")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	env := &Env{
		DBDriver: os.Getenv("DB_Driver"),
		DBUser:   os.Getenv("DB_User"),
		DBPass:   os.Getenv("DB_Pass"),
		DBName:   os.Getenv("DB_name"),
	}
	fmt.Println(env)

	return env, nil

}
