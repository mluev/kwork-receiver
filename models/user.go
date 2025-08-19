package models

import "github.com/google/uuid"

type User struct {
	ID         uuid.UUID
	Name       string
	Keywords   string
	Username   string
	Blocked    bool
	Receiving  bool
	TelegramId int64
}
