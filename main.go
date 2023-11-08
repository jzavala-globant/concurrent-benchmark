package main

import (
	"sync"
)

const (
	numOfGorutines = 50
	desiredPos     = 30
	numOfWorkers   = 3
)

func fib(desiredPos int) int {
	if desiredPos <= 1 {
		return 0
	}
	if desiredPos == 2 {
		return 1
	}
	return fib(desiredPos-1) + fib(desiredPos-2)
}

// worker function takes in an id, a jobs channel, and a results channel
func worker(id int, jobs <-chan int, results chan<- int) {
	// iterates over the jobs channel
	for j := range jobs {
		// fmt.Println("worker", id, "started job", j)
		res := fib(desiredPos)
		_ = res
		_ = j
		// fmt.Println("worker", id, "finished job", j, "res:", res)
		// sends the result to the results channel
		results <- res
	}
}

func f1(iterations int) {
	wg := sync.WaitGroup{}
	for i := 1; i <= iterations; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			// log.Printf("Gorutine(%d started)\n", i)
			res := fib(desiredPos)
			_ = res
			// log.Printf("Gorutine(%d) completed: %d\n", i, res)
		}(i)
	}
	wg.Wait()
}

func f2(iterations int) {
	// creates a jobs channel with a buffer size of numJobs
	jobs := make(chan int, iterations)
	// creates a results channel with a buffer size of numJobs
	results := make(chan int, iterations)

	// creates 3 goroutines for the worker function
	for w := 1; w <= numOfWorkers; w++ {

		go worker(w, jobs, results)
	}
	// sends numJobs integers to the jobs channel
	for j := 1; j <= iterations; j++ {
		jobs <- j
	}
	// close the jobs channel to signal that all jobs have been sent
	close(jobs)

	// receives numJobs integers from the results channel
	for a := 1; a <= iterations; a++ {
		<-results
	}
	// close the results channel to signal that all results received
	close(results)
}

func main() {
	// f1(numOfGorutines)
	f2(numOfGorutines)
}
