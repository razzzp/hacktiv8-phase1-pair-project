package main

import (
	"log"
	"roc-gameshop-app/config"
)

func main() {
	config.InitGoDotEnv()

	db, err := config.CreateDBInstance()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Println("Connect to DB.")
}
