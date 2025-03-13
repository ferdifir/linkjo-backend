package validators

type OrderRequest struct {
	CustomerName  string         `json:"customer_name,omitempty"`
	TableNumber   string         `json:"table_number,omitempty"`
	PaymentStatus string         `json:"payment_status" validate:"required,oneof=pending paid cancelled"`
	PaymentMethod string         `json:"payment_method" validate:"required,oneof=cash qris"`
	Products      []OrderProduct `json:"products" validate:"required,dive"`
}

type OrderProduct struct {
	ProductID int `json:"product_id" validate:"required"`
	Quantity  int `json:"quantity" validate:"required,min=1"`
}

type PaymentRequest struct {
	PaymentStatus string `json:"payment_status" validate:"required,oneof=pending paid cancelled"`
}
