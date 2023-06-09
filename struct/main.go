package main

import (
	"encoding/json"
	"fmt"

	models "github.com/NIRVXSH/go-basic/struct/model"
)

type book struct {
	id     int
	title  string
	author string
}

func Debug(obj any) {
	raw, _ := json.MarshalIndent(&obj, "", "\t")
	fmt.Println(string(raw))
}

type book1 struct {
	*models.Book
}

func (d *book1) GetBook() *book1 {
	return d
}

func (d book1) SetBook(title string) {
	d.Title = title
}

func main() {
	b := book{
		id:     1,
		title:  "Golang Basic",
		author: "NIRVXSH",
	}
	fmt.Println(b)

	c := models.Book{
		Id:     2,
		Title:  "Sedtanan",
		Author: "Feel special",
	}
	fmt.Println(c)
	Debug(c)

	d := book1{
		&models.Book{
			Id:     3,
			Title:  "test",
			Author: "struct",
		}}
	d.GetBook()
	Debug(d)
	d.SetBook("Exam")
	Debug(d)
}
