package filters

import "github.com/go-telegram/bot/models"

func IsPhoto(update *models.Update) bool {
	return update.Message != nil && len(update.Message.Photo) > 0
}

func IsVideo(update *models.Update) bool {
	return update.Message != nil && update.Message.Video != nil
}
