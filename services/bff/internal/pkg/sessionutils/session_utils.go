package sessionutils

import (
	"bff-service/internal/core/models"
	"context"
	"errors"

	"github.com/gofiber/fiber/v2"
)

func GetSession(ctx context.Context) (*models.Session, error) {
	session, ok := ctx.Value("session").(*models.Session)
	if !ok {
		return nil, errors.New("session not found")
	}
	return session, nil
}

func SetSession(c *fiber.Ctx, session *models.Session) {
	ctx := context.WithValue(c.UserContext(), "session", session)
	c.SetUserContext(ctx)
}
