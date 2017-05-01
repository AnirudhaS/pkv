package main

import "fmt"
import "github.com/jinzhu/gorm"
import _ "github.com/jinzhu/gorm/dialects/postgres"

func connectDB() *gorm.DB {

	db, err := gorm.Open("postgres", "host=localhost user=postgres database=pkv sslmode=disable password=password")
	if err != nil {
		panic("failed to connect database")
	}
	if err != nil {
		fmt.Print(err)
		defer db.Close()
	}
	return db
}

func migrate() *gorm.DB {
	db := connectDB()
	db.AutoMigrate(&KeyValue{})
	return db
}
