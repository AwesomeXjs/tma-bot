package filters

import "github.com/go-telegram/bot/models"

const (
	start = "/start"
	help  = "/help"
	id    = "/id"
)

func IsStart(update *models.Update) bool {
	return update.Message != nil && update.Message.Text == start
}

func IsHelp(update *models.Update) bool {
	return update.Message != nil && update.Message.Text == help
}

func IsMyID(update *models.Update) bool {
	return update.Message != nil && update.Message.Text == id
}
