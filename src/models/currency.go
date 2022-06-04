package models

type Currency struct {
	Id int

	Name             string
	Code             string
	Symbol           string
	Stripe_supported string
}
