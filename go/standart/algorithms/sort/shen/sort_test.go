package main

import (
	"reflect"
	"testing"
)

func TestSort72(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "пустой слайс",
			input:    []int{},
			expected: []int{},
		},
		{

			name:     "Одно число",
			input:    []int{1},
			expected: []int{1},
		}, {
			name:     "Уже отсортированный слайс",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Перевернутый слайс",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Нормальная сортировка",
			input:    []int{1, 4, 2, 1, 6, 4, 3, 8, 7, 9},
			expected: []int{1, 1, 2, 3, 4, 4, 6, 7, 8, 9},
		},
		{
			name:     "Все одинаковые",
			input:    []int{2, 2, 2, 2},
			expected: []int{2, 2, 2, 2},
		},
		{
			name:     "Отрицательные числа",
			input:    []int{-3, -1, -2, 0, 2, 1},
			expected: []int{-3, -2, -1, 0, 1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputCopy := make([]int, len(tt.input))
			copy(inputCopy, tt.input)

			result := sort72(inputCopy)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("sort72(%v) = %v; expected %v", tt.input, result, tt.expected)
			}
		})
	}

}

func TestSort73(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Пустой слайс",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Одно число",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "Уже отсортированный",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Перевернутый",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Нормальная сортировка",
			input:    []int{1, 4, 2, 1, 6, 4, 3, 8, 7, 9},
			expected: []int{1, 1, 2, 3, 4, 4, 6, 7, 8, 9},
		},
		{
			name:     "Одинаковые элементы",
			input:    []int{2, 2, 2, 2},
			expected: []int{2, 2, 2, 2},
		},
		{
			name:     "Отрицательные числа",
			input:    []int{-3, -1, -2, 0, 2, 1},
			expected: []int{-3, -2, -1, 0, 1, 2},
		},
		{
			name:     "Большие числа",
			input:    []int{1000, 500, 2000, 1500, 3000},
			expected: []int{500, 1000, 1500, 2000, 3000},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputCopy := make([]int, len(tt.input))
			copy(inputCopy, tt.input)

			result := sort73(inputCopy)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("sortFunc(%v) = %v; expected %v", tt.input, result, tt.expected)
			}
		})
	}
}
