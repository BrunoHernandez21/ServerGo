package models

type Comment struct {
	Id             int
	User_id        int
	Commentable_id int
	Date_added     int
	Last_modified  int

	Body             string
	Commentable_type string
}
