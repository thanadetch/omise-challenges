package utils

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"go-tamboon/models"
	"sort"
)

func SummaryDonations(donations []*models.Donation, successfulDonations []*models.Donation) {
	totalReceived := int64(0)
	successfullyDonated := int64(0)
	var topDonors []string

	for _, donation := range donations {
		totalReceived += donation.Amount
	}

	sort.Slice(successfulDonations, func(i, j int) bool {
		return successfulDonations[i].Amount > successfulDonations[j].Amount
	})

	for i, donation := range successfulDonations {
		successfullyDonated += donation.Amount
		if i < 3 {
			topDonors = append(topDonors, donation.Name)
		}
	}

	donationSummary := &models.DonationSummary{
		TotalReceived:       totalReceived,
		SuccessfullyDonated: successfullyDonated,
		FaultyDonation:      totalReceived - successfullyDonated,
		AveragePerPerson:    float64(successfullyDonated) / float64(len(donations)),
		TopDonors:           topDonors,
	}

	PrintSummaryDonations(donationSummary)
}

func PrintSummaryDonations(donationSummary *models.DonationSummary) {
	format := func(amount float64) string {
		return humanize.FormatFloat("#,###.##", amount)
	}

	fmt.Println()
	fmt.Printf("%22s THB %16s\n", "total received:", format(float64(donationSummary.TotalReceived)))
	fmt.Printf("%22s THB %16s\n", "successfully donated:", format(float64(donationSummary.SuccessfullyDonated)))
	fmt.Printf("%22s THB %16s\n", "faulty donation:", format(float64(donationSummary.FaultyDonation)))
	fmt.Println()
	fmt.Printf("%22s THB %16s\n", "average per person:", format(donationSummary.AveragePerPerson))
	fmt.Printf("%22s", "top donors:")
	for _, donor := range donationSummary.TopDonors {
		fmt.Printf(" %s \n", donor)
		fmt.Printf("%22s", "")
	}
	fmt.Println()
}
