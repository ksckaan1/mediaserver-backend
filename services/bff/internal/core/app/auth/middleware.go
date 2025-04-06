package auth

import (
	"bff-service/internal/core/models"
	"bff-service/internal/pkg/sessionutils"
	"shared/pb/authpb"

	"github.com/gofiber/fiber/v2"
)

type Middleware struct {
	authClient authpb.AuthServiceClient
}

func NewMiddleware(authClient authpb.AuthServiceClient) *Middleware {
	return &Middleware{
		authClient: authClient,
	}
}

type mwCookie struct {
	SessionID string `cookie:"session_id"`
}

func (m *Middleware) Handle(c *fiber.Ctx) error {
	var cookie mwCookie
	err := c.CookieParser(&cookie)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "session not found",
		})
	}

	session, err := m.authClient.GetSession(c.UserContext(), &authpb.GetSessionRequest{
		SessionId: cookie.SessionID,
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "unauthorized",
		})
	}

	sessionutils.SetSession(c, &models.Session{
		SessionId: session.SessionId,
		UserId:    session.UserId,
		UserAgent: session.UserAgent,
		IpAddress: session.IpAddress,
		CreatedAt: session.CreatedAt.AsTime(),
	})

	return c.Next()
}
