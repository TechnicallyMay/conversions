package main

import "fmt"

func main() {
	list := &UnitNode{}

    list.AddConversion(10, "one", 1, "ten")
    list.AddConversion(5, "two", 1, "ten")
    list.AddConversion(8, "two", 2, "eight")
    list.AddConversion(2, "ten", 1, "twenty")

    //TODO: Insert at beginning of list
	printUnitList(list)
}

func printUnitList(firstElementInList *UnitNode) {
	var curr *UnitNode = firstElementInList

    fmt.Print("\n\nStarting final exam:\n\n")
	var i = 0
	for curr != nil {
		fmt.Printf("\n[Final Exam] Unit at position %v is %v. ScaleToNext is %v", i, *curr.name, curr.ScaleToNext)
        curr = curr.Next
		i++
	}
}
