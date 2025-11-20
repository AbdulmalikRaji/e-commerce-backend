package authDto

import "time"

type LoginByEmailRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type LoginByEmailResponse struct {
	AccessToken string    `json:"access_token"`
	ExpiresAt   time.Time `json:"expires_at"`
}

type RefreshTokenResponse struct {
	AccessToken string    `json:"access_token"`
	ExpiresAt   time.Time `json:"expires_at"`
}

type LogoutRequest struct {
	AccessToken string `json:"access_token" validate:"required"`
}

type SignUpByEmailRequest struct {
	Firstname       string      `json:"first_name" validate:"required"`
	Lastname        string      `json:"last_name" validate:"required"`
	Email           string      `json:"email" validate:"required,email"`
	PhoneNumber     string      `json:"phone_number" validate:"required,e164"`
	Password        string      `json:"password" validate:"required,min=8"`
	ConfirmPassword string      `json:"confirm_password" validate:"required,eqfield=Password"`
	Address         UserAddress `json:"address" validate:"required,dive"`
	Role            string      `json:"role" validate:"required,oneof=buyer seller admin"`
}

type UserAddress struct {
	Line1      string `json:"line1" validate:"required"`
	Line2      string `json:"line2"`
	City       string `json:"city" validate:"required"`
	State      string `json:"state" validate:"required"`
	PostalCode string `json:"postal_code"`
	Country    string `json:"country" validate:"required"`
	IsDefault  bool   `json:"is_default"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type ResetPasswordRequest struct {
	ResetToken      string `json:"reset_token" validate:"required"`
	NewPassword     string `json:"new_password" validate:"required,min=8"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=NewPassword"`
}

type ForgotPasswordResponse struct {
	Message string `json:"message"`
}
