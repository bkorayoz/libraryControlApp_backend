package Models

type Type struct {
	Id   int64
	Name string
}

func CreateType(id int64, name string) Type {
	t := Type{
		Id:   id,
		Name: name,
	}
	return t
}

func (t Type) Print() {
	println("Id: ", t.Id, " name: ", t.Name)
}
