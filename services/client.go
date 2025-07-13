package services

import (
	"github.com/omise/omise-go"
	"go-tamboon/config"
)

func GetOmiseClient() (*omise.Client, error) {
	client, err := omise.NewClient(
		config.AppConfig.OmisePublicKey,
		config.AppConfig.OmiseSecretKey,
	)

	if err != nil {
		return nil, err
	}
	return client, nil
}
