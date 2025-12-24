package entity

import "time"

type Product struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	TokoID      uint      `json:"toko_id"` // Nanti diisi otomatis dari toko si user
	NamaProduct string    `json:"nama_product"`
	Harga       int       `json:"harga"`
	Stok        int       `json:"stok"`
	Deskripsi   string    `json:"deskripsi"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
