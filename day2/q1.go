package main

import (
	"fmt"
	"sync"
)

func count(input []string) map[string]int {
	freq := make(map[string]int) // have to try once with slice as well
	var wg sync.WaitGroup
	var mu sync.Mutex
	for _, str := range input {
		wg.Add(1)
		go func(str string) {
			defer wg.Done()           // to keep going the go routines
			m := make(map[string]int) // local map for storing each string frequency

			for _, char := range str {
				char := string(char)
				m[char]++
			}
			mu.Lock()
			for char, count := range m {
				freq[char] += count
			}
			mu.Unlock()
		}(str)

	}
	wg.Wait()
	return freq
}
func main() {
	var s string
	i := 0
	for i < 100000 {
		s = s + "z"
		i++
	}
	fmt.Println(s)
	input := []string{"quick", "brown", "fox", "lay", "dog", s}
	frequencies := count(input)
	fmt.Println(frequencies)
}
