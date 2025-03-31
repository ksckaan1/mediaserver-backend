package logger

import (
	"context"
	"testing"
)

func TestLogger(t *testing.T) {
	l := New()

	l.Info(context.TODO(), "test", "key1", "field1", "key2", 42)
}
