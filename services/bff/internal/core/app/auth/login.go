package auth

import (
	"shared/pb/authpb"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type Login struct {
	authClient authpb.AuthServiceClient
}

func NewLogin(authClient authpb.AuthServiceClient) *Login {
	return &Login{
		authClient: authClient,
	}
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	SessionId string `json:"session_id"`
}

func (h *Login) Handle(c *fiber.Ctx) error {
	var body LoginRequest
	err := c.BodyParser(&body)
	if err != nil {
		return err
	}

	resp, err := h.authClient.Login(c.UserContext(), &authpb.LoginRequest{
		Username:  body.Username,
		Password:  body.Password,
		UserAgent: c.Get("User-Agent"),
		IpAddress: c.IP(),
	})
	if err != nil {
		if strings.Contains(err.Error(), "password does not match") ||
			strings.Contains(err.Error(), "user not found") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "wrong username or password",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "session_id",
		Value:    resp.SessionId,
		HTTPOnly: true,
		Secure:   true,
	})

	return c.SendStatus(fiber.StatusNoContent)
}
