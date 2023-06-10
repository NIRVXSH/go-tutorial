package main

import (
	"fmt"
	"sync"
)

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
	//chan
	//Worker Pool
}
