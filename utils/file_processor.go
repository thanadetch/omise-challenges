package utils

import (
	"encoding/csv"
	"fmt"
	"github.com/gocarina/gocsv"
	"go-tamboon/cipher"
	"go-tamboon/models"
	"os"
)

func ValidateFileExists(path string) error {
	if path == "" {
		return fmt.Errorf("file path is empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("file does not exist: %s", path)
	}

	return nil
}

func ReadDonationsFromEncryptedFile(filePath string) ([]*models.Donation, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	rot128Reader, err := cipher.NewRot128Reader(file)
	if err != nil {
		return nil, fmt.Errorf("error creating Rot128Reader: %v", err)
	}

	csvReader := csv.NewReader(rot128Reader)

	var donations []*models.Donation
	err = gocsv.UnmarshalCSV(csvReader, &donations)
	if err != nil {
		return nil, fmt.Errorf("error parsing CSV with tags: %v", err)
	}
	return donations, nil
}
