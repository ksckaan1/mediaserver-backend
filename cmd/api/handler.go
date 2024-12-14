package main

import (
	"context"
	"fmt"
	"mediaserver/internal/pkg/gh"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/ksckaan1/m2s"
)

func H[I, O any](genericHandler func(context.Context, *gh.Request[I]) (*gh.Response[O], error)) fiber.Handler {
	return func(c *fiber.Ctx) error {
		body, err := parseBody[I](c)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		req := gh.NewRequestContainer(body, c.AllParams(), c.Queries(), c.GetReqHeaders())
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

func parseBody[I any](c *fiber.Ctx) (I, error) {
	body := new(I)
	if len(c.Body()) == 0 {
		return *body, nil
	}
	switch {
	case c.Get("Content-Type") == "application/json":
		err := c.BodyParser(body)
		if err != nil {
			return *body, err
		}
	case strings.HasPrefix(c.Get("Content-Type"), "multipart/form-data"):
		mpf, err := c.MultipartForm()
		if err != nil {
			return *body, err
		}
		err = m2s.Convert(mpf, body)
		if err != nil {
			return *body, err
		}
	}

	err := validateBody(*body)
	if err != nil {
		return *body, err
	}

	return *body, nil
}

func validateBody(body any) error {
	validate := validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	err := validate.Struct(body)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return fmt.Errorf("validation error: %q field must be %s", validationErrors[0].Field(), validationErrors[0].ActualTag())
	}
	return nil
}
