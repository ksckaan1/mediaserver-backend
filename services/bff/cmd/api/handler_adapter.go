package main

import (
	"context"
	"strings"

	validator "github.com/go-playground/validator/v10"
	fiber "github.com/gofiber/fiber/v2"
	m2s "github.com/ksckaan1/m2s"
)

type Handler[Req, Resp any] interface {
	Handle(ctx context.Context, req *Req) (*Resp, int, error)
}

func h[Req, Resp any](handler Handler[Req, Resp]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := new(Req)

		err := parseRequest(c, req)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		validate := validator.New()
		err = validate.Struct(req)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		resp, statusCode, err := handler.Handle(c.UserContext(), req)
		if err != nil {
			return c.Status(statusCode).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if statusCode == fiber.StatusNoContent {
			return c.SendStatus(statusCode)
		}

		return c.Status(statusCode).JSON(resp)
	}
}

func parseRequest(c *fiber.Ctx, req any) error {
	contentType := c.Get("Content-Type")
	if contentType == "application/json" {
		err := c.BodyParser(req)
		if err != nil {
			return err
		}
	} else if strings.HasPrefix(contentType, "multipart/form-data") {
		mpf, err := c.MultipartForm()
		if err != nil {
			return err
		}
		err = m2s.Convert(mpf, req)
		if err != nil {
			return err
		}
	}
	err := c.ParamsParser(req)
	if err != nil {
		return err
	}
	err = c.QueryParser(req)
	if err != nil {
		return err
	}
	err = c.ReqHeaderParser(req)
	if err != nil {
		return err
	}
	return nil
}
