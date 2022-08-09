package model

type User struct {
	UID      int    `json:"uid"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	ImageURL string `json:"imageUrl"`
	Website  string `json:"website"`
}
