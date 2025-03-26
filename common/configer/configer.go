package configer

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type Configer[T any] struct {
	Data T
}

func New[T any]() *Configer[T] {
	return &Configer[T]{
		Data: *new(T),
	}
}

func (c *Configer[T]) Load() error {
	err := env.Parse(c)
	if err != nil {
		return fmt.Errorf("env.Parse: %w", err)
	}
	return nil
}
