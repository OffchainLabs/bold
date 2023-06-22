package retry

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/metrics"
	"github.com/sirupsen/logrus"
)

const sleepTime = time.Second * 1

var (
	log          = logrus.WithField("prefix", "util")
	retryCounter = metrics.NewRegisteredCounter("arb/validator/runtime/retry", nil)
)

// UntilSucceeds retries the given function until it succeeds or the context is cancelled.
func UntilSucceeds[T any](ctx context.Context, fn func() (T, error)) (T, error) {
	for {
		if ctx.Err() != nil {
			return zeroVal[T](), ctx.Err()
		}
		got, err := fn()
		if err != nil {
			log.WithError(err).Error("Failed to call function")
			retryCounter.Inc(1)
			time.Sleep(sleepTime)
			continue
		}
		return got, nil
	}
}

func zeroVal[T any]() T {
	var result T
	return result
}
