package commonStructs

type Filter struct {
	Author    []string `json:"author"`
	Type      []string `json:"type"`
	Date      []string `json:"date"`
	Publisher []string `json:"publisher"`
}
