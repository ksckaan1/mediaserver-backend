package auth

import (
	"bff-service/internal/core/models"
	"bff-service/internal/pkg/sessionutils"
	"shared/enums/usertype"
	"shared/pb/authpb"
	"shared/pb/userpb"

	"github.com/gofiber/fiber/v2"
)

type AuthMiddleware struct {
	authClient authpb.AuthServiceClient
	userClient userpb.UserServiceClient
}

func NewAuthMiddleware(authClient authpb.AuthServiceClient, userClient userpb.UserServiceClient) *AuthMiddleware {
	return &AuthMiddleware{
		authClient: authClient,
		userClient: userClient,
	}
}

type mwCookie struct {
	SessionID string `cookie:"session_id"`
}

func (m *AuthMiddleware) Handle(c *fiber.Ctx) error {
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

	user, err := m.userClient.GetUserByID(c.UserContext(), &userpb.GetUserByIDRequest{
		Id: session.UserId,
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
		UserType:  usertype.FromString(user.UserType),
	})

	return c.Next()
}
