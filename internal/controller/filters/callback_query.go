package filters

import "github.com/go-telegram/bot/models"

func IsQuery(update *models.Update) bool {
	return update.CallbackQuery != nil && update.CallbackQuery.Data == "info"
}
