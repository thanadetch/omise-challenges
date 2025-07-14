package utils

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"go-tamboon/models"
	"sort"
)

func convertSatangToTHB(satang float64) float64 {
	return satang / 100
}

func SummaryDonations(donations []*models.Donation, successfulDonations []*models.Donation) {
	totalReceived := float64(0)
	successfullyDonated := float64(0)
	var topDonors []string

	for _, donation := range donations {
		totalReceived += float64(donation.Amount)
	}
	totalReceived = convertSatangToTHB(totalReceived)

	sort.Slice(successfulDonations, func(i, j int) bool {
		return successfulDonations[i].Amount > successfulDonations[j].Amount
	})

	for i, donation := range successfulDonations {
		successfullyDonated += float64(donation.Amount)
		if i < 3 {
			topDonors = append(topDonors, donation.Name)
		}
	}
	successfullyDonated = convertSatangToTHB(successfullyDonated)

	donationSummary := &models.DonationSummary{
		TotalReceived:       totalReceived,
		SuccessfullyDonated: successfullyDonated,
		FaultyDonation:      totalReceived - successfullyDonated,
		AveragePerPerson:    successfullyDonated / float64(len(donations)),
		TopDonors:           topDonors,
	}

	PrintSummaryDonations(donationSummary)
}

func PrintSummaryDonations(donationSummary *models.DonationSummary) {
	format := func(amount float64) string {
		return humanize.FormatFloat("#,###.##", amount)
	}

	fmt.Println()
	fmt.Printf("%22s THB %16s\n", "total received:", format(donationSummary.TotalReceived))
	fmt.Printf("%22s THB %16s\n", "successfully donated:", format(donationSummary.SuccessfullyDonated))
	fmt.Printf("%22s THB %16s\n", "faulty donation:", format(donationSummary.FaultyDonation))
	fmt.Println()
	fmt.Printf("%22s THB %16s\n", "average per person:", format(donationSummary.AveragePerPerson))
	fmt.Printf("%22s", "top donors:")
	for _, donor := range donationSummary.TopDonors {
		fmt.Printf(" %s \n", donor)
		fmt.Printf("%22s", "")
	}
	fmt.Println()
}
