package author

import (
	model "Picus-Security-Golang-Bootcamp/homework-3-week-4-TheOryZ/pkg/model"

	"gorm.io/gorm"
)

type AuthorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) *AuthorRepository {
	return &AuthorRepository{db: db}
}
func (a *AuthorRepository) Migrations() {
	a.db.AutoMigrate(&Author{})
}

//FindAll Get all records <SELECT * FROM Authors>
func (a *AuthorRepository) FindAll() []Author {
	var authors []Author
	a.db.Find(&authors)
	return authors
}

//FindById Get By Id <SELECT * FROM Authors WHERE ID = id>
func (a *AuthorRepository) FindById(id int) Author {
	var author Author
	a.db.Where("id = ?", id).Find(&author)
	return author
}

//FindByName Get by Name <SELECT * FROM Authors WHERE name = name>
func (a *AuthorRepository) FindByName(name string) []Author {
	var authors []Author
	a.db.Where("name = ?", name).Find(&authors)
	return authors
}

//GetNonDeleted Get non deleted authors
func (a *AuthorRepository) GetNonDeleted() []Author {
	var authors []Author
	a.db.Where("deleted_at = ?", 0).Find(&authors)
	return authors
}

//GetByIdWithBooks Get Author by id and with books
func (a *AuthorRepository) GetByIdWithBooks(id int) model.BookWithAuthor {
	var model model.BookWithAuthor
	a.db.Select("Authors.id, Authors.name, Books.id ,Books.title").Joins("left join Books on Authors.id = Books.author_id").Where("Authors.id = ?", id).Scan(model)
	return model
}

//Insert Create new Author
func (a *AuthorRepository) Insert(author Author) error {
	result := a.db.Create(author)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//Update Update book
func (a *AuthorRepository) Update(author Author) error {
	result := a.db.Save(author)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//Delete Delete author
func (a *AuthorRepository) Delete(author Author) error {
	result := a.db.Delete(author)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//DeleteById Delete by Id author
func (a *AuthorRepository) DeleteById(id int) error {
	result := a.db.Delete(&Author{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
