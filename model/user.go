package model

type UserAPIResponce struct {
	Users []User `json:"users"`
}

type User struct {
	UserId     int    `json:"id"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	MiddleName string `json:"maidenName"`
	Age        int    `json:"age"`
}
