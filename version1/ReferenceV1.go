package version1

type ReferenceV1 struct {
	Id   string `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
}

func NewReferenceV1(id, typee, name string) *ReferenceV1 {
	return &ReferenceV1{
		Id:   id,
		Type: typee,
		Name: name,
	}
}
