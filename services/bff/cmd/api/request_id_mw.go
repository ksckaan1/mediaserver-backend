package main

import (
	"shared/ports"

	fiber "github.com/gofiber/fiber/v2"
	"google.golang.org/grpc/metadata"
)

func requestIDMW(idGenerator ports.IDGenerator) fiber.Handler {
	return func(c *fiber.Ctx) error {
		requestID := c.Get("x-request-id", idGenerator.NewID())
		md := metadata.MD{}
		md.Set("request-id", requestID)
		ctx := metadata.NewOutgoingContext(c.UserContext(), md)
		c.SetUserContext(ctx)
		c.Set("x-request-id", requestID)
		return c.Next()
	}
}
