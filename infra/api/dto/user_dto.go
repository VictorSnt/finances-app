package dto

type CreateUserDTO struct {
	Username string  `json:"username" binding:"required"`
	Income   float64 `json:"income" binding:"required,gt=0"`
}

type UpdateUser struct {
	Username string  `json:"username" binding:"omitempty"`
	Income   float64 `json:"income" binding:"omitempty,gt=0"`
}

type UserDTO struct {
	ID       int     `json:"user_id"`
	Username string  `json:"username"`
	Income   float64 `json:"income"`
}
