package entity

import "time"

type Transaction struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	UserID     uint      `json:"user_id"`
	AddressID  uint      `json:"address_id"`
	ProductID  uint      `json:"product_id"`
	Quantity   int       `json:"quantity"`
	TotalPrice int       `json:"total_price"`
	Status     string    `json:"status"` // e.g., "PENDING", "SUCCESS"
	CreatedAt  time.Time `json:"created_at"`

	// Optional: Biar respon JSON-nya lengkap
	Product Product `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	Address Address `gorm:"foreignKey:AddressID" json:"address,omitempty"`
}
