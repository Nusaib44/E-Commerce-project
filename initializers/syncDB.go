package initializers

import (
	"project/pkg/models"
)

func SyncDB() {
	DB.AutoMigrate(&models.Userdata{})
	DB.AutoMigrate(&models.Product{})
	DB.AutoMigrate(&models.Category{})
	DB.AutoMigrate(models.Cart{})
	DB.AutoMigrate(models.ShoppingCartItem{})
	DB.AutoMigrate(models.Address{})
	DB.AutoMigrate(models.Payment{})
	DB.AutoMigrate(models.Order{})
	DB.AutoMigrate(models.Charge{})
	DB.AutoMigrate(models.Coupen{})
	DB.AutoMigrate(models.ProductOffers{})
	DB.AutoMigrate(models.CategoryOffers{})
	DB.AutoMigrate(models.Return{})
	DB.AutoMigrate(models.Walet{})
}
