package Models

type Author struct {
	Id   int64
	Name string
}

func CreateAuthor(id int64, name string) Author {
	a := Author{
		Id:   id,
		Name: name,
	}
	return a
}

func (a Author) Print() {
	println("Id: ", a.Id, " name: ", a.Name)
}
