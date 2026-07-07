package handlers

import (
	"fmt"
	"morng-dev/internal/core/domain/entities"
	"morng-dev/internal/core/domain/ports/services"
	"morng-dev/pkg/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type TodoHandler struct {
	todoService services.TodoService
}

func NewTodoHandler(todoService services.TodoService) *TodoHandler {
	return &TodoHandler{todoService: todoService}
}

func (h *TodoHandler) CreateTodo(c *fiber.Ctx) error {
	var req entities.TodoRequest
	userID := c.Locals("userID").(uuid.UUID)
	fmt.Println(userID)
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
			Message: "รูปแบบข้อมูลไม่ถูกต้อง",
		})
	}

	todo, err := h.todoService.CreateTodo(c.Context(), userID, &req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ErrorResponse{
			Success: false,
			Message: "ไม่สามารถสร้างได้",
		})
	}
	return c.Status(fiber.StatusOK).JSON(entities.ApiResponse{
		Success: true,
		Message: "สร้างสำเร็จ",
		Data:    todo,
	})
}

func (h *TodoHandler) GetTodoByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ApiResponse{
			Success: false,
			Message: "รูปแบบ ID ไม่ถูกต้อง",
		})
	}
	todo, err := h.todoService.GetTodoByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ErrorResponse{
			Success: false,
			Message: "ไม่พบรายการ",
		})
	}

	return c.Status(fiber.StatusOK).JSON(entities.ApiResponse{
		Success: true,
		Message: "ดึงรายการสำเร็จ",
		Data:    todo,
	})
}

func (h *TodoHandler) GetAllTodo(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	todos, pagination, err := h.todoService.GetAllTodo(c.Context(), page, limit)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ErrorResponse{
			Success: false,
			Message: "ไม่สามารถดึงผู้ใช้ได้",
		})
	}
	return c.Status(fiber.StatusOK).JSON(entities.ApiResponse{
		Success:    true,
		Message:    "ดึงtodoสำเร็จ",
		Data:       todos,
		Pagination: pagination,
	})
}

func (h *TodoHandler) UpdateTodoStatus(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	userID := c.Locals("userID").(uuid.UUID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ErrorResponse{
			Success: false,
			Message: "id invalid",
		})
	}

	var req entities.TodoStatusRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ErrorResponse{
			Success: false,
			Message: "ข้อมูลไม่ถูกต้อง",
		})
	}

	if err := utils.ValidateStruct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ErrorResponse{
			Success: false,
			Message: "รูปแบบข้อมูลไม่ถูกต้อง",
		})
	}

	if err := h.todoService.UpdateStatusTodo(c.Context(), userID, id, &req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entities.ErrorResponse{
			Success: false,
			Message: "ไม่สามารถอัพเดตได้",
			Error:   err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(entities.ApiResponse{
		Success: true,
		Message: "อัพเดตสำเร็จ",
	})
}

func (h *TodoHandler) DeleteTodo(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	userID := c.Locals("userID").(uuid.UUID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ErrorResponse{
			Success: false,
			Message: "id ไม่ถูกต้อง",
		})
	}
	if err := h.todoService.DeleteTodo(c.Context(), userID, id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entities.ErrorResponse{
			Success: false,
			Message: "ไม่สามารถลบได้",
			Error:   err.Error(),
		})
	}
	return c.JSON(entities.ApiResponse{
		Success: true,
		Message: "ลบข้อมูลสำเร็จ",
	})
}
