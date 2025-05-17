package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Int()
	}
	return arr
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	max := data[0]
	for i := range data {
		if data[i] > max {
			max = data[i]
		}
	}
	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	n := len(data)
	chunkSize := n / CHUNKS

	var wg sync.WaitGroup
	var mu sync.Mutex
	results := make([]int, 0, CHUNKS)

	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		end := (i + 1) * chunkSize
		if end > n {
			end = n
		}

		wg.Add(1)
		go func(slice []int) {
			defer wg.Done()
			localMax := maximum(slice)

			mu.Lock()
			results = append(results, localMax)
			mu.Unlock()
		}(data[start:end])
	}

	wg.Wait()

	return maximum(results)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	arr := generateRandomElements(SIZE)

	timeFirstWay := time.Now()
	fmt.Println("Ищем максимальное значение в один поток")
	max := maximum(arr)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, time.Since(timeFirstWay).Microseconds())

	timeSecondWay := time.Now()
	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	max = maxChunks(arr)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, time.Since(timeSecondWay).Microseconds())
}
