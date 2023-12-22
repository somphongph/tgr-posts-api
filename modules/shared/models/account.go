package models

type AccountProfile struct {
	Id          string `json:"id"`
	Username    string `json:"username"`
	DisplayName string `json:"displayName"`
	PhotoUrl    string `json:"photoUrl"`
}
