# Final Task Evermoss - Backend Developer

Aplikasi E-Commerce sederhana berbasis RESTful API yang dibangun menggunakan **Golang**. Project ini mencakup manajemen user, toko, produk, alamat, dan sistem transaksi yang terintegrasi.

## 🚀 Fitur & Poin Penilaian
- **Auth & Shop System**: Registrasi user otomatis membuat toko baru [Poin 1-4].
- **Address Management**: Menambahkan alamat pengiriman untuk transaksi [Poin 5].
- **Product Management**: CRUD Produk terhubung ke toko user [Poin 9].
- **Transaction System**: Checkout produk dengan hitung otomatis & potong stok [Poin 10].

## 📸 Dokumentasi (Screenshot Postman)

### 1. Registrasi & Login (Poin 1-4)
User berhasil mendaftar dan mendapatkan JWT Token untuk akses fitur terproteksi.
> ![Register](assets/register_success.jpeg)
> ![Login](assets/login_success.jpg)

### 2. Tambah Alamat & Produk (Poin 5 & 9)
Setiap user bisa mendaftarkan alamat pengiriman dan menambahkan produk ke toko mereka.
> ![Add Address](assets/address_success.jpeg)
> ![Add Product](assets/product_success.jpeg)

### 3. Transaksi / Checkout (Poin 10)
Proses checkout otomatis menghitung `total_price` dan mengubah status menjadi `SUCCESS`.
> ![Checkout](assets/checkout_success.jpeg)

## 🛠️ Struktur Project (Dependency Injection)
```text
.
├── cmd/                # Entry Point (main.go)
├── config/             # DB Connection & Migration
├── internal/           # Core Logic
│   ├── entity/         # Database Models
│   ├── handler/        # Controller/HTTP Handler
│   ├── repository/     # Data Access Layer
│   ├── usecase/        # Business Logic Layer
│   ├── middleware/     # JWT & Admin Security
│   └── utils/          # JWT Helper
└── go.mod
