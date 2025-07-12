package models

type DonationSummary struct {
	TotalReceived       int64
	SuccessfullyDonated int64
	FaultyDonation      int64
	AveragePerPerson    float64
	TopDonors           []string
}
