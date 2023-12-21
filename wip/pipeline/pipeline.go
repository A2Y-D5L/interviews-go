// Filename: pipeline.go
// This file demonstrates the concurrent pipeline pattern in Go for processing streams of data.
// It filters odd numbers from a slice, sums them up, and outputs the total sum using Go's concurrency primitives.

package main

import (
	"context"
)

// filterOddNumbers returns a channel emitting only the odd numbers from the input slice.
// It listens to the context for cancellation.
func filterOddNumbers(ctx context.Context, numbers []int) <-chan int {
	out := make(chan int, len(numbers)) // Use buffered channel to reduce context switches.
	go func() {
		defer close(out)
		for _, n := range numbers {
			select {
			case <-ctx.Done():
				return
			default:
				if n%2 != 0 {
					out <- n
				}
			}
		}
	}()
	return out
}

// sumNumbers returns a channel emitting the sum of numbers received from the input channel.
// It listens to the context for cancellation.
func sumNumbers(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		sum := 0
		for {
			select {
			case n, ok := <-in:
				if !ok {
					out <- sum
					return
				}
				sum += n
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

// getSumOfOdds returns the sum of odd integers in a slice of int.
func getSumOfOdds(ctx context.Context, ints []int) int {
	pipeline := sumNumbers(ctx, filterOddNumbers(ctx, ints))
	sum := 0
	for {
		select {
		case val, ok := <-pipeline:
			if !ok {
				return sum
			}
			sum += val
		case <-ctx.Done():
			return sum // Return the sum calculated so far.
		}
	}
}
