package main

import (
	"fmt"
	"github.com/marcos-nsantos/e-commerce/auth-service/database"
	"github.com/marcos-nsantos/e-commerce/auth-service/routes"
	"log"
	"net/http"
)

const webPort = "80"

func init() {
	if err := database.AutoMigrateUser(); err != nil {
		log.Panicf("Fail to migrate database: %v", err)
	}
}

func main() {
	log.Println("Starting authentication service")

	r := routes.HandleRequests()
	if err := http.ListenAndServe(fmt.Sprintf(":%s", webPort), r); err != nil {
		log.Panic(err)
	}
}
