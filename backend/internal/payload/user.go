package payload

type UserRequest struct {
	Age       int      `json:"age" validate:"required"`
	Sex       string   `json:"sex" validate:"required"`
	Interests []string `json:"interests" validate:"required"`
}
