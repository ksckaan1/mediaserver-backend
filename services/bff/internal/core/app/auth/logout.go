package auth

import (
	"bff-service/internal/pkg/sessionutils"
	"shared/pb/authpb"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Logout struct {
	authClient authpb.AuthServiceClient
}

func NewLogout(authClient authpb.AuthServiceClient) *Logout {
	return &Logout{
		authClient: authClient,
	}
}

func (h *Logout) Logout(c *fiber.Ctx) error {
	session, err := sessionutils.GetSession(c.UserContext())
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "session not found",
		})
	}

	all := c.QueryBool("all", false)

	if all {
		_, err = h.authClient.LogoutAll(c.UserContext(), &authpb.LogoutAllRequest{
			UserId: session.UserId,
		})
	} else {
		_, err = h.authClient.Logout(c.UserContext(), &authpb.LogoutRequest{
			SessionId: session.SessionId,
		})
	}
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "session_id",
		Value:    "",
		HTTPOnly: true,
		Secure:   true,
		Expires:  time.Now().Add(-time.Hour * 24),
	})

	return c.SendStatus(fiber.StatusNoContent)
}
