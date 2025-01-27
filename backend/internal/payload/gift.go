package payload

type GiftRequest struct {
	Description string `json:"description" validate:"required"`
}

type GiftResponse struct {
	Name        string `json:"gift_name"`
	Description string `json:"gift_description"`
	Price       int    `json:"price"`
	ImageURL    string `json:"image_url"`
}
