package postgresql

import (
	"libraryControlApp_backend/Models"
	. "libraryControlApp_backend/util"
)

type LibraryData struct{}

var libraryData *LibraryData

func (LibraryData) GetAllBooks() []Models.Book { // FIX ME
	rows, err := DB.Query("SELECT * FROM books")
	CheckErr(err)
	defer rows.Close()

	var (
		id          int64
		name        string
		authorId    int64
		typeId      int64
		publishDate string
		publisherId int64
		locationId  int64
	)

	books := make([]Models.Book, 0)
	for rows.Next() {
		err := rows.Scan(&id, &name, &authorId, &typeId, &publishDate, &publisherId, &locationId)
		CheckErr(err)
		books = append(books, Models.CreateBook(id, name, authorId, typeId, publishDate, publisherId, locationId))
	}

	return books

}

func (LibraryData) GetBookById(input int64) {

}

func (LibraryData) GetBookByName(input string) {

}

func (LibraryData) GetBooksByAuthor(input string) {

}

func (LibraryData) GetBooksByType(input string) {

}

func (LibraryData) GetBooksByDate(start, end string) {

}

func (LibraryData) GetBooksByLocation(input string) {

}

func (LibraryData) AddBook(name string, authorId int64, typeId int64, publishDate string, locationId int64, publisherId int64) error {
	println("AddBook")
	return nil
}

func (LibraryData) DeleteBook(id int64) {

}

func (LibraryData) GetAllLocations() {

}

func (LibraryData) GetLocationById(input int64) {

}

func (LibraryData) AddLocation(blok string, shelf string) {

}

func (LibraryData) RemoveLocation(input int64) {

}

func (LibraryData) GetAll(tableName string) {

}

func (LibraryData) GetById(tableName string, id int64) {

}

func (LibraryData) Add(tableName string, name string) error {
	println("Add -> ", tableName, " ", name)
	return nil
}

func (LibraryData) Remove(tableName string, id int64) error {
	println("Remove -> ", tableName, " ", id)
	return nil
}
