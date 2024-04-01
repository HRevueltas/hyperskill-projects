package main

import (
	"fmt"
)

type Item struct {
	Name   string
	Price  float64
	Earned float64
}

// income calcula el total de ingresos
func income(items []Item) float64 {
	var total float64
	for _, item := range items {
		total += item.Earned
	}
	return total
}

// netIncome calcula el ingreso neto después de deducir gastos
func netIncome(total, staff, other float64) float64 {
	net := total - staff - other
	return net
}

func main() {
	items := []Item{
		{"bubblegum", 2, 202},
		{"toffe", 0.2, 118},
		{"icecream", 5, 2250},
		{"milkchoco", 4, 1680},
		{"doughnut", 2.5, 1075},
		{"pancake", 3.2, 80},
	}

	// Calcular ingreso total
	totalIncome := income(items)

	// Imprimir lista de ingresos
	fmt.Println("Ingresos:")
	fmt.Println("Nombre       | Precio | Ganado")
	fmt.Println("-------------------------------")
	for _, item := range items {
		fmt.Printf("%-12s | $%-6.2f | $%.2f\n", item.Name, item.Price, item.Earned)
	}
	fmt.Println("-------------------------------")
	fmt.Printf("Ingreso total: $%.2f\n\n", totalIncome)

	// Solicitar al usuario los gastos de personal y otros gastos
	var staff, other float64
	fmt.Print("Gastos de personal: $")
	fmt.Scan(&staff)
	fmt.Print("Otros gastos: $")
	fmt.Scan(&other)

	// Calcular ingreso neto
	net := netIncome(totalIncome, staff, other)

	// Imprimir resultado del ingreso neto
	fmt.Println("Ingreso neto: $", net)
}
