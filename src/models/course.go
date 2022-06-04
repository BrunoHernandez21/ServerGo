package models

type Course struct {
	Id                       int
	Title                    string
	Short_description        string
	Description              string
	Outcomes                 string
	Language                 string
	Category_id              int
	Sub_category_id          int
	Section                  string
	Requirements             string
	Price                    int
	Discount_flag            int
	Discounted_price         int
	Level                    string
	User_id                  string
	Thumbnail                string
	Video_url                string
	Date_added               int
	Last_modified            int
	Course_type              string
	Is_top_course            int
	Is_admin                 int
	Status                   string
	Course_overview_provider string
	Meta_keywords            string
	Meta_description         string
	Is_free_course           int
	Multi_instructor         int
	Creator                  int
}
