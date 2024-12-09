package service

import (
	"github.com/AwesomeXjs/tma-bot/internal/client/http"
	"github.com/AwesomeXjs/tma-bot/internal/config"
)

type Service struct {
	httpClient http.IHttpClient
	resources  config.IResources
}

func New(httpClient http.IHttpClient, resources config.IResources) IService {
	return &Service{
		httpClient: httpClient,
		resources:  resources,
	}
}
