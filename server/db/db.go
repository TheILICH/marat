package db

import (
	"os"
	"fmt"
	"github.com/yerkebayev/go-final-go/models"
	"github.com/joho/godotenv"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {

	if err := godotenv.Load(); err != nil {
		fmt.Println("NO ENV FILE!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	}

	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")	
	hostname := os.Getenv("HOST")
	port := os.Getenv("PORT")
	schema := os.Getenv("SCHEMA")
	database := os.Getenv("POSTGRES_DB")


	DSN := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?search_path=%s", username, password, hostname, port, database, schema)

	db, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(&models.Teacher{}, &models.Student{}, &models.Attendance{}, &models.Course{}, &models.Session{}, &models.Image{})

	return db
}
