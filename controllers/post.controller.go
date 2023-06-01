package controllers

type PostData struct {
	Email string `json:"email" binding:"required"`
}
