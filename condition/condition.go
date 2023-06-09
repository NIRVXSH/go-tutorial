package main

import (
	"errors"
	"fmt"
)

func genErr() error { //func genErr() error: เป็นฟังก์ชันที่ไม่รับพารามิเตอร์และส่งคืนค่าเป็น error ซึ่งเป็นประเภทข้อมูลที่ใช้ใน Go เพื่อระบุข้อผิดพลาดหรือข้อความเกิดขึ้นในระหว่างการประมวลผล
	return errors.New("some error")
	//return errors.New("some error"): ในฟังก์ชัน genErr() นี้มีคำสั่ง return ที่ส่งคืนค่าเป็นข้อผิดพลาดที่ถูกสร้างขึ้น
	//โดยใช้ฟังก์ชัน errors.New() ซึ่งรับพารามิเตอร์เป็นสตริง (string) และสร้าง error ขึ้นมาใหม่ด้วยข้อความที่ระบุในพารามิเตอร์นั้น ในที่นี้คือ "some error".
}

func main() {
	a := 2
	if a != 2 {
		fmt.Println("is not 2")
	} else {
		fmt.Println("a value is :", a)
	}

	score := 80
	if score >= 80 {
		fmt.Println("Ur grade is A")
	} else if score >= 70 {
		fmt.Println("Ur grade is B")
	} else if score >= 60 {
		fmt.Println("Ur grade is C")
	} else if score >= 50 {
		fmt.Println("Ur grade is D")
	} else {
		fmt.Println("Ur grade is F")
	}

	if err := genErr(); err != nil { //err := genErr(): ในบรรทัดนี้ เราใช้เครื่องหมาย := เพื่อประกาศตัวแปร err และใช้ฟังก์ชัน genErr() เพื่อเรียกใช้และรับค่าข้อผิดพลาดที่สร้างขึ้นจากฟังก์ชันนั้น.
		//err != nil: เป็นการตรวจสอบว่าค่าข้อผิดพลาด err ไม่เป็น nil ซึ่งหมายความว่ามีข้อผิดพลาดที่เกิดขึ้น. **err เรียกใช้นอก if ไม่ได้ **
		panic(err.Error()) //panic(err.Error()): ถ้ามีข้อผิดพลาดเกิดขึ้น (ค่า err ไม่เป็น nil) โค้ดจะเรียกใช้ panic()
		//เพื่อทำให้โปรแกรมสิ้นสุดการทำงานทันทีและระบุข้อผิดพลาดที่เกิดขึ้นในรูปแบบข้อความโดยใช้ err.Error() เพื่อรับข้อความจาก error นั้น.
	}
}
