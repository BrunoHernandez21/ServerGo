package models

type Addons struct {
	Id         int
	Status     int
	Created_at int
	Updated_at int

	Name              string
	Unique_identifier string
	Version           string
	About             string
	Purchase_code     string
}
