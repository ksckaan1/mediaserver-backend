package logger

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLogger(t *testing.T) {
	logger, err := New()
	require.NoError(t, err)

	logger.Info(context.TODO(), "test", "key1", "field1", "key2", 42)
}
