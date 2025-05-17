package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name     string
		size     int
		expected int
	}{
		{
			name:     "Положительный размер",
			size:     10,
			expected: 10,
		},
		{
			name:     "Нулевой размер",
			size:     0,
			expected: 0,
		},
		{
			name:     "Отрицательный размер",
			size:     -10,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := generateRandomElements(tt.size)
			require.Len(t, result, tt.expected)
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name       string
		data       []int
		expected   int
		shouldFail bool
	}{
		{
			name:     "Один элемент",
			data:     []int{5},
			expected: 5,
		},
		{
			name:     "Все элементы одинаковые",
			data:     []int{3, 3, 3},
			expected: 3,
		},
		{
			name:     "Максимум в начале",
			data:     []int{10, 1, 2, 3},
			expected: 10,
		},
		{
			name:     "Максимум в середине",
			data:     []int{4, 7, 99, 2, 5},
			expected: 99,
		},
		{
			name:     "Максимум в конце",
			data:     []int{1, 2, 3, 100},
			expected: 100,
		},
		{
			name:     "Пустой слайс",
			data:     []int{},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maximum(tt.data)
			require.Equal(t, tt.expected, result)
		})
	}
}
