package services

import (
	"github.com/omise/omise-go"
	"go-tamboon/config"
	"sync"
)

var (
	client *omise.Client
	err    error
	once   sync.Once
)

func GetOmiseClient() (*omise.Client, error) {
	once.Do(func() {
		client, err = omise.NewClient(
			config.AppConfig.OmisePublicKey,
			config.AppConfig.OmiseSecretKey,
		)
	})

	if err != nil {
		return nil, err
	}
	return client, nil
}
