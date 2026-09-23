package handlers

import (
	"morng-dev/internal/core/domain/entities"
	"morng-dev/internal/core/domain/ports/services"
	"morng-dev/pkg/utils"
	"time"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authService services.AuthService
}

func NewAuthHandler(authService services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req entities.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ErrorResponse{
			Success: false,
			Message: "ข้อมูลไม่ถูกต้อง",
			Error:   err.Error(),
		})
	}

	if err := utils.ValidateStruct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ErrorResponse{
			Success: false,
			Message: "ข้อมูลไม่ครบถ้วน",
			Error:   err.Error(),
		})
	}

	user, err := h.authService.Register(c.Context(), &req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ErrorResponse{
			Success: false,
			Message: "ไม่สามารถลงทะเบียนได้",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(entities.ApiResponse{
		Success: true,
		Message: "ลงทะเบียนสำเร็จ",
		Data:    user,
	})
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req entities.LoginRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ErrorResponse{
			Success: false,
			Message: "ข้อมูลไม่ถูกต้อง",
			Error:   err.Error(),
		})
	}

	if err := utils.ValidateStruct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ErrorResponse{
			Success: false,
			Message: "ข้อมูลไม่ครบถ้วน",
			Error:   err.Error(),
		})
	}

	user, err := h.authService.Login(c.Context(), &req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ErrorResponse{
			Success: false,
			Message: "ไม่สามารถเข้าสู่ระบบได้",
			Error:   err.Error(),
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    user.Token,
		MaxAge:   int((12 * time.Hour).Seconds()),
		HTTPOnly: true,
		SameSite: "Strict",
	})

	return c.Status(fiber.StatusOK).JSON(entities.ApiResponse{
		Success: true,
		Message: "เข้าสู่ระบบสำเร็จ",
		Data:    user,
	})
}

func (h *AuthHandler) ForgotPassword(c *fiber.Ctx) error {
	var req entities.ForgotPasswordRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ErrorResponse{
			Success: false,
			Message: "Body Invalid Request",
			Error:   err.Error(),
		})
	}
	if err := h.authService.ForgotPassword(c.Context(), &req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ErrorResponse{
			Success: false,
			Error:   err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(entities.ApiResponse{
		Success: true,
		Message: "Password reset email sent successfully",
	})
}

func (h *AuthHandler) ResetPassword(c *fiber.Ctx) error {
	token := c.Query("token")
	var req entities.ResetPasswordRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ErrorResponse{
			Success: false,
			Message: "body invalid request",
		})
	}
	req.Token = token
	if err := h.authService.ResetPassword(c.Context(), &req); err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(entities.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(entities.ApiResponse{
		Success: true,
		Message: "password reset successfully",
	})

}
