package main

import (
	"log"
	"warehouse-management/config"
	"warehouse-management/routes"
)

func main() {
	// Inisialisasi Database
	config.InitDB()

	// Inisialisasi Router
	router := routes.SetupRouter()

	// Menjalankan Server
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Gagal menjalankan server:", err)
	}
}
