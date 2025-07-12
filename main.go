package main

import (
	"fmt"
	"go-tamboon/config"
	"go-tamboon/models"
	"go-tamboon/services"
	"go-tamboon/utils"
	"os"
)

func main() {
	fmt.Println("performing donations...")
	csvFilePath := os.Args[1]

	err := config.LoadConfig()

	if err != nil {
		fmt.Printf("Error loading configuration: %v\n", err)
		os.Exit(1)
	}

	donations, err := utils.ReadDonationsFromEncryptedFile(csvFilePath)
	if err != nil {
		fmt.Printf("Error reading encrypted file: %v\n", err)
		os.Exit(1)
	}

	omiseClient, err := services.GetOmiseClient()
	if err != nil {
		fmt.Printf("Error creating Omise client: %v\n", err)
		os.Exit(1)
	}

	donationService := services.NewDonationService(omiseClient)
	var successfulDonations []*models.Donation

	for _, donation := range donations {
		_, err := donationService.Donate(donation)
		if err == nil {
			successfulDonations = append(successfulDonations, donation)
		}
	}
	fmt.Println("done.")
	utils.SummaryDonations(donations, successfulDonations)
}
