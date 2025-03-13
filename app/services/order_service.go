package services

import (
	"errors"
	"linkjo/app/models"
	"linkjo/app/validators"
	"time"

	"linkjo/config"
)

type OrderResponse struct {
	ID            uint    `json:"id"`
	CustomerName  string  `json:"customer_name"`
	TableNumber   string  `json:"table_number"`
	TotalPrice    float64 `json:"total_price"`
	PaymentStatus string  `json:"payment_status"`
	PaymentMethod string  `json:"payment_method"`
	CreatedAt     time.Time
	OrderDetails  []OrderDetailResponse
}

type OrderDetailResponse struct {
	OrderID      uint    `json:"order_id"`
	ProductID    uint    `json:"product_id"`
	ProductName  string  `json:"product_name"`
	ProductPrice float64 `json:"product_price"`
	Quantity     int     `json:"quantity"`
	Subtotal     float64 `json:"subtotal"`
}

func CreateOrder(order validators.OrderRequest, userID *uint) (*models.Order, error) {
	// Mulai transaksi database
	tx := config.DB.Begin()

	// Validasi apakah daftar produk ada
	if len(order.Products) == 0 {
		return nil, errors.New("produk tidak boleh kosong")
	}

	// Hitung total harga berdasarkan produk dan kuantitas
	var totalPrice float64
	var orderDetails []models.OrderDetail

	for _, item := range order.Products {
		var product models.Product
		if err := tx.Where("id = ?", item.ProductID).First(&product).Error; err != nil {
			tx.Rollback() // Batalkan transaksi jika produk tidak ditemukan
			return nil, errors.New("produk dengan ID " + string(rune(item.ProductID)) + " tidak ditemukan")
		}

		subtotal := float64(item.Quantity) * product.Price
		totalPrice += subtotal

		orderDetails = append(orderDetails, models.OrderDetail{
			ProductID: uint(item.ProductID),
			Quantity:  item.Quantity,
			Subtotal:  subtotal,
		})
	}

	// Simpan data Order ke database
	newOrder := models.Order{
		UserID:        *userID,
		CustomerName:  order.CustomerName,
		TableNumber:   order.TableNumber,
		TotalPrice:    totalPrice,
		PaymentStatus: order.PaymentStatus,
		PaymentMethod: order.PaymentMethod,
	}

	if err := tx.Create(&newOrder).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Tambahkan order_id ke setiap OrderDetail yang dibuat
	for i := range orderDetails {
		orderDetails[i].OrderID = newOrder.ID
	}

	// Simpan semua OrderDetail dalam batch insert
	if err := tx.Create(&orderDetails).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// kurangi stok produk
	for _, item := range order.Products {
		var product models.Product
		if err := tx.Where("id = ?", item.ProductID).First(&product).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
		product.Stock -= item.Quantity
		if err := tx.Save(&product).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	// Commit transaksi jika semua berhasil
	tx.Commit()

	return &newOrder, nil
}

func GetOrdersByUserID(userID *uint) ([]OrderResponse, error) {
	var orders []OrderResponse

	query := `
		SELECT o.id, o.customer_name, o.table_number, o.total_price, 
		       o.payment_status, o.payment_method, 
		       p.name AS product_name, od.quantity, od.subtotal, p.id AS product_id,
			   p.price AS product_price
		FROM orders o
		LEFT JOIN order_details od ON o.id = od.order_id
		LEFT JOIN products p ON od.product_id = p.id
		WHERE o.user_id = ?
		ORDER BY o.id;
	`

	rows, err := config.DB.Raw(query, userID).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Gunakan map untuk mengelompokkan order dan order_details
	orderMap := make(map[uint]*OrderResponse)

	for rows.Next() {
		var orderID uint
		var o OrderResponse
		var d OrderDetailResponse

		err = rows.Scan(&orderID, &o.CustomerName, &o.TableNumber, &o.TotalPrice,
			&o.PaymentStatus, &o.PaymentMethod, &d.ProductName, &d.Quantity, &d.Subtotal, &d.ProductID, &d.ProductPrice)
		if err != nil {
			return nil, err
		}

		// Jika order sudah ada di map, tambahkan detail produk
		if existingOrder, exists := orderMap[orderID]; exists {
			existingOrder.OrderDetails = append(existingOrder.OrderDetails, d)
		} else {
			// Buat order baru dan tambahkan detail produk pertama
			o.ID = orderID
			o.OrderDetails = []OrderDetailResponse{d}
			orderMap[orderID] = &o
		}
	}

	// Konversi map ke slice
	for _, order := range orderMap {
		orders = append(orders, *order)
	}

	return orders, nil
}

// update payment status
func UpdatePaymentStatus(orderID uint, paymentStatus string) error {
	return config.DB.Model(&models.Order{}).Where("id = ?", orderID).Update("payment_status", paymentStatus).Error
}
