package validators

type ProductRequest struct {
	CategoryID  int     `json:"category_id" validate:"required"`
	Name        string  `json:"name" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
	Stock       int     `json:"stock" validate:"required"`
	Unit        string  `json:"unit" validate:"required"`
	Image       string  `json:"image"`
	Description string  `json:"description"`
}
