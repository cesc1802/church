package dto

type RegisterRequest struct {
	UserID   string `json:"userID" binding:"required"`
	Password string `json:"password" binding:"required"`
}
