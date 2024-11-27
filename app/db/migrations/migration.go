package main

import (
    "go-restapi-boiltertemplate/app/infra"
    "go-restapi-boiltertemplate/app/models"
)

func main() {
	infra.Initializer()
	db := infra.SetupDB()

	if err := db.AutoMigrate(&models.Post{},&models.User{}); err != nil {
		panic("Failed to migrate database")
	}
}