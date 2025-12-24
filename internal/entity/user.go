package entity

import "time"

type User struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Nama        string    `json:"nama"`
	Email       string    `gorm:"unique;not null" json:"email"`        // Ketentuan: Email unik
	PhoneNumber string    `gorm:"unique;not null" json:"phone_number"` // Ketentuan: No Telp unik
	Password    string    `json:"password"`                            // Tidak ditampilkan di JSON
	IsAdmin     bool      `gorm:"default:false" json:"is_admin"`       // Untuk kelola kategori
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Toko struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"not null" json:"user_id"` // Foreign Key ke User
	NamaToko  string    `json:"nama_toko"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
