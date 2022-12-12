package dto

type LoginRequest struct {
	LoginID  string `json:"login_id" binding:"required"`
	Password string `json:"password" binding:"required"`
}
