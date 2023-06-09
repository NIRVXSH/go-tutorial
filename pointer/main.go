package main

import "fmt"

func swap(x, y *int) {
	temp := *x
	*x = *y
	*y = temp
}

func incrementPointer(p *int) {
	*p = *p + 1
}

func main() {
	a := new(int)
	b := 10

	a = &b //ในบรรทัดนี้ เรากำหนดค่าของ a ให้ชี้ไปยังตัวแปร b โดยใช้ตัวดำเนินการ & ที่ใช้ในการเข้าถึงที่อยู่ของตัวแปร.

	fmt.Println("a อ้างอิงไป pointer b :", a) //เราใช้ fmt.Println() เพื่อพิมพ์ค่าของ a ที่เป็น pointer ซึ่งจะแสดงที่อยู่ของ b ที่ a กำหนดให้ชี้ไป.

	fmt.Println("&b อ้างอิงไป pointer b :", &b) //เราใช้ fmt.Println() เพื่อพิมพ์ที่อยู่ของ b โดยใช้ตัวดำเนินการ & ที่ใช้ในการเข้าถึงที่อยู่ของตัวแปร.

	fmt.Println("&a อ้างอิงไป pointer a ", &a) // เราใช้ fmt.Println() เพื่อพิมพ์ที่อยู่ของ a โดยใช้ตัวดำเนินการ & ที่ใช้ในการเข้าถึงที่อยู่ของตัวแปร.

	fmt.Println("*a อ้างอิง Value b ", *a) //เราใช้ fmt.Println() พร้อมกับ %p เพื่อพิมพ์ค่าที่อยู่ของตัวแปร a โดยใช้ตัวผู้รับ *a ที่ใช้ในการอ้างอิงค่าของตัวแปร a ซึ่งในที่นี้เป็นค่าที่อยู่ของ b.

	x := 5
	w := new(int)
	w = &x
	y := 10
	z := new(int)
	z = &y

	fmt.Println("ยังไม่ swap:", x, ",", y)
	swap(w, z)
	fmt.Println("swap:แล้ว", x, ",", y)

	num := 10
	p := &num
	fmt.Println("ค่าก่อนเพิ่ม:", num)
	incrementPointer(p)
	fmt.Println("ค่าหลังเพิ่ม:", num)
}
