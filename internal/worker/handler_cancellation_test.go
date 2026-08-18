package worker

import (
	"context"
	"errors"
	"testing"
	"time"

	"sanitation-operations/internal/clock"
	"sanitation-operations/internal/repository"
)

func TestRunnerCancelsInFlightHandler(t *testing.T) {
	store, enqueueCtx, now := workerStore(t)
	if _, err := store.Enqueue(enqueueCtx, "trip.started", []byte(`{}`), now); err != nil {
		t.Fatal(err)
	}
	started := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())
	runner := Runner{Store: store, Clock: clock.Fixed{Current: now}, Interval: time.Hour, Handler: HandlerFunc(func(ctx context.Context, _ repository.OutboxJob) error {
		close(started)
		<-ctx.Done()
		return ctx.Err()
	})}
	done := make(chan error, 1)
	go func() { done <- runner.Run(ctx) }()
	<-started
	cancel()
	select {
	case err := <-done:
		if !errors.Is(err, context.Canceled) {
			t.Fatalf("runner error=%v", err)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("worker handler ignored cancellation")
	}
}
