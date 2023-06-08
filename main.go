package main

import "fmt"

func main() {
	//var name type =value
	//8bit = 1byte
	//32bit= 4byte
	var intmsg string = "int  มี int8 int16 int32 int64 แตกต่างกันที่ประกาศได้สูงสุดต่างกัน"
	var a int = 2
	var d uint = 4
	var msg string = "Test"

	b := 2
	c := float64(b)
	fmt.Println(intmsg)
	fmt.Println("var a int =2 show:", a)
	fmt.Println("var d uint = 4 show:", d)
	fmt.Println("var msg string show: ", msg)
	fmt.Println("b := 2  show:", b)
	fmt.Println("c:=float64(b) show:", c)

	var e any
	e = "test any"
	fmt.Println("ทดลอง any ด้วย var e any e= test any show: ", e)

	var f any
	f = "แปลง any เป็น String"
	g := f.(string)
	fmt.Println(g)
}
