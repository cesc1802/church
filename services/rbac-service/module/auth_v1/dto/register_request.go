package dto

type RegisterRequest struct {
	LoginID   string  `json:"login_id" binding:"required"`
	Password  string  `json:"password" binding:"required"`
	LastName  *string `json:"last_name" binding:"omitempty"`
	FirstName *string `json:"first_name" binding:"omitempty"`
}
