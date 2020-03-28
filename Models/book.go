package Models

type Book struct {
	Id          int64
	Name        string
	AuthorId    int64
	TypeId      int64
	PublishDate string
	PublisherId int64
	LocationId  int64
}

func CreateBook(id int64, name string, authorId int64, typeId int64, publishDate string, publisherId int64, locationId int64) Book {
	b := Book{
		Id:          id,
		Name:        name,
		AuthorId:    authorId,
		TypeId:      typeId,
		PublishDate: publishDate,
		PublisherId: publisherId,
		LocationId:  locationId,
	}
	return b
}

func (b Book) Print() {
	println("Id: ", b.Id, " name: ", b.Name, " author: ", b.AuthorId, " type: ", b.TypeId, " publisher: ", b.PublisherId, " publish date: ", b.PublishDate)

}
