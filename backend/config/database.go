package config

import (
	"culturstone/models"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("../culturstone.db"), &gorm.Config{})
	if err != nil {
		panic("Gagal koneksi database")
	}

	// AutoMigrate will create/update tables based on the models
	err = DB.AutoMigrate(&models.Admin{}, &models.Category{}, &models.ProductCategory{}, &models.Product{}, &models.Portfolio{}, &models.Testimoni{}, &models.Message{}, &models.Visitor{}, &models.User{})
	if err != nil {
		panic("Failed to auto migrate database: " + err.Error())
	}

	seedData()
}

func seedData() {
	var c int64
	DB.Model(&models.Product{}).Count(&c)
	if c == 0 {
		DB.Create(&models.ProductCategory{Name: "Sanitary"})
		DB.Create(&models.ProductCategory{Name: "Dekorasi"})
		DB.Create(&models.ProductCategory{Name: "Eksterior"})

		cats := []models.Category{
			{Name: "Granit", Description: "Batu keras.", ImageURL: "https://images.unsplash.com/photo-1620302247770-c9a933e2594e?w=500"},
			{Name: "Marmer", Description: "Mewah.", ImageURL: "https://images.unsplash.com/photo-1617791160505-6f00504e35d9?w=500"},
		}
		DB.Create(&cats)

		DB.Create(&models.Product{
			Name: "Wastafel Marmer", Category: "Sanitary",
			ImageURL: "https://images.unsplash.com/photo-1584622050111-993a426fbf0a?w=500&q=80",
			Quality:  "Export", Size: "40cm", Finishing: "Polish",
			Description: "Wastafel batu alam minimalis.",
		})
	}
}
