package interfaces

import "libraryControlApp_backend/commonStructs"

type LibraryUsecase interface {
	Search(input string) string
	Filter(input commonStructs.Filter) string
	AddBook(input commonStructs.BookInput) error
	AddAuthor(input string) error
	AddType(input string) error
	AddPublisher(input string) error
}
