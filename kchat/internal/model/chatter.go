package model

type Chatter struct {
	UID      int    `json:"uid"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	ImageURL string `json:"img_url"`
}
