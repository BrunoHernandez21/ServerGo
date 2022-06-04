package models

type Payment struct {
	Id                        int
	User_id                   int
	Payment_type              string
	Course_id                 int
	Amount                    int
	Date_added                int
	Last_modified             int
	Admin_revenue             string
	Instructor_revenue        string
	Instructor_payment_status int
	Transaction_id            string
	Session_id                string
	Coupon                    string
}
