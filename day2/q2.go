package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func takeRatings(wg *sync.WaitGroup, ratings chan<- int) {
	defer wg.Done()
	time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond) // to generate random response
	studentRating := rand.Intn(10)                               // rating between 1 and 10
	ratings <- studentRating
}
func main() {
	var wg sync.WaitGroup
	wg.Add(200)                    // add 200 students to wait group
	ratings := make(chan int, 200) // buffered channel for 200 students
	for i := 0; i < 200; i++ {
		go takeRatings(&wg, ratings)
	}

	wg.Wait()
	close(ratings) // to avoid deadlock  fatal error: all goroutines are asleep - deadlock! this error was coming

	var total int
	for r := range ratings {
		total += r
	}

	averageRating := float64(total) / 200.0
	fmt.Printf("Average rating is : %.2f\n", averageRating)
}
