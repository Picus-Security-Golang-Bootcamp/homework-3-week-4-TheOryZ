package book

import (
	"gorm.io/gorm"

	model "Picus-Security-Golang-Bootcamp/homework-3-week-4-TheOryZ/pkg/model"
)

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
	b.db.Where("ID = ?", id).Find(&book)
	return book
}

//FindByAuthorID Get by Author ID <SELECT * FROM Books WHERE AuthorID = authorID ORDER BY ID DESC>
func (b *BookRepository) FindByAuthorID(authorID int) []Book {
	var books []Book
	b.db.Where("AuthorID = ?", authorID).Order("Id desc").Find(&books)
	return books
}

//FindByTitle Get by Title <SELECT * FROM Books WHERE Title = title>
func (b *BookRepository) FindByTitle(title string) []Book {
	var books []Book
	b.db.Where("Title = ?", title).Find(&books)
	return books
}

//GetNonDeleted Get non deleted books
func (b *BookRepository) GetNonDeleted() []Book {
	var books []Book
	b.db.Where("deleted_at = ?", 0).Find(&books)
	return books
}

//GetByIdWithAuthorName Get all books and with authors names
func (b *BookRepository) GetByIdWithAuthorName() model.BookWithAuthor {
	var model model.BookWithAuthor
	b.db.Select("Books.id, Books.title, Books.number_of_pages, Books.number_of_stocks, Books.price, Books.isbn, Books.release_date, Books.author_id, Authors.name").Joins("left join Authors on Authors.id = Books.author_id").Scan(model)
	return model
}

//Insert Create new Book
func (b *BookRepository) Insert(book Book) error {
	result := b.db.Create(book)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//Update Update book
func (b *BookRepository) Update(book Book) error {
	result := b.db.Save(book)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//Delete Delete book
func (b *BookRepository) Delete(book Book) error {
	result := b.db.Delete(book)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
