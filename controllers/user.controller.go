package controllers

type Data struct {
	Email string `json:"email" binding:"required"`
}
