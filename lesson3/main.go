package main

import "fmt"

type Branch struct {
	Id      int
	Name    string
	Address string
}

type Transaction struct {
	Id        int
	BranchId  int
	ProductId int
	Quantity  int
}

type Product struct {
	Id    int
	Name  string
	Price int
}

var product = []Product{
	{Id: 1, Name: "Olma", Price: 12000},
	{Id: 2, Name: "Banan", Price: 22000},
	{Id: 3, Name: "Olcha", Price: 25000},
}
var transaction = []Transaction{
	{Id: 1, BranchId: 1, ProductId: 1, Quantity: 12},
	{Id: 2, BranchId: 2, ProductId: 2, Quantity: 10},
	{Id: 3, BranchId: 1, ProductId: 3, Quantity: 8},
	{Id: 4, BranchId: 2, ProductId: 1, Quantity: 14},
	{Id: 5, BranchId: 1, ProductId: 2, Quantity: 13},
	{Id: 6, BranchId: 2, ProductId: 3, Quantity: 7},
}
var branches = []Branch{
	{Id: 1, Name: "MGorkiy", Address: "Mirzo Ulug'bek 17 - uy"},
	{Id: 2, Name: "Chilonzor", Address: "Chilonzor Metro"},
}

func main() {

	for _, branchesV := range branches {
		sum := 0
		fmt.Print("Branch Name: ", branchesV.Name)
		for _, transactionV := range transaction {
			if transactionV.BranchId == branchesV.Id {
				for _, productV := range product {
					if productV.Id == transactionV.ProductId {
						sum += productV.Price * transactionV.Quantity
					}
				}
			}
		}
		fmt.Print(" Total Sum: ", sum, "\n")
	}

	counterMap := make(map[int]int)
	productMap := make(map[int]Product)
	branchesMap := make(map[int]Branch)

	for _, v := range product {
		productMap[v.Id] = v
	}

	for _, v := range transaction {
		counterMap[v.BranchId] += productMap[v.ProductId].Price * v.Quantity
	}

	for _, v := range branches {

		branchesMap[v.Id] = v

	}

	for i := range counterMap {
		fmt.Printf("%s -> %d \n", branchesMap[i].Name, counterMap[i])
	}

}
