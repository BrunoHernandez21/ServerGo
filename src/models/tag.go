package models

type Tag struct {
	Id            int
	Tag           string
	Tagable_id    int
	Tagable_type  string
	Date_added    int
	Last_modified int
}
