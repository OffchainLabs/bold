// Copyright 2023-2024, Offchain Labs, Inc.
// For license information, see:
// https://github.com/offchainlabs/bold/blob/main/LICENSE.md

// Package runtime defines utilities that deal with managing lifecycles of
// functions and important behaviors at the application runtime, such as
// retrying errored functions until they succeed.
package retry

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/metrics"
)

const defaultSleepTime = time.Second * 30

var (
	retryCounter = metrics.NewRegisteredCounter("arb/validator/runtime/retry", nil)
)

type RetryConfig struct {
	sleepTime time.Duration
}

type Opt func(*RetryConfig)

// WithInterval specifies how often to retry an errored function.
func WithInterval(d time.Duration) Opt {
	return func(rc *RetryConfig) {
		rc.sleepTime = d
	}
}

func UntilSucceedsMultipleReturnValue[T, U any](ctx context.Context, fn func() (T, U, error), opts ...Opt) (T, U, error) {
	cfg := &RetryConfig{
		sleepTime: defaultSleepTime,
	}
	for _, o := range opts {
		o(cfg)
	}
	count := 0
	for {
		if ctx.Err() != nil {
			return zeroVal[T](), zeroVal[U](), ctx.Err()
		}
		got, got2, err := fn()
		if err != nil {
			count++
			log.Error("Could not succeed function after retries",
				"retryCount", count,
				"err", err,
			)
			retryCounter.Inc(1)
			select {
			case <-ctx.Done():
				return zeroVal[T](), zeroVal[U](), ctx.Err()
			case <-time.After(cfg.sleepTime):
			}
			continue
		}
		return got, got2, nil
	}
}

// UntilSucceeds retries the given function until it succeeds or the context is cancelled.
func UntilSucceeds[T any](ctx context.Context, fn func() (T, error), opts ...Opt) (T, error) {
	result, _, err := UntilSucceedsMultipleReturnValue(ctx, func() (T, struct{}, error) {
		got, err := fn()
		return got, struct{}{}, err
	}, opts...)
	return result, err
}

func zeroVal[T any]() T {
	var result T
	return result
}
