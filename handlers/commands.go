package handlers

import (
	"errors"
	"fmt"
	"kworker/repositories"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
)

func Commands(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	switch update.Message.Command() {
		case "start":
			start(bot, update)
		case "on":
			toggleStatus(bot, update, true)
		case "off":
			toggleStatus(bot, update, false)
	}
}


func toggleStatus(bot *tgbotapi.BotAPI, update tgbotapi.Update, status bool) {
	sender := update.Message.From

	user, err := repositories.GetUser(sender.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			user, err = repositories.CreateUser(sender, update.Message.Chat.ID)
			if err != nil {
				log.Fatalf("Error creating user: %v", err)
			}
		} else {
			log.Fatalf("Database error: %v", err)
		}
	}

	user.Receiving = !user.Receiving
	err = repositories.UpdateUser(user)
	if err != nil {
		log.Fatalf("Error updating user: %v", err)
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

	if user.Receiving {
		msg.Text = "You are now receiving orders"
	} else {
		msg.Text = "You are no longer receiving orders"
	}

	bot.Send(msg)
}

func start(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	sender := update.Message.From

	_, err := repositories.GetUser(sender.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("new user")
			_, err = repositories.CreateUser(sender, update.Message.Chat.ID)
			if err != nil {
				log.Fatalf("Error creating user: %v", err)
			}
		} else {
			log.Fatalf("Database error: %v", err)
		}
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	msg.Text = "Start getting orders - /on\nStop getting orders - /off"

	bot.Send(msg)
}

func me(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	user, err := repositories.GetUser(update.Message.From.ID)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	if (err != nil) {
		msg.Text = "Error fetching user information"
	} else {
		msg.Text = fmt.Sprintf("User ID: %d\nUsername: %s\nFirst Name: %s", user.ID, user.Username, user.Name)
	}

	bot.Send(msg)
}
