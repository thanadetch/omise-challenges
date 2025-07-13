package utils

import "go-tamboon/models"

func ClearSensitiveData(donations []*models.Donation) {
	for _, donation := range donations {
		if donation != nil {
			donation.Card = ""
			donation.CVV = ""
			donation.ExpMonth = 0
			donation.ExpYear = 0
		}
	}
}
