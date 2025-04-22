package auth

import (
	"bff-service/internal/pkg/sessionutils"
	"fmt"
	"shared/enums/usertype"

	"github.com/gofiber/fiber/v2"
)

type UserTypeMiddleware struct {
}

func NewUserTypeMiddleware() *UserTypeMiddleware {
	return &UserTypeMiddleware{}
}

func (m *UserTypeMiddleware) RequiredUserType(userType usertype.UserType) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		session, err := sessionutils.GetSession(c.UserContext())
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "unauthorized",
			})
		}
		if session.UserType != userType {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": fmt.Sprintf("user type %s is required", userType.String()),
			})
		}
		return c.Next()
	}
}
