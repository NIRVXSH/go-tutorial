package main

import (
	"fmt"
	"sync"
)

type numbers struct {
	a int
	b int
}

type sum struct {
	r int
}

func workers(jobCh <-chan numbers, resultsCh chan<- sum) {
	for job := range jobCh {
		resultsCh <- sum{job.a + job.b}
	}
}

func main() {
	//Basic -> sync.Waitgroup
	a := []int{1, 2, 3, 4, 5}

	//for i := range a {
	//	fmt.Printf("%v ", a[i])
	//}

	var wg sync.WaitGroup
	for i := range a {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("%v ", a[i])
		}(i)
	}
	wg.Wait()
	fmt.Println("")
	fmt.Println("----------------------------------------")
	wg.Add(len(a))
	go func() {
		for i := range a {
			defer wg.Done()
			fmt.Printf("%v ", a[i])

		}
	}()
	wg.Wait()
	fmt.Println("")
	fmt.Println("-----------------chan-----------------------")
	fmt.Println("")
	//chan
	numberCh := make(chan int)
	msgCh := make(chan string)
	//number
	go func(numberCh chan int) {
		numberCh <- 10
	}(numberCh)

	//msg
	go func(msgCh chan string) {
		msgCh <- "Hello NIRVXSH"
	}(msgCh)

	number := <-numberCh
	msg := <-msgCh

	fmt.Println(number)
	fmt.Println(msg)

	//Worker Pool
	fmt.Println("")
	fmt.Println("-----------------Worker Pool-----------------------")

	nums := []numbers{
		{a: 1, b: 2},
		{a: 1, b: 2},
		{a: 1, b: 2},
		{a: 1, b: 2},
		{a: 1, b: 2},
	}
	fmt.Println(nums)
	jobCh := make(chan numbers, len(nums))
	resultsCH := make(chan sum, len(nums))

	for _, n := range nums {
		jobCh <- n
	}
	close(jobCh)

	numWorkers := 2
	for w := 0; w < numWorkers; w++ {
		go workers(jobCh, resultsCH)
	}
	results := make([]sum, 0)
	for a := 0; a < len(nums); a++ {
		temp := <-resultsCH
		results = append(results, temp)
	}
	fmt.Println(results)
}
