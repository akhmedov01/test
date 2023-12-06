package main

import "fmt"

type Client struct {
	ID     int
	Name   string
	Basket []CardProducts
}

type CardProducts struct {
	ProductId int
	SizeId    int
	Quantity  int
}

type Product struct {
	ID    int
	Name  string
	Sizes []Size
}

type Size struct {
	ID    int
	Name  string
	Price int
}

var clients = []Client{

	{ID: 1, Name: "Akhmedov", Basket: []CardProducts{
		{ProductId: 1, SizeId: 2, Quantity: 4},
		{ProductId: 4, SizeId: 1, Quantity: 3},
		{ProductId: 3, SizeId: 3, Quantity: 1},
	},
	},

	{ID: 2, Name: "Komilov", Basket: []CardProducts{
		{ProductId: 2, SizeId: 1, Quantity: 1},
		{ProductId: 2, SizeId: 2, Quantity: 6},
		{ProductId: 5, SizeId: 1, Quantity: 5},
	},
	},
}

var products = []Product{

	{ID: 1, Name: "Pizza", Sizes: []Size{
		{ID: 1, Name: "25sm", Price: 30000},
		{ID: 2, Name: "30sm", Price: 45000},
		{ID: 3, Name: "35sm", Price: 70000},
	}},

	{ID: 2, Name: "Burger", Sizes: []Size{
		{ID: 1, Name: "Small", Price: 15000},
		{ID: 2, Name: "Medium", Price: 25000},
		{ID: 3, Name: "Big", Price: 30000},
	}},

	{ID: 3, Name: "Hot-Dog", Sizes: []Size{
		{ID: 1, Name: "Small", Price: 20000},
		{ID: 2, Name: "Medium", Price: 25000},
		{ID: 3, Name: "Big", Price: 30000},
	}},

	{ID: 4, Name: "Coca-cola", Sizes: []Size{
		{ID: 1, Name: "0.5", Price: 5000},
		{ID: 2, Name: "1.0", Price: 8000},
		{ID: 3, Name: "1.5", Price: 10000},
	}},

	{ID: 5, Name: "Fanta", Sizes: []Size{
		{ID: 1, Name: "0.5", Price: 5000},
		{ID: 2, Name: "1.0", Price: 8000},
		{ID: 3, Name: "1.5", Price: 11000},
	}},
}

func main() {

	/* max := 1

	for _, clientV := range clients {

		for i := range clientV.Basket {

			if clientV.Basket[i].Quantity > max {

				max = clientV.Basket[i].Quantity

			}

		}

	} */

	for _, clientV := range clients {

		counter := 0
		max := 1

		for i := range clientV.Basket {

			if clientV.Basket[i].Quantity > max {

				max = clientV.Basket[i].Quantity

			}

		}

		for i, cardProductV := range clientV.Basket {

			for _, productV := range products {

				if cardProductV.ProductId == productV.ID {
					for _, sizeV := range productV.Sizes {
						if cardProductV.SizeId == sizeV.ID {

							counter += sizeV.Price * cardProductV.Quantity

						}
					}
				}

				if max == clientV.Basket[i].Quantity && clientV.Basket[i].ProductId == productV.ID {

					fmt.Printf("%s \t %s : %d \n", clientV.Name, productV.Name, max)

				}

			}

		}

		fmt.Printf("%s Count: %d\n", clientV.Name, counter)

	}

}
