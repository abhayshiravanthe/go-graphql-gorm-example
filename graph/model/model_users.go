package model

type User struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	FirstName *string `json:"firstName"`
	LastName  *string `json:"lastName"`
	Mobile    *string `json:"mobile"`
	IsLocked  *bool   `json:"isLocked"`
	Password  string  `json:"password"`
}
