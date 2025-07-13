package services

import (
	"context"
	"fmt"
	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
	"go-tamboon/models"
	"golang.org/x/time/rate"
	"sync"
	"time"
)

type DonationService struct {
	client  *omise.Client
	limiter *rate.Limiter
}

func NewDonationService(client *omise.Client) *DonationService {
	return &DonationService{
		client:  client,
		limiter: rate.NewLimiter(rate.Limit(5), 1),
	}
}

func (s *DonationService) getCard(donation *models.Donation) (*omise.Card, error) {
	s.limiter.Wait(context.Background())
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

	s.limiter.Wait(context.Background())
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

func (s *DonationService) ProcessDonations(donations []*models.Donation) []*models.Donation {
	numWorkers := 3
	donationChan := make(chan *models.Donation)
	resultChan := make(chan *models.Donation)

	var wg sync.WaitGroup

	// Start workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for donation := range donationChan {
				_, err := s.Donate(donation)
				if err == nil {
					resultChan <- donation
				}
			}
		}()
	}

	// Send donations to workers
	go func() {
		for _, donation := range donations {
			donationChan <- donation
		}
		close(donationChan)
	}()

	// Close result channel when all workers are done
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Collect successful donations - wait for channel to close
	var successfulDonations []*models.Donation
	for donation := range resultChan {
		successfulDonations = append(successfulDonations, donation)
	}

	return successfulDonations
}
