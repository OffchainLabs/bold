package validator

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_retryUntilSucceeds(t *testing.T) {
	hello := func() (string, error) {
		return "hello", nil
	}

	ctx := context.Background()
	got, err := retryUntilSucceeds(ctx, hello)
	require.NoError(t, err)
	require.Equal(t, "hello", got)

	newCtx, cancel := context.WithCancel(ctx)
	cancel()
	_, err = retryUntilSucceeds(newCtx, hello)
	require.ErrorContains(t, err, "context canceled")
}
