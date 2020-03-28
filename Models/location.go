package Models

type Location struct {
	Id    int64
	Blok  int
	Shelf int
}

func CreateLocation(id int64, blok int, shelf int) Location {
	l := Location{
		Id:    id,
		Blok:  blok,
		Shelf: shelf,
	}
	return l
}

func (l Location) Print() {
	println("Id: ", l.Id, " blok: ", l.Blok, " shelf: ", l.Shelf)
}
