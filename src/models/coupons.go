package models

type Coupons struct {
	Id          int
	Created_at  int
	Expiry_date int

	Code                string
	Discount_percentage string
}
