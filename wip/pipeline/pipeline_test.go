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

func Test_concPipelines(t *testing.T) {
	tests := []struct {
		name string
		ints [][]int
		want int
	}{
		{
			name: "single slice",
			ints: [][]int{{1, 2, 3, 4}},
			want: 4,
		},
		{
			name: "multiple slices",
			ints: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			want: 25,
		},
		{
			name: "empty slice",
			ints: [][]int{{}},
			want: 0,
		},
		{
			name: "multiple empty slices",
			ints: [][]int{{}, {}, {}},
			want: 0,
		},
		{
			name: "all even numbers",
			ints: [][]int{{2, 4, 6}, {8, 10, 12}},
			want: 0,
		},
		{
			name: "negative numbers",
			ints: [][]int{{-1, -3, -5}},
			want: -9,
		},
		{
			name: "mixed positive and negative numbers",
			ints: [][]int{{1, -2, 3, -4}},
			want: 4,
		},
		{
			name: "large number of slices",
			ints: func() [][]int {
				var slices [][]int
				for i := 0; i < 100; i++ {
					slices = append(slices, []int{1, 2, 3, 4, 5})
				}
				return slices
			}(),
			want: 900,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := concPipelines(tt.ints...); got != tt.want {
				t.Errorf("concPipelines() = %v, want %v", got, tt.want)
			}
		})
	}
}
