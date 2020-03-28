package interactors

import (
	"libraryControlApp_backend/commonStructs"
	"libraryControlApp_backend/gateways/postgresql"
)

type LibraryInteractor struct{}

func (LibraryInteractor) Search(input string) string {
	return "Search"
}

func (LibraryInteractor) Filter(input commonStructs.Filter) string {
	return "Filter"
}

func (LibraryInteractor) AddBook(input commonStructs.BookInput) error {
	lib := postgresql.LibraryData{}
	return lib.AddBook(input.Name, input.AuthorId, input.TypeId, input.PublishDate, input.LocationId, input.PublisherId)
}

func (LibraryInteractor) AddAuthor(input string) error {
	lib := postgresql.LibraryData{}
	return lib.Add("authors", input)
}

func (LibraryInteractor) AddType(input string) error {
	lib := postgresql.LibraryData{}
	return lib.Add("types", input)
}

func (LibraryInteractor) AddPublisher(input string) error {
	lib := postgresql.LibraryData{}
	return lib.Add("publishers", input)
}
