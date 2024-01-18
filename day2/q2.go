package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(200)                   // add 200 students to wait group
	rating := make(chan int, 200) // buffered channel for 200 students
	for i := 0; i < 200; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond) // to generate random response
			student := rand.Intn(10)                                     // rating between 1 and 10
			rating <- student
		}()
	}

	wg.Wait()
	close(rating) // to avoid deadlock  fatal error: all goroutines are asleep - deadlock! this error was coming

	var total int
	for r := range rating {
		total += r
	}

	averageRating := float64(total) / 200.0
	fmt.Printf("Average rating is : %.2f\n", averageRating)
}
