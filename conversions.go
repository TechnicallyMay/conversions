package main

import (
	"fmt"
	"strconv"
)

func main() {
    unitList := &UnitNode{}

    for {
        var firstQty, firstUnit, secondQty, secondUnit string
        fmt.Println("Type your command in the format `<firstQuantity> <firstUnit> = <secondQuantity> <secondUnit>`")
        fmt.Println("For example, to add a new conversion: `4 cup = 1 quart`")
        fmt.Println("Or, to get a conversion: `10 cup = ? quart`")
        fmt.Println()
        fmt.Scanf("%s %s = %s %s",  &firstQty, &firstUnit, &secondQty, &secondUnit)

        //TODO: Break down commands
        //TODO: Functionality to print whole list
        if firstQty != "?" && secondQty != "?" {
            firstQtyFloat, _ := strconv.ParseFloat(firstQty, 10)
            secondQtyFloat, _ := strconv.ParseFloat(secondQty, 10)

            unitList.AddConversion(float32(firstQtyFloat), firstUnit, float32(secondQtyFloat), secondUnit)
            fmt.Printf("Added conversion, %v %v = %v %v\n\n", firstQtyFloat, firstUnit, secondQtyFloat, secondUnit)
            continue
        } 

        var qty float64
        var fromName string
        var toName string

        if firstQty == "?" {
            qty, _ = strconv.ParseFloat(secondQty, 10)
            fromName = secondUnit
            toName = firstUnit

        } else {
            qty, _ = strconv.ParseFloat(firstQty, 10)
            fromName = firstUnit
            toName = secondUnit
        }

        unit, err := unitList.GetConversion(float32(qty), fromName, toName)

        if err != nil {
            fmt.Println(err)
            continue
        }

        fmt.Printf("%v %v = %v %v\n\n", qty, fromName, unit, toName)
    }
}
