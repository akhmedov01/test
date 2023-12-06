package main

import "fmt"

type Client struct {
	Id     int
	Name   string
	Basket []BasketProducts
}

type BasketProducts struct {
	ProductId int
	Quantity  int
}

type Product struct {
	Id    int
	Name  string
	Price int
}

var clients = []Client{
	{Id: 1, Name: "Anvar", Basket: []BasketProducts{
		{ProductId: 1, Quantity: 11},
		{ProductId: 2, Quantity: 6},
		{ProductId: 3, Quantity: 9},
	}},
	{Id: 2, Name: "Dilmurod", Basket: []BasketProducts{
		{ProductId: 1, Quantity: 12},
		{ProductId: 2, Quantity: 6},
		{ProductId: 3, Quantity: 9},
	}},
}

var products = []Product{
	{Id: 1, Name: "Olma", Price: 12000},
	{Id: 2, Name: "Banan", Price: 22000},
	{Id: 3, Name: "Olcha", Price: 25000},
}

func main() {

	productMap := make(map[int]int)

	for _, v := range products {
		productMap[v.Id] = v.Price
	}

	//fmt.Println(productMap[3])

	for _, v := range clients {

		sum := 0

		fmt.Print("Client Name: ", v.Name)

		for _, vBasket := range v.Basket {

			sum += productMap[vBasket.ProductId] * vBasket.Quantity

		}

		fmt.Print(" Total Sum: ", sum, "\n")
	}

}
