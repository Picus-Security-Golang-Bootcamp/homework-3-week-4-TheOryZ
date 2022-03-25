package service

import (
	model "Picus-Security-Golang-Bootcamp/homework-3-week-4-TheOryZ/pkg/model"
	"encoding/json"
	"io/ioutil"
	"os"
)

// GetAllBooks Function that returns all books
func GetAllBooks() (*model.Books, error) {
	var books model.Books
	jsonFile, err := os.Open("../../pkg/docs/books.json")
	if err != nil {
		return nil, err
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &books)
	defer jsonFile.Close()
	return &books, nil
}

//GetAllAuthors Function that returns all authors
func GetAllAuthors() (*model.Authors, error) {
	var authors model.Authors
	jsonFile, err := os.Open("../../pkg/docs/authors.json")
	if err != nil {
		return nil, err
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &authors)
	defer jsonFile.Close()
	return &authors, nil
}
