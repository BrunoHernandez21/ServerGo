package models

import "encoding/json"

type Users_IO struct {
	Id                int
	First_name        string
	Last_name         string
	Email             string
	Password          string
	Social_links      map[string]string
	Biography         string
	Role_id           int
	Date_added        int
	Last_modified     int
	Watch_history     string
	Wishlist          string
	Title             string
	Paypal_keys       []map[string]string
	Stripe_keys       string
	Verification_code string
	Image             string
	Status            int
	Is_instructor     int
	Skills            string
	Payment_keys      string
	Number            string
	Name              string
	BirthDay          string
}

func (a Users_IO) To_user() Users {
	newUser := Users{
		Id:                a.Id,
		First_name:        a.First_name,
		Last_name:         a.Last_name,
		Email:             a.Email,
		Password:          a.Password,
		Biography:         a.Biography,
		Role_id:           a.Role_id,
		Date_added:        a.Date_added,
		Last_modified:     a.Last_modified,
		Watch_history:     a.Watch_history,
		Wishlist:          a.Wishlist,
		Title:             a.Title,
		Stripe_keys:       a.Stripe_keys,
		Verification_code: a.Verification_code,
		Image:             a.Image,
		Status:            a.Status,
		Is_instructor:     a.Is_instructor,
		Skills:            a.Skills,
		Payment_keys:      a.Payment_keys,
	}
	jsonStr, err := json.Marshal(a.Social_links)
	if err != nil {
		return newUser
	}
	jsonStr2, err2 := json.Marshal(a.Social_links)
	if err2 != nil {
		return newUser
	}
	newUser.Social_links = string(jsonStr)
	newUser.Paypal_keys = string(jsonStr2)
	return newUser
}

type Users struct {
	Id                int
	First_name        string
	Last_name         string
	Email             string
	Password          string
	Social_links      string
	Biography         string
	Role_id           int
	Date_added        int
	Last_modified     int
	Watch_history     string
	Wishlist          string
	Title             string
	Paypal_keys       string
	Stripe_keys       string
	Verification_code string
	Image             string
	Status            int
	Is_instructor     int
	Skills            string
	Payment_keys      string
	Number            string
	Name              string
	BirthDay          string
}

func (a Users) To_user_io() Users_IO {
	newUser := Users_IO{
		Id:                a.Id,
		First_name:        a.First_name,
		Last_name:         a.Last_name,
		Email:             a.Email,
		Password:          a.Password,
		Biography:         a.Biography,
		Role_id:           a.Role_id,
		Date_added:        a.Date_added,
		Last_modified:     a.Last_modified,
		Watch_history:     a.Watch_history,
		Wishlist:          a.Wishlist,
		Title:             a.Title,
		Stripe_keys:       a.Stripe_keys,
		Verification_code: a.Verification_code,
		Image:             a.Image,
		Status:            a.Status,
		Is_instructor:     a.Is_instructor,
		Skills:            a.Skills,
		Payment_keys:      a.Payment_keys,
	}

	json.Unmarshal([]byte(a.Social_links), &newUser.Social_links)
	json.Unmarshal([]byte(a.Paypal_keys), &newUser.Paypal_keys)

	return newUser

}
