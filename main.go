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
		panic("empty slice")
	}

	max := 0
	for i := 0; i < len(data); i++ {
		if data[i] > max {
			max = data[i]
		}
	}
	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) == 0 {
		panic("empty slice")
	}

	chunkSize := len(data) / CHUNKS
	if chunkSize == 0 {
		return maximum(data)
	}

	var wg sync.WaitGroup
	results := make(chan int, CHUNKS)

	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == CHUNKS-1 {
			end = len(data)

			wg.Add(1)
			go func(slice []int) {
				defer wg.Done()
				results <- maximum(slice)
			}(data[start:end])
		}

		go func() {
			wg.Wait()
			close(results)
		}()
	}
	max := <-results
	for res := range results {
		if res > max {
			max = res
		}
	}

	return max
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
