package models

type DonationSummary struct {
	TotalReceived       float64
	SuccessfullyDonated float64
	FaultyDonation      float64
	AveragePerPerson    float64
	TopDonors           []string
}
