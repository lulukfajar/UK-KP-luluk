package main

import (
	"UjianKetrampilan/db"
	"UjianKetrampilan/models"
	"UjianKetrampilan/routers"
	"log"
)

func init() {
	db.InitializeDB()
}

func main() {
	pg := db.GetDB()

	err := pg.AutoMigrate(models.User{}, models.Photo{}, models.SocialMedia{}, models.Comment{})

	if err != nil {
		log.Fatalf("pg.AutoMigrate: %s\n", err.Error())
	}

	router := routers.StartServer()
	router.Run()
}
