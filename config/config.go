package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	OmiseSecretKey string
	OmisePublicKey string
}

var AppConfig *Config

func LoadConfig() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	AppConfig = &Config{
		OmiseSecretKey: os.Getenv("OMISE_SECRET_KEY"),
		OmisePublicKey: os.Getenv("OMISE_PUBLIC_KEY"),
	}

	if err := AppConfig.validate(); err != nil {
		return err
	}

	return nil
}

func (c *Config) validate() error {
	if c.OmiseSecretKey == "" {
		return fmt.Errorf("OMISE_SECRET_KEY environment variable is required")
	}
	if c.OmisePublicKey == "" {
		return fmt.Errorf("OMISE_PUBLIC_KEY environment variable is required")
	}
	return nil
}
