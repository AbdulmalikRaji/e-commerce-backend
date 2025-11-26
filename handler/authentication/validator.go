package authentication

import (
	"github.com/abdulmalikraji/e-commerce/dto/authDto"
	"github.com/abdulmalikraji/e-commerce/utils"
	"github.com/abdulmalikraji/e-commerce/utils/genericResponse"
	"github.com/abdulmalikraji/e-commerce/utils/messages"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func LoginByEmailRequestValidator(ctx *fiber.Ctx) error {

	var request authDto.LoginByEmailRequest
	err := ctx.BodyParser(&request)
	if err != nil {
		log.Error(err)
		return genericResponse.ErrorResponse(ctx, fiber.StatusBadRequest, err.Error())
	}

	// Collect all validation errors and return them together as a map[field] -> []messages
	errs := make(map[string][]string)

	if request.Email == "" {
		errs["email"] = append(errs["email"], messages.CreateMsg(ctx, messages.RequiredField, map[string]string{"Field": "Email"}))
	} else if !utils.EmailRegex(request.Email) {
		errs["email"] = append(errs["email"], messages.CreateMsg(ctx, messages.InvalidFormat, map[string]string{"Field": "Email"}))
	}

	if request.Password == "" {
		errs["password"] = append(errs["password"], messages.CreateMsg(ctx, messages.RequiredField, map[string]string{"Field": "Password"}))
	}

	if len(errs) > 0 {
		return genericResponse.ErrorResponse(ctx, fiber.StatusBadRequest, "Validation failed", errs)
	}

	return ctx.Next()
}

func SignUpByEmailRequestValidator(ctx *fiber.Ctx) error {

	var request authDto.SignUpByEmailRequest
	err := ctx.BodyParser(&request)
	if err != nil {
		log.Error(err)
		return genericResponse.ErrorResponse(ctx, fiber.StatusBadRequest, err.Error())
	}

	errs := make(map[string][]string)

	// Email
	if request.Email == "" {
		errs["email"] = append(errs["email"], messages.CreateMsg(ctx, messages.RequiredField, map[string]string{"Field": "Email"}))
	} else if !utils.EmailRegex(request.Email) {
		errs["email"] = append(errs["email"], messages.CreateMsg(ctx, messages.InvalidFormat, map[string]string{"Field": "Email"}))
	}

	//Phone Number
	if request.PhoneNumber == "" {
		errs["phone_number"] = append(errs["phone_number"], messages.CreateMsg(ctx, messages.RequiredField, map[string]string{"Field": "Phone Number"}))
	} else if !utils.PhoneNumberRegex(request.PhoneNumber) {
		errs["phone_number"] = append(errs["phone_number"], messages.CreateMsg(ctx, messages.InvalidFormat, map[string]string{"Field": "Phone Number"}))
	}

	// Password
	if request.Password == "" {
		errs["password"] = append(errs["password"], messages.CreateMsg(ctx, messages.RequiredField, map[string]string{"Field": "Password"}))
	} else if len(request.Password) < 8 {
		errs["password"] = append(errs["password"], messages.CreateMsg(ctx, messages.MinLength, map[string]string{"Field": "Password", "Length": "8"}))
	}

	// First and last name
	if request.Firstname == "" {
		errs["first_name"] = append(errs["first_name"], messages.CreateMsg(ctx, messages.RequiredField, map[string]string{"Field": "First Name"}))
	}
	if request.Lastname == "" {
		errs["last_name"] = append(errs["last_name"], messages.CreateMsg(ctx, messages.RequiredField, map[string]string{"Field": "Last Name"}))
	}

	// Confirm password
	if request.ConfirmPassword == "" {
		errs["confirm_password"] = append(errs["confirm_password"], messages.CreateMsg(ctx, messages.RequiredField, map[string]string{"Field": "Confirm Password"}))
	} else if request.Password != request.ConfirmPassword {
		errs["confirm_password"] = append(errs["confirm_password"], messages.CreateMsg(ctx, messages.PasswordsDoNotMatch, nil))
	}

	// Role
	validRoles := map[string]bool{"buyer": true, "seller": true, "admin": true}
	if request.Role != "" && !validRoles[request.Role] {
		errs["role"] = append(errs["role"], messages.CreateMsg(ctx, messages.InvalidValue, map[string]string{"Field": "Role"}))
	}

	if len(errs) > 0 {
		return genericResponse.ErrorResponse(ctx, fiber.StatusBadRequest, "Validation failed", errs)
	}

	return ctx.Next()
}
