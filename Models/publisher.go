package Models

type Publisher struct {
	Id   int64
	Name string
}

func CreatePublisher(id int64, name string) Publisher {
	p := Publisher{
		Id:   id,
		Name: name,
	}
	return p
}

func (p Publisher) Print() {
	println("Id: ", p.Id, " name: ", p.Name)
}
