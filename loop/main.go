package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("ค่าloop แบบ for รอบที่ ", i)
	}

	i := 0
	for i < 10 {
		fmt.Println("ค่าloop แบบ while รอบที่ ", i)
		i++
	}
}
