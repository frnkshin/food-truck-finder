package main

import (
	"fmt"
	"food-truck-finder/pkg/foodtruck"
	tm "github.com/buger/goterm"
	"github.com/eiannone/keyboard"
	"os"
)

// Handles show all page display of all food trucks
func handleShowAllView(trucks *foodtruck.FoodTrucks) {
	for _, truck := range *trucks {
		fmt.Println(truck.String())
	}
}

// Handles page display of food trucks in paginated view
func handlePaginatedView(trucks *foodtruck.FoodTrucks) {
	tm.Clear()
	for _, truck := range *trucks {
		fmt.Println(truck.String())
		tm.Flush()
	}
	fmt.Println("Press any key to show next results\nPress Q to quit")
	char, _, err := keyboard.GetSingleKey()
	if (err != nil) {
		panic(err)
	}
	switch char {
	case 'q', 'Q':
		fmt.Println("Pressed Q -- Quitting program")
		os.Exit(0)
		break
	default:
		tm.Clear()
		return
	}
}

