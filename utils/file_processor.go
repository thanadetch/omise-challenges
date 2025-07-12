package utils

import (
	"fmt"
	"github.com/gocarina/gocsv"
	"go-tamboon/cipher"
	"go-tamboon/models"
	"io"
	"os"
	"strings"
)

func ReadDonationsFromEncryptedFile(filePath string) ([]*models.Donation, error) {
	decryptedContent, err := readAndDecryptFile(filePath)
	if err != nil {
		return nil, err
	}

	donations, err := parseCSVToDonations(decryptedContent)
	if err != nil {
		return nil, err
	}

	return donations, nil
}

func readAndDecryptFile(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	rot128Reader, err := cipher.NewRot128Reader(file)
	if err != nil {
		return nil, fmt.Errorf("error creating Rot128Reader: %v", err)
	}

	return io.ReadAll(rot128Reader)
}

func parseCSVToDonations(content []byte) ([]*models.Donation, error) {
	reader := strings.NewReader(string(content))
	var donations []*models.Donation
	err := gocsv.Unmarshal(reader, &donations)
	if err != nil {
		return nil, fmt.Errorf("error parsing CSV with tags: %v", err)
	}

	return donations, nil
}
