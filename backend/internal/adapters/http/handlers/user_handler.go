package handlers

import (
	"morng-dev/internal/core/domain/entities"
	"morng-dev/internal/core/domain/ports/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	users, pagination, err := h.userService.GetUsers(c.Context(), page, limit)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ErrorResponse{
			Success: false,
			Message: "ไม่สามารถดึงข้อมูลผู้ใช้ได้",
		})
	}
	return c.Status(fiber.StatusOK).JSON(entities.ApiResponse{
		Success:    true,
		Message:    "ดึงข้อมูลสำเร็จ",
		Data:       users,
		Pagination: pagination,
	})
}
