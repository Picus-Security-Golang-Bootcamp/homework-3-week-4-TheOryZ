package main

import (
	"Picus-Security-Golang-Bootcamp/homework-3-week-4-TheOryZ/pkg/store/author"
	"Picus-Security-Golang-Bootcamp/homework-3-week-4-TheOryZ/pkg/store/book"
	postgres "Picus-Security-Golang-Bootcamp/homework-3-week-4-TheOryZ/pkg/store/common/db"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	//Set enviroment variables
	err := godotenv.Load("pkg/env/connection.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db, err := postgres.NewPsqlDB()
	if err != nil {
		log.Fatal("Postgres cannot init", err)
	}
	log.Println("Postgres connected")

	//Connection DB and migrations
	authorRepo := author.NewAuthorRepository(db)
	authorRepo.Migrations()

	bookRepo := book.NewBookRepository(db)
	bookRepo.Migrations()

	//Insert Seed Datas
	authorRepo.InsertSeedData()
	bookRepo.InsertSeedData()
}
