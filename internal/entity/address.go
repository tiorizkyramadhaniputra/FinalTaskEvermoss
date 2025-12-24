package entity

import "time"

type Address struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	UserID       uint      `gorm:"not null" json:"user_id"` // Poin 12: Untuk validasi kepemilikan
	NamaPenerima string    `json:"nama_penerima"`           // Standar alamat kirim
	NoTelepon    string    `json:"no_telepon"`              // Standar alamat kirim
	ProvinsiID   string    `json:"provinsi_id"`             // Dari API Wilayah
	KotaID       string    `json:"kota_id"`                 // Dari API Wilayah
	KecamatanID  string    `json:"kecamatan_id"`            // Dari API Wilayah
	KelurahanID  string    `json:"kelurahan_id"`            // Dari API Wilayah
	DetailAlamat string    `json:"detail_alamat"`           // Detail jalan/nomor rumah
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
