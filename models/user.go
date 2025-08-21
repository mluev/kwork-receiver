package models

import "github.com/google/uuid"

type User struct {
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	Name       string
	Keywords   string
	Username   string
	Blocked    bool
	Receiving  bool
	TelegramId int64     `gorm:"column:telegram_id"`
	ChatId     int64     `gorm:"column:chat_id"`
}
