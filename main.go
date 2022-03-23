package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	//Set enviroment variables
	err := godotenv.Load("pkg/env/connection.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println(os.Getenv("DB_NAME"))
}
