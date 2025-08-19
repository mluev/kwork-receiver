package repositories

import (
	"fmt"
	"kworker/models"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
)

func CreateUser(sender *tgbotapi.User) (models.User, error) {
	user := models.User{
		TelegramId: sender.ID,
		Username:   sender.UserName,
		Name:       fmt.Sprintf("%s %v", sender.FirstName, sender.LastName),
		ID:         uuid.New(),
	}

	err := DB.Create(&user).Error;

	log.Println(err)

	return user, err
}

func GetUser(telegramId int64) (models.User, error) {
	var result models.User

	err := DB.Model(&models.User{}).Where("telegram_id = ?", telegramId).First(&result).Error

	return result, err
}

func UpdateUser(user models.User) error {
	err := DB.Save(&user).Error

	return err
}
