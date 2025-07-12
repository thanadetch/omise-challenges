package models

type Donation struct {
	Name     string `csv:"Name"`
	Amount   int64  `csv:"AmountSubunits"`
	Card     string `csv:"CCNumber"`
	CVV      string `csv:"CVV"`
	ExpMonth int    `csv:"ExpMonth"`
	ExpYear  int    `csv:"ExpYear"`
}
