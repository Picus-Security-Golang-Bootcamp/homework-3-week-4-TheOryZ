package main

import (
	"Picus-Security-Golang-Bootcamp/homework-3-week-4-TheOryZ/pkg/store/author"
	"Picus-Security-Golang-Bootcamp/homework-3-week-4-TheOryZ/pkg/store/book"
	postgres "Picus-Security-Golang-Bootcamp/homework-3-week-4-TheOryZ/pkg/store/common/db"
	"fmt"
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
	log.Println("Seed Datas inserted")

	//**Sample Queries for Authors

	//*Find all authors
	fmt.Println(authorRepo.FindAll())

	//*Find by name (search)
	//fmt.Println(authorRepo.FindByName("Gene"))

	//*Get author name by id with books
	//fmt.Println(authorRepo.GetByIdWithBooks(3))

	//*insert new author
	// model := author.Author{Name: "Gene Wolfe"}
	// fmt.Println(authorRepo.Insert(model))

	//**Sample Queries for Books

	//*Find all books
	fmt.Println(bookRepo.FindAll())

	//*Find by title (search)
	//fmt.Println(bookRepo.FindByTitle("Dune"))

	//*Get Book by id with author name
	//fmt.Println(bookRepo.GetByIdWithAuthorName(1))

	//*insert new book
	// model := book.Book{Title: "Animal Farm", NumberOfPages: 112, NumberOfStocks: 65, Price: 11, ISBN: "ISBN987654321", ReleaseDate: "1945", AuthorID: 3}
	// fmt.Println(bookRepo.Insert(model))
}
