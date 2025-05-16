package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	size := 100
	arr := generateRandomElements(100)
	lenOfArr := len(arr)
	if lenOfArr <= 0 || lenOfArr < size {
		t.Errorf("ожидаемый размер %d, получено %d", size, lenOfArr)
	}

	for _, v := range arr {
		if v <= 0 {
			t.Errorf("ожидаем положительное число, получено %d", v)
		}
	}
}

func TestMaximum(t *testing.T) {
	size := 100
	arr := generateRandomElements(size)
	max := maximum(arr)
	if max <= 0 {
		t.Errorf("ожидаем положительное число, получено %d", max)
	}

	if len(arr) == 0 {
		t.Errorf("ожидаемый размер %d, получено %d", size, 0)
	}

	if len(arr) == 1 && arr[0] != max {
		t.Errorf("неверно определен максимум с одним элементом в списке, получено: %d, ожидалось: %d", max, arr[0])
	}

	arr = []int{5, 5, 5}
	if maximum(arr) != 5 {
		t.Errorf("неверно определен максимум списка из одинаковых элементов, получено: %d, ожидалось: %d", maximum(arr), 5)
	}

	arr = []int{10, 5, 5}
	if maximum(arr) != 10 {
		t.Errorf("максимум в начале списка определен неверно, получено: %d, ожидалось: %d", maximum(arr), 10)
	}

	arr = []int{5, 5, 10}
	if maximum(arr) != 10 {
		t.Errorf("максимум в конце списка определен неверно, получено: %d, ожидалось: %d", maximum(arr), 10)
	}
}
