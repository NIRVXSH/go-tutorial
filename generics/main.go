package main

import "fmt"

type Number interface {
	int | float64
}

type GameorMovie interface {
	getPrice() int
}

type Game struct {
	Title    string
	Platform string
	Price    int
}

type Movie struct {
	Title string
	Price int
}

func main() {
	numsInt := []int{1, 2, 3, 4, 5}
	numsFloat64 := []float64{1.1, 2.2, 3.3, 4.4, 5.5}

	fmt.Println(sum(numsInt))
	fmt.Println(sum(numsFloat64))
	fmt.Println("--------------------------------------------")

	games := []Game{
		{
			Title:    "Minecraft",
			Platform: "PC",
			Price:    30,
		},
		{
			Title:    "The Sims 4",
			Platform: "PC",
			Price:    20,
		},
	}

	movies := []Movie{
		{
			Title: "Pacific Rim",
			Price: 30,
		},
		{
			Title: "Harry Potter",
			Price: 10,
		},
	}

	fmt.Println(sumprice(games))
	fmt.Println(sumprice(movies))

}

func sum[V int | float64](nums []V) V { //type int ,float64 ก็ได้
	var sum V
	for _, num := range nums {
		sum += num
	}
	return sum
}

func sumwithtype[V Number](nums []V) V { //type int ,float64 ก็ได้
	var sum V
	for _, num := range nums {
		sum += num
	}
	return sum
}

func sumInt(nums []int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	return sum
}

func sumFloat64(nums []float64) float64 {
	var sum float64
	for _, num := range nums {
		sum += num
	}
	return sum
}

func (g Game) getPrice() int {
	return g.Price
}

func (g Movie) getPrice() int {
	return g.Price
}

func sumprice[V GameorMovie](objs []V) int {
	var result int
	for _, obj := range objs {
		result += obj.getPrice()
	}
	return result
}
