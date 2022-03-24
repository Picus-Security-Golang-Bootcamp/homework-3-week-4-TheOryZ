package model

type BookWithAuthor struct {
	ID             int
	Title          string
	NumberOfPages  int
	NumberOfStocks int
	Price          float64
	ISBN           string
	ReleaseDate    string
	AuthorID       int
	AuthorName     string
}
