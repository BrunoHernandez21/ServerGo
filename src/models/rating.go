package models

type Rating struct {
	Id            int
	Rating        int
	User_id       int
	Ratable_id    int
	Ratable_type  string
	Date_added    int
	Last_modified int
	Review        string
}
