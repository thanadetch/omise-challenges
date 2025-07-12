package services

import (
	"fmt"
	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
	"go-tamboon/models"
	"time"
)

type DonationService struct {
	client *omise.Client
}

func NewDonationService(client *omise.Client) *DonationService {
	return &DonationService{
		client: client,
	}
}

func (s *DonationService) getCard(donation *models.Donation) (*omise.Card, error) {
	card := &omise.Card{}
	err := s.client.Do(card, &operations.CreateToken{
		Name:            donation.Name,
		Number:          donation.Card,
		ExpirationMonth: time.Month(donation.ExpMonth),
		ExpirationYear:  donation.ExpYear,
		SecurityCode:    donation.CVV,
	})
	if err != nil {
		return nil, err
	}
	return card, nil
}

func (s *DonationService) Donate(donation *models.Donation) (*omise.Charge, error) {
	card, err := s.getCard(donation)
	if err != nil {
		return nil, err
	}

	result := &omise.Charge{}
	err = s.client.Do(result, &operations.CreateCharge{
		Amount:      donation.Amount,
		Currency:    "thb",
		Description: fmt.Sprintf("Donation from %s", donation.Name),
		Card:        card.ID,
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}
