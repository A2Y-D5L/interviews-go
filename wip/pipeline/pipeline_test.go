package main

import (
	"context"
	"testing"
)

func Test_getSumOfOdds(t *testing.T) {
	tests := []struct {
		name string
		ctx  context.Context
		ints []int
		want int
	}{
		{
			name: "empty slice",
			ctx:  context.Background(),
			ints: []int{},
			want: 0,
		},
		{
			name: "all even numbers",
			ctx:  context.Background(),
			ints: []int{2, 4, 6, 8},
			want: 0,
		},
		{
			name: "all odd numbers",
			ctx:  context.Background(),
			ints: []int{1, 3, 5, 7},
			want: 16,
		},
		{
			name: "mixed numbers",
			ctx:  context.Background(),
			ints: []int{1, 2, 3, 4, 5},
			want: 9,
		},
		{
			name: "negative numbers",
			ctx:  context.Background(),
			ints: []int{-1, -2, -3, -4, -5},
			want: -9,
		},
		{
			name: "large slice",
			ctx:  context.Background(),
			ints: func() []int {
				var largeSlice []int
				for i := 0; i < 1000; i++ {
					largeSlice = append(largeSlice, i)
				}
				return largeSlice
			}(),
			want: 250000, // The sum of all odd numbers from 0 to 999.
		},
		{
			name: "context cancelled",
			ctx: func() context.Context {
				ctx, cancel := context.WithCancel(context.Background())
				cancel() // Immediately cancel the context
				return ctx
			}(),
			ints: []int{1, 3, 5, 7},
			want: 0, // Should return 0 as context is cancelled before processing.
		},
	}

	// Execute the test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSumOfOdds(tt.ctx, tt.ints); got != tt.want {
				t.Errorf("getSumOfOdds() = %v, want %v", got, tt.want)
			}
		})
	}
}
