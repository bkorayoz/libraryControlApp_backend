package interactors

type LibraryInteractor struct{}

func (LibraryInteractor) GetAllBooks() string {
	return "GetAllBooks()"
}

func (LibraryInteractor) GetBookById(input int64) string {
	return "GetBookById"
}

func (LibraryInteractor) GetBookByName(input string) string {
	return "GetBookByName"
}

func (LibraryInteractor) GetBooksByAuthor(input string) string {
	return "GetBooksByAuthor"
}

func (LibraryInteractor) GetBooksByType(input string) {

}

func (LibraryInteractor) GetBooksByDate(start, end string) {

}

func (LibraryInteractor) GetBooksByLocation(input string) {

}

func (LibraryInteractor) AddBook(name string, authorId int64, typeId int64, publishDate string, locationId int64, publisherId int64) {

}

func (LibraryInteractor) DeleteBook(id int64) {

}

func (LibraryInteractor) GetAllLocations() {

}

func (LibraryInteractor) GetLocationById(input int64) {

}

func (LibraryInteractor) AddLocation(blok string, shelf string) {

}

func (LibraryInteractor) RemoveLocation(input int64) {

}

func (LibraryInteractor) GetAll(tableName string) {

}

func (LibraryInteractor) GetById(tableName string, id int64) {

}

func (LibraryInteractor) Add(tableName string, name string) {

}

func (LibraryInteractor) Remove(tableName string, id int64) {

}
