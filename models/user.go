package models

type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"` // "-" means this field won't be shown in JSON responses
}
