// Copyright 2021-2022, Offchain Labs, Inc.
// For license information, see https://github.com/nitro/blob/master/LICENSE

package stopwaiter

import (
	"context"
	"errors"
	"reflect"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/log"
)

const stopDelayWarningTimeout = 30 * time.Second

type StopWaiterSafe struct {
	mutex     sync.Mutex // protects started, stopped, ctx, parentCtx, stopFunc
	started   bool
	stopped   bool
	ctx       context.Context
	parentCtx context.Context
	stopFunc  func()
	name      string
	waitChan  <-chan interface{}

	wg sync.WaitGroup
}

func (s *StopWaiterSafe) Started() bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.started
}

func (s *StopWaiterSafe) Stopped() bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.stopped
}

func (s *StopWaiterSafe) GetContextSafe() (context.Context, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.getContext()
}

// this context is not cancelled even after someone calls Stop
func (s *StopWaiterSafe) GetParentContextSafe() (context.Context, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.getParentContext()
}

// Only call this internally with the mutex held.
func (s *StopWaiterSafe) getContext() (context.Context, error) {
	if s.started {
		return s.ctx, nil
	}
	return nil, errors.New("not started")
}

// Only call this internally with the mutex held.
func (s *StopWaiterSafe) getParentContext() (context.Context, error) {
	if s.started {
		return s.parentCtx, nil
	}
	return nil, errors.New("not started")
}

func getParentName(parent any) string {
	// remove asterisk in case the type is a pointer
	return strings.Replace(reflect.TypeOf(parent).String(), "*", "", 1)
}

// start-after-start will error, start-after-stop will immediately cancel
func (s *StopWaiterSafe) Start(ctx context.Context, parent any) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.started {
		return errors.New("start after start")
	}
	s.started = true
	s.name = getParentName(parent)
	s.parentCtx = ctx
	s.ctx, s.stopFunc = context.WithCancel(s.parentCtx)
	if s.stopped {
		s.stopFunc()
	}
	return nil
}

func (s *StopWaiterSafe) StopOnly() {
	_ = s.stopOnly()
}

// returns true if stop function was called
func (s *StopWaiterSafe) stopOnly() bool {
	stopWasCalled := false
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.started && !s.stopped {
		s.stopFunc()
		stopWasCalled = true
	}
	s.stopped = true
	return stopWasCalled
}

// StopAndWait may be called multiple times, even before start.
func (s *StopWaiterSafe) StopAndWait() error {
	return s.stopAndWaitImpl(stopDelayWarningTimeout)
}

func getAllStackTraces() string {
	buf := make([]byte, 64*1024*1024)
	size := runtime.Stack(buf, true)
	builder := strings.Builder{}
	builder.Write(buf[0:size])
	return builder.String()
}

func (s *StopWaiterSafe) stopAndWaitImpl(warningTimeout time.Duration) error {
	if !s.stopOnly() {
		return nil
	}
	waitChan, err := s.GetWaitChannel()
	if err != nil {
		return err
	}
	timer := time.NewTimer(warningTimeout)

	select {
	case <-timer.C:
		traces := getAllStackTraces()
		log.Warn("taking too long to stop", "name", s.name, "delay[s]", warningTimeout.Seconds())
		log.Warn(traces)
	case <-waitChan:
		timer.Stop()
		return nil
	}
	<-waitChan
	return nil
}

func (s *StopWaiterSafe) GetWaitChannel() (<-chan interface{}, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.waitChan == nil {
		ctx, err := s.getContext()
		if err != nil {
			return nil, err
		}
		waitChan := make(chan interface{})
		go func() {
			<-ctx.Done()
			s.wg.Wait()
			close(waitChan)
		}()
		s.waitChan = waitChan
	}
	return s.waitChan, nil
}

// If stop was already called, thread might silently not be launched
func (s *StopWaiterSafe) LaunchThreadSafe(foo func(context.Context)) error {
	ctx, err := s.GetContextSafe()
	if err != nil {
		return err
	}
	if s.Stopped() {
		return nil
	}
	s.wg.Add(1)
	go func() {
		foo(ctx)
		s.wg.Done()
	}()
	return nil
}

// This calls go foo() directly, with the benefit of being easily searchable.
// Callers may rely on the assumption that foo runs even if this is stopped.
func (s *StopWaiterSafe) LaunchUntrackedThread(foo func()) {
	go foo()
}

// CallIteratively calls function iteratively in a thread.
// input param return value is how long to wait before next invocation
func (s *StopWaiterSafe) CallIterativelySafe(foo func(context.Context) time.Duration) error {
	return s.LaunchThreadSafe(func(ctx context.Context) {
		for {
			interval := foo(ctx)
			if ctx.Err() != nil {
				return
			}
			if interval == time.Duration(0) {
				continue
			}
			timer := time.NewTimer(interval)
			select {
			case <-ctx.Done():
				timer.Stop()
				return
			case <-timer.C:
			}
		}
	})
}

// StopWaiter may panic on race conditions instead of returning errors
type StopWaiter struct {
	StopWaiterSafe
}

func (s *StopWaiter) Start(ctx context.Context, parent any) {
	if err := s.StopWaiterSafe.Start(ctx, parent); err != nil {
		panic(err)
	}
}

func (s *StopWaiter) StopAndWait() {
	if err := s.StopWaiterSafe.StopAndWait(); err != nil {
		panic(err)
	}
}

// If stop was already called, thread might silently not be launched
func (s *StopWaiter) LaunchThread(foo func(context.Context)) {
	if err := s.LaunchThreadSafe(foo); err != nil {
		panic(err)
	}
}

func (s *StopWaiter) CallIteratively(foo func(context.Context) time.Duration) {
	if err := s.CallIterativelySafe(foo); err != nil {
		panic(err)
	}
}

func (s *StopWaiter) GetContext() context.Context {
	ctx, err := s.GetContextSafe()
	if err != nil {
		panic(err)
	}
	return ctx
}

func (s *StopWaiter) GetParentContext() context.Context {
	ctx, err := s.GetParentContextSafe()
	if err != nil {
		panic(err)
	}
	return ctx
}
