package service

import "github.com/AwesomeXjs/tma-bot/internal/model"

type IService interface {
	Registration(user *model.User) error
}
