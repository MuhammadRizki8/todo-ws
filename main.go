package main

import (
	"todo-service/config"
	"todo-service/models"
	"todo-service/routes"
)

func main() {
    // Koneksi ke database
    config.ConnectDB()

    // Migrasi model ke database
    models.Migrate(config.DB)
	// seeder data dummy
	models.SeedTodos(config.DB)

    // Inisialisasi router
    r := routes.SetupRoutes()

    // Jalankan server
    r.Run(":8080")
}
