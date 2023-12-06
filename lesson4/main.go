package main

import "fmt"

type Client struct {
	Name     string
	Age      int
	Products []Product
}

type Product struct {
	Name  string
	Price float64
}

var client = []Client{
	{Name: "Anvar",
		Age: 16,
		Products: []Product{
			{Name: "Olma", Price: 123000},
			{Name: "Olcha", Price: 122000},
		}},
	{Name: "Anvar",
		Age: 16,
		Products: []Product{
			{Name: "Olma", Price: 123000},
			{Name: "Olcha", Price: 122000},
		}},
}

func main() {

	//var abc = make(map[int]Product)

	for _, value := range client {

		for _, v := range value.Products {

			fmt.Println(v.Name)
		}

	}

}
