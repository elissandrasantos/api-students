package db

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite" // Sqlite driver based on CGO
	// "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name   string `json:"name"`
	CPF    int    `json:"cpf"`
	Email  string `json:"email"`
	Age    int    `json:"age"`
	Active bool   `json:"registration"`
}

func Init() *gorm.DB {

	// github.com/mattn/go-sqlite3
	db, err := gorm.Open(sqlite.Open("student.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&Student{})

	return db
}

func AddStudent(student Student) error {
	db := Init() //quando está dentro do pacote ja pode chamar direto pelo nome, por isso não usei db.Init

	//result := db.Create(&student)
	//if result.Error != nil {
	//	fmt.Println("Error to create student")
	//}

	if result := db.Create(&student); result.Error != nil {
		return result.Error
	}

	fmt.Println("Create student!")
	return nil
}
