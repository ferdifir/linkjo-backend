package validators

type ProductRequest struct {
	TenantID    uint    `json:"tenant_id" validate:"required"`
	CategoryID  *uint   `json:"category_id"`
	Name        string  `json:"name" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
	Stock       int     `json:"stock" validate:"required"`
	Unit        string  `json:"unit" validate:"required"`
	Image       string  `json:"image"`
	Description string  `json:"description"`
}
