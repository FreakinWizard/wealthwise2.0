package db

import "github.com/FreakinWizard/wealthwise2.0/models"

func SyncDatabase(){
	DB.AutoMigrate(&models.User{})
}