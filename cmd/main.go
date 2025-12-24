package main

import (
	"fmt"
	"log"
	"net/http"

	"FinalTaskEvermoss/config"
	"FinalTaskEvermoss/internal/handler"
	"FinalTaskEvermoss/internal/middleware"
	"FinalTaskEvermoss/internal/repository"
	"FinalTaskEvermoss/internal/usecase"

	"github.com/gorilla/mux"
)

func main() {
	// 1. Inisialisasi Database
	config.ConnectDB()

	// 2. Dependency Injection — USER
	userRepo := repository.NewUserRepository(config.DB)
	userUsc := usecase.NewUserUsecase(userRepo)
	userHdl := handler.NewUserHandler(userUsc)

	// 3. Dependency Injection — ADDRESS
	addressRepo := repository.NewAddressRepository(config.DB)
	addressUsc := usecase.NewAddressUsecase(addressRepo)
	addressHdl := handler.NewAddressHandler(addressUsc)

	// 4. Dependency Injection — PRODUCT
	productRepo := repository.NewProductRepository(config.DB)
	productUsc := usecase.NewProductUsecase(productRepo)
	productHdl := handler.NewProductHandler(productUsc)

	// 5. Dependency Injection — TRANSACTION (BARU: Sikat Poin 10)
	txRepo := repository.NewTransactionRepository(config.DB)
	txUsc := usecase.NewTransactionUsecase(txRepo, productRepo)
	txHdl := handler.NewTransactionHandler(txUsc)

	// 6. Router utama
	r := mux.NewRouter()

	// 7. API Grouping
	api := r.PathPrefix("/api/v1").Subrouter()

	// ===== PUBLIC AUTH =====
	api.HandleFunc("/register", userHdl.Register).Methods("POST")
	api.HandleFunc("/login", userHdl.Login).Methods("POST")

	// ===== PUBLIC PRODUCTS (Poin 9: Bisa dilihat tanpa login) =====
	api.HandleFunc("/products", productHdl.List).Methods("GET")

	// ===== PROTECTED ROUTES (Butuh Token) =====
	protected := api.PathPrefix("").Subrouter()
	protected.Use(middleware.AuthMiddleware)

	// ADDRESS (Poin 5)
	protected.HandleFunc("/addresses", addressHdl.Create).Methods("POST")
	protected.HandleFunc("/addresses", addressHdl.GetMyAddresses).Methods("GET")

	// PRODUCT (Poin 9: Tambah produk butuh login)
	protected.HandleFunc("/products", productHdl.Create).Methods("POST")

	// TRANSACTION (Poin 10: Checkout produk)
	protected.HandleFunc("/checkout", txHdl.Checkout).Methods("POST")

	fmt.Println("🚀 Server jalan di port 8080...")

	// 8. Run server
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
