package main

import (
	"context"
	"fmt"
	"mediaserver/internal/pkg/gh"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/ksckaan1/m2s"
)

func H[I, O any](genericHandler func(context.Context, *gh.Request[I]) (*gh.Response[O], error)) fiber.Handler {
	return func(c *fiber.Ctx) error {
		icBody := new(I)
		if len(c.Body()) > 0 && c.Get("Content-Type") == "application/json" {
			err := c.BodyParser(icBody)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": err.Error(),
				})
			}
		} else if len(c.Body()) > 0 && strings.HasPrefix(c.Get("Content-Type"), "multipart/form-data") {
			mpf, err := c.MultipartForm()
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": err.Error(),
				})
			}
			err = m2s.Convert(mpf, icBody)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": fmt.Errorf("convert multipart form to struct: %w", err).Error(),
				})
			}
		}

		req := gh.NewRequestContainer(*icBody, c.AllParams(), c.Queries(), c.GetReqHeaders())

		out, err := genericHandler(c.UserContext(), req)
		if err != nil {
			return c.Status(getStatusCode(out, 500)).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(getStatusCode(out, 200)).JSON(out.Body)
	}
}

func getStatusCode[O any](out *gh.Response[O], defaultStatusCode int) int {
	if out != nil && out.StatusCode != 0 {
		return out.StatusCode
	}
	return defaultStatusCode
}
