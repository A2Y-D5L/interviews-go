package main

import (
	"context"
	"fmt"
)

// filterEvenNumbers filters out even numbers.
func filterEvenNumbers(ctx context.Context, numbers []int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for idx, n := range numbers {
			select {
			case <-ctx.Done():
				return
			default:
				fmt.Println(idx, ": ", n)
				if n%2 != 0 {
					out <- n
				}
			}
		}
	}()
	return out
}

// sumNumbers receives numbers and sums them.
func sumNumbers(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		sum := 0
		for n := range in {
			sum += n
		}
		out <- sum
	}()
	return out
}

// monitorPipeline monitors the pipeline and returns the calculated sum.
func monitorPipeline(ctx context.Context, sumStage <-chan int) int {
	sum := 0
	for {
		select {
		case val, ok := <-sumStage:
			if !ok {
				return sum
			}
			sum += val
		case <-ctx.Done():
			return sum
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	pipeline := sumNumbers(ctx, filterEvenNumbers(ctx, []int{3, 4, 6, 1, 22, 59, 784, 32, 121}))
	fmt.Println("Expected:", 3+1+59+121)
	fmt.Println("Result:", monitorPipeline(ctx, pipeline))
}
