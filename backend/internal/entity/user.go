package entity

type User struct {
	Age int    `json:"age"`
	Job string `json:"job"`
	Interests []string `json:"interests"`
}
