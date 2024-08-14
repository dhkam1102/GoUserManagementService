// internal/models.go
package models

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	ConfirmPass string `json:"confirm_password"`
	Role        string `json:"role"` // make it optional
}

type UpdateRequest struct {
	OldEmail    string `json:"oldemail"`
	OldPassword string `json:"oldpassword"`
	NewEmail    string `json:"email"`
	NewPassword string `json:"password"`
}
