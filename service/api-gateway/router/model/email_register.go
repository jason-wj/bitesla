package model

type EmailRegister struct {
	Password string `json:"password" example:"sja123"`
	Email    string `json:"email" example:"idea_wj@163.com"`
}
