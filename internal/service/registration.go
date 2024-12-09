package service

import (
	"net/http"

	"github.com/AwesomeXjs/tma-bot/internal/model"
)

func (s *Service) Registration(user *model.User) error {
	_, err := s.httpClient.NewRequest(http.MethodPost, s.resources.RegisterURL(), user)
	if err != nil {
		return err
	}
	return nil
}
