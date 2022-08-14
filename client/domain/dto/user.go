package dto

//UserValidator -
type UserValidator struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
}

// GetUserId -
type GetUserId struct {
	ID string `json:"id" uri:"id" binding:"required" `
}
