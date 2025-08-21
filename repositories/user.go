package repositories

import (
	"fmt"
	"kworker/models"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
)

func CreateUser(sender *tgbotapi.User, chatId int64) (models.User, error) {
	log.Println(sender.ID)
	user := models.User{
		TelegramId: sender.ID,
		ChatId:     chatId,
		Username:   sender.UserName,
		Name:       fmt.Sprintf("%s %v", sender.FirstName, sender.LastName),
		ID:         uuid.New(),
	}

	err := DB.Create(&user).Error;

	log.Println(err)

	return user, err
}

func GetUser(telegramId int64) (models.User, error) {
	log.Println(telegramId)
	var result models.User

	err := DB.Model(&models.User{}).Where("telegram_id = ?", telegramId).First(&result).Error

	return result, err
}

func UpdateUser(user models.User) error {
	err := DB.Save(&user).Error

	return err
}

func GetReadyUsers() ([]models.User, error) {
	var users []models.User

	err := DB.Model(&models.User{}).Where("receiving = ?", true).Find(&users).Error

	return users, err
}
