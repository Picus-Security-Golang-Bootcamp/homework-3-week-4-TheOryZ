package book

import "gorm.io/gorm"

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}
func (b *BookRepository) Migrations() {
	b.db.AutoMigrate(&Book{})
}

//FindAll Get all records <SELECT * FROM Books>
func (b *BookRepository) FindAll() []Book {
	var books []Book
	b.db.Find(&books)
	return books
}

//FindById Get By Id <SELECT * FROM Books WHERE ID = id>
func (b *BookRepository) FindById(id int) Book {
	var book Book
	b.db.Where(`"ID" = ?`, id).Find(&book)
	return book
}

//FindByAuthorID Get by Author ID <SELECT * FROM Books WHERE AuthorID = authorID ORDER BY ID DESC>
func (b *BookRepository) FindByAuthorID(authorID int) []Book {
	var books []Book
	b.db.Where(`"AuthorID" = ?`, authorID).Order("Id desc").Find(&books)
	return books
}
