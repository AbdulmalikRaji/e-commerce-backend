package authentication

import (
	"github.com/abdulmalikraji/e-commerce/dto/authDto"
	"github.com/abdulmalikraji/e-commerce/services"
	"github.com/abdulmalikraji/e-commerce/utils/genericResponse"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler interface {
	SignUp(ctx *fiber.Ctx) error
	LoginByEmail(ctx *fiber.Ctx) error
	ValidateToken(ctx *fiber.Ctx) error
	RefreshToken(ctx *fiber.Ctx) error
	Logout(ctx *fiber.Ctx) error
	ForgotPassword(ctx *fiber.Ctx) error
}

type authHandler struct {
	service services.AuthService
}

func New(service services.AuthService) AuthHandler {
	return authHandler{
		service: service,
	}
}

func (c authHandler) LoginByEmail(ctx *fiber.Ctx) error {
	var loginRequest authDto.LoginByEmailRequest
	if err := ctx.BodyParser(&loginRequest); err != nil {
		return genericResponse.ErrorResponse(ctx, fiber.StatusBadRequest, err.Error())
	}

	data, status, err := c.service.LoginByEmail(ctx, loginRequest)
	if err != nil {
		return genericResponse.ErrorResponse(ctx, fiber.StatusUnauthorized, err.Error())
	}

	return genericResponse.SuccessResponse(ctx, status, data, "Login successful")
}

func (c authHandler) SignUp(ctx *fiber.Ctx) error {
	var signUpRequest authDto.SignUpByEmailRequest
	if err := ctx.BodyParser(&signUpRequest); err != nil {
		return genericResponse.ErrorResponse(ctx, fiber.StatusBadRequest, err.Error())
	}

	status, err := c.service.SignupByEmail(ctx, signUpRequest)
	if err != nil {
		return genericResponse.ErrorResponse(ctx, fiber.StatusBadRequest, err.Error())
	}

	return genericResponse.SuccessResponse(ctx, status, nil, "Sign up successful, Check your email for verification link")
}

// ValidateToken validates the current token
func (c authHandler) ValidateToken(ctx *fiber.Ctx) error {
	// Get token from Authorization header
	authHeader := ctx.Get("Authorization")
	if authHeader == "" || len(authHeader) <= 7 || authHeader[:7] != "Bearer " {
		return genericResponse.ErrorResponse(ctx, fiber.StatusUnauthorized, "Invalid Authorization header format")
	}

	token := authHeader[7:]
	if status, err := c.service.ValidateToken(ctx, token); err != nil {
		return genericResponse.ErrorResponse(ctx, status, "Invalid token")
	}

	return genericResponse.SuccessResponse(ctx, fiber.StatusOK, nil, "Token is valid")
}

// RefreshToken refreshes the current access token using the refresh token
func (c authHandler) RefreshToken(ctx *fiber.Ctx) error {
	data, status, err := c.service.RefreshToken(ctx)
	if err != nil {
		return genericResponse.ErrorResponse(ctx, status, "Failed to refresh token")
	}

	return genericResponse.SuccessResponse(ctx, status, data, "Token refreshed successfully")
}

func (c authHandler) Logout(ctx *fiber.Ctx) error {
	var logoutRequest authDto.LogoutRequest

	authHeader := ctx.Get("Authorization")
	if authHeader == "" || len(authHeader) <= 7 || authHeader[:7] != "Bearer " {
		return genericResponse.ErrorResponse(ctx, fiber.StatusUnauthorized, "Invalid Authorization header format")
	}

	token := authHeader[7:]
	if status, err := c.service.ValidateToken(ctx, token); err != nil {
		return genericResponse.ErrorResponse(ctx, status, "Invalid token")
	}

	logoutRequest.AccessToken = token

	status, err := c.service.Logout(ctx, logoutRequest)
	if err != nil {
		return genericResponse.ErrorResponse(ctx, status, err.Error())
	}

	return genericResponse.SuccessResponse(ctx, status, nil, "Logout successful")
}

func (c authHandler) ForgotPassword(ctx *fiber.Ctx) error {
	var forgotPasswordRequest authDto.ForgotPasswordRequest
	if err := ctx.BodyParser(&forgotPasswordRequest); err != nil {
		return genericResponse.ErrorResponse(ctx, fiber.StatusBadRequest, err.Error())
	}

	status, err := c.service.ForgotPassword(ctx, forgotPasswordRequest)
	if err != nil {
		return genericResponse.ErrorResponse(ctx, status, err.Error())
	}

	return genericResponse.SuccessResponse(ctx, status, nil, "Password recovery email sent")
}

func (c authHandler) ResetPassword(ctx *fiber.Ctx) error {
	var req authDto.ResetPasswordRequest
	if err := ctx.BodyParser(&req); err != nil {
		return genericResponse.ErrorResponse(ctx, fiber.StatusBadRequest, err.Error())
	}

	if req.Type != "recovery" {
		return genericResponse.ErrorResponse(ctx, fiber.StatusBadRequest, "Invalid reset type")
	}

	status, err := c.service.ResetPassword(ctx, req)
	if err != nil {
		return genericResponse.ErrorResponse(ctx, status, err.Error())
	}

	return genericResponse.SuccessResponse(ctx, status, nil, "Password reset successful")
}

func (c authHandler) ResetPasswordPage(ctx *fiber.Ctx) error {
	html := `
<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8" />
  <title>Reset Password</title>
</head>
<body>
  <h2>Reset Password</h2>

  <form id="resetForm">
    <input
      type="password"
      id="password"
      placeholder="New password"
      required
    />
    <br /><br />
    <button type="submit">Reset Password</button>
  </form>

  <p id="status"></p>

  <script>
    const status = document.getElementById("status");

    // Supabase puts tokens in the URL hash
    const hash = window.location.hash.substring(1);
    const params = new URLSearchParams(hash);

    const accessToken = params.get("access_token");
    const refreshToken = params.get("refresh_token");
    const type = params.get("type");

    if (!accessToken || type !== "recovery") {
      status.innerText = "Invalid or missing recovery token.";
      throw new Error("Invalid recovery token");
    }

    document.getElementById("resetForm").addEventListener("submit", async (e) => {
      e.preventDefault();

      const password = document.getElementById("password").value;

      const res = await fetch("/auth/reset-password", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          access_token: accessToken,
          refresh_token: refreshToken,
          password: password,
		  type: type
        })
      });

      const data = await res.json();
      status.innerText = data.message || data.error;
    });
  </script>
</body>
</html>
`
	return ctx.Type("html").SendString(html)
}
