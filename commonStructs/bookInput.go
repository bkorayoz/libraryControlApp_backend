package commonStructs

type BookInput struct {
	Name        string `json:name`
	AuthorId    int64  `json:authorid`
	TypeId      int64  `json:typeid`
	PublishDate string `json:publishdate`
	PublisherId int64  `json:publisherid`
	LocationId  int64  `json:locationid`
}
