package repositories

import (
	"gorm.io/gorm"
	"kworker/database"
)

var DB *gorm.DB = database.Init()
