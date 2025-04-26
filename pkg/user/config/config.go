package config

type User struct {
	Name   string `json:"name" binding:"required"`
	PAN    string `json:"pan" binding:"required,pan"`
	Mobile string `json:"mobile" binding:"required,mobile"`
	Email  string `json:"email" binding:"required,email"`
}
