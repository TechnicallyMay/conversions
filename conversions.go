package main

import "fmt"

func main() {
	list := &UnitNode{}
	list.AddConversion(16, "cup", 1, "gallon")
	list.AddConversion(4, "cup", 1, "quart")
	list.AddConversion(2, "cup", 1, "pint")

    //TODO: Insert with one unit not being a 1
	printUnitList(list)
}

func printUnitList(firstElementInList *UnitNode) {
	var curr *UnitNode = firstElementInList

	var i = 0
	for curr != nil {
		fmt.Printf("Unit at position %v is %v. ScaleToNext is %v\n", i, curr.Unit.name, curr.ScaleToNext)
        curr = curr.Next
		i++
	}
}
