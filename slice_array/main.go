package main

import "fmt"

func arrayWatcher(data *[]int) {
	//data = append(data, 6) ใช้ append ไม่ได้
	*data = append(*data, 6)
	//data[0] = 0
}

func main() {
	//array ขยายขนาด array ไม่ได้
	a := [3]int{1, 2, 3}
	fmt.Println(a[1])
	b := [11]int{}
	for i := 0; i <= 10; i++ {
		b[i] = i
	}
	fmt.Println(b)
	fmt.Println("----------------------------------------")
	//slice เหมือน array แต่ขยายได้
	c := []int{1, 2, 3, 4}
	c = append(c, 5)
	fmt.Println(c)
	//arrayWatcher(c) ใช้กับ data[0] = 0 ใน arrayWatcher
	arrayWatcher(&c)
	fmt.Println(c)
	fmt.Println("----------------------------------------")
	d := make([]int, 0)
	d = append(d, 5)
	fmt.Println(d)
	//arrayWatcher(c) ใช้กับ data[0] = 0 ใน arrayWatcher
	arrayWatcher(&d)
	fmt.Println(d)
	fmt.Println("----------------------------------------")

	e := []string{"hi", "I'm ", "Fook"}
	for i := range e {
		fmt.Println(a[i])
	}
	fmt.Println("print value in slice")
	for _, V := range e {
		fmt.Println(V)
	}
}
