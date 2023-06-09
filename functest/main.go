package main

import (
	"fmt"

	"github.com/NIRVXSH/go-basic/functest/module"
)

func sum(name string, a int, b int) (int, string) {
	return a + b, "Hello :" + name
}

func changer(x int) {
	x = 20
}

func main() {
	fmt.Println(sum("Fook", 10, 20))
	x := 10
	changer(x)
	fmt.Println("x เรียกใช้ changer แล้วค่า x จะไม่เปลี่ยน เป็น x =10 เหมือนเดิม : ", x)

	fmt.Println("ค่าจาก func Sum ในโฟลเดอร์ module: ", module.Sum(44, 28))
}
