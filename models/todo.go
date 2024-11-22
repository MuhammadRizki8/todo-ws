package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID        uint      `gorm:"primaryKey"`
	Title     string    `gorm:"type:varchar(255);not null"`
	Note      string    `gorm:"type:varchar(500)"`
	Completed bool      `gorm:"default:false"`
	Deadline  time.Time 
	CreatedAt time.Time `gorm:"autoUpdateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

// Fungsi untuk menjalankan migrasi
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Todo{})
}

// Fungsi untuk seeding data awal
func SeedTodos(db *gorm.DB) {
	todos := []Todo{
		{
			Title:     "Belajar golang",
			Note:      "Mulai dengan dasar-dasar golang",
			Deadline:  time.Now().AddDate(0, 0, 7), // Tenggat waktu 7 hari dari sekarang
			Completed: false,
		},
		{
			Title:     "Belajar ORM di golang pakai GORM",
			Note:      "Pelajari migrasi, seeder, dan CRUD",
			Deadline:  time.Now().AddDate(0, 0, 10), // Tenggat waktu 10 hari dari sekarang
			Completed: false,
		},
		{
			Title:     "Implementasi Seeder berhasil!",
			Note:      "Pastikan seeder tidak dijalankan di production",
			Deadline:  time.Now().AddDate(0, 0, 1), // Tenggat waktu 1 hari dari sekarang
			Completed: true,
		},
	}
	for _, todo := range todos {
		db.FirstOrCreate(&todo, Todo{Title: todo.Title})
	}
}
