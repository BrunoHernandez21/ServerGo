package models

type Category struct {
	Id            int
	Parent        int
	Date_added    int
	Last_modified int

	Code               string
	Name               string
	Slug               string
	Font_awesome_class string
	Thumbnail          string
}
