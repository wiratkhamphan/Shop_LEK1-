package condb

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// ประกาศ type Env ไว้รับค่าจากไฟล์ mysql.env
type Env struct {
	DBDriver string
	DBUser   string
	DBPass   string
	DBName   string
}

// NewEnv หาที่อยู่ไฟล์ mysql.env
func NewEnv() (*Env, error) {
	err := godotenv.Load("database/Env/file_Env/mysql.env") // ระบุที่อยู่ไฟล์
	if err != nil {
		log.Fatal("Error loading .env file:", err) //  ถ้าเกิดมีค่า nil ให้ แสดง log err
	}
	//รับค่า เข้าที่ไปใน  type Env
	env := &Env{
		DBDriver: os.Getenv("DB_Driver"),
		DBUser:   os.Getenv("DB_User"),
		DBPass:   os.Getenv("DB_Pass"),
		DBName:   os.Getenv("DB_name"),
	}
	// fmt.Println(env)

	return env, nil

}
