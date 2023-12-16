package model

type Auth struct {
	Username string `json:"username" validate:"required,alphanum,max=100"`
	Password string `json:"password" validate:"required"`
}