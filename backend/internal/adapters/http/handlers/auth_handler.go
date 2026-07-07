package handlers

import (
	"morng-dev/internal/core/domain/entities"
	"morng-dev/internal/core/domain/ports/services"
	"morng-dev/pkg/utils"

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
		return c.Status(fiber.StatusInternalServerError).JSON(entities.ErrorResponse{
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
	return c.Status(fiber.StatusOK).JSON(entities.ApiResponse{
		Success: true,
		Message: "เข้าสู่ระบบาสำเร็จ",
		Data:    user,
	})
}
