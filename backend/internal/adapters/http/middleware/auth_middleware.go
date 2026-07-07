package middleware

import (
	"morng-dev/internal/core/domain/entities"
	"morng-dev/pkg/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type AuthMiddleware struct {
	JWT_SECRET string
}

func NewAuthMiddleware(JWT_SECRET string) *AuthMiddleware {
	return &AuthMiddleware{
		JWT_SECRET: JWT_SECRET,
	}
}

func (m *AuthMiddleware) Authrequire() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(entities.ApiResponse{
				Success: false,
				Message: "header Unauthorization",
			})
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			return c.Status(fiber.StatusUnauthorized).JSON(entities.ApiResponse{
				Success: false,
				Message: "รูปแบบ token ไม่ถูกต้อง",
			})
		}
		token, err := jwt.ParseWithClaims(tokenString, &utils.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(m.JWT_SECRET), nil
		})
		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(entities.ApiResponse{
				Success: false,
				Message: "token ไม่ถูกต้องหรือหมดอายุ",
			})
		}
		claims, ok := token.Claims.(*utils.Claims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(entities.ApiResponse{
				Success: false,
				Message: "Claims ไม่ถูกต้อง",
			})
		}
		userID, err := uuid.Parse(claims.UserId)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(entities.ApiResponse{
				Success: false,
				Message: "userID invalid",
			})
		}
		c.Locals("userID", userID)
		c.Locals("email", claims.Email)

		return c.Next()

	}
}
