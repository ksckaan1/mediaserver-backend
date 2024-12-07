package main

import (
	"context"
	"mediaserver/internal/pkg/generichandler"

	"github.com/gofiber/fiber/v2"
)

func H[I, O any](gh func(context.Context, *generichandler.Request[I]) (*generichandler.Response[O], error)) fiber.Handler {
	return func(c *fiber.Ctx) error {
		icBody := new(I)
		if len(c.Body()) > 0 {
			err := c.BodyParser(icBody)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": err.Error(),
				})
			}
		}

		req := generichandler.NewRequestContainer(*icBody, c.AllParams(), c.Queries(), c.GetReqHeaders())

		out, err := gh(c.UserContext(), req)
		if err != nil {
			return c.Status(getStatusCode(out, 500)).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(getStatusCode(out, 200)).JSON(out.Body)
	}
}

func getStatusCode[O any](out *generichandler.Response[O], defaultStatusCode int) int {
	if out != nil && out.StatusCode != 0 {
		return out.StatusCode
	}
	return defaultStatusCode
}
