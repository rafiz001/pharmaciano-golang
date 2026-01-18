package initializers

import "pharmaciano/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
