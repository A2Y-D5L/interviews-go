package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"
)

// customContext carries both the standard context and an error channel.
type customContext struct {
	context.Context
	errChan    chan error
	cancelFunc context.CancelFunc
}

// newCustomContext creates a new customContext with an error channel.
func newCustomContext(ctx context.Context) *customContext {
	cancelCtx, cancelFunc := context.WithCancel(ctx)
	return &customContext{
		Context:    cancelCtx,
		errChan:    make(chan error, 1),
		cancelFunc: cancelFunc,
	}
}

// cancelWithError cancels the context and sends an error to the error channel.
func (c *customContext) cancelWithError(err error) {
	select {
	case c.errChan <- err:
	default: // Avoid blocking if an error is already set.
	}
	c.cancelFunc()
}

// generateNumbers sends numbers into a channel, simulating potential errors.
func generateNumbers(ctx *customContext, max int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 1; i <= max; i++ {
			select {
			case <-ctx.Done():
				return
			case out <- i:
			}
		}
	}()
	return out
}

// filterEvenNumbers filters out even numbers.
func filterEvenNumbers(ctx *customContext, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			if n%2 != 0 {
				select {
				case <-ctx.Done():
					return
				case out <- n:
				}
			}
		}
	}()
	return out
}

// sumNumbers receives numbers and sums them.
func sumNumbers(ctx *customContext, in <-chan int) <-chan int {
	out := make(chan int, 1) // Buffered channel to send the sum
	go func() {
		defer close(out)
		sum := 0
		for n := range in {
			sum += n
		}
		out <- sum // Send sum after processing all numbers or if the loop is exited early
	}()
	return out
}

// newPipeline sets up the pipeline and returns a function to start it with an int argument.
func newPipeline(ctx *customContext) func(int) <-chan int {
	return func(max int) <-chan int {
		firstStage := generateNumbers(ctx, max)
		secondStage := filterEvenNumbers(ctx, firstStage)
		sumStage := sumNumbers(ctx, secondStage)
		return sumStage
	}
}

// monitorPipeline monitors the pipeline and returns the calculated sum and any error that occurs.
func monitorPipeline(ctx *customContext, sumStage <-chan int) (int, error) {
	var pipelineErr error
	done := make(chan struct{}) // Signal completion of error handling goroutine

	go func() {
		select {
		case err := <-ctx.errChan:
			pipelineErr = err
			ctx.cancelFunc()
		case <-ctx.Done():
		}
		close(done) // Signal completion
	}()

	sum := 0
	for val := range sumStage {
		sum += val
	}

	<-done // Wait for the error handling goroutine to signal completion
	return sum, pipelineErr
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ctx := newCustomContext(context.Background())
	defer ctx.cancelFunc()

	startPipeline := newPipeline(ctx)

	sum, err := monitorPipeline(ctx, startPipeline(10))
	if err != nil {
		log.Printf("Pipeline error: %v", err)
	} else {
		fmt.Println("Sum of odd numbers:", sum)
	}
}
