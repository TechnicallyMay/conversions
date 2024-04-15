package main

import (
    "fmt"
    "strconv"
)

func main() {
    unitList := &UnitNode{}
    unitList.addDefaultConversions()

    fmt.Println("Type your command in the format `<firstQuantity> <firstUnit> = <secondQuantity> <secondUnit>`")
    fmt.Println("For example, to add a new conversion: `4 cup = 1 quart`")
    fmt.Println("Or, to get a conversion: `10 cup = ? quart`")

    for {
        var firstQty, firstUnit, secondQty, secondUnit string
        fmt.Println()
        fmt.Scanf("%s %s = %s %s",  &firstQty, &firstUnit, &secondQty, &secondUnit)

        if firstQty == "" || firstUnit == "" || secondQty == "" || secondUnit == "" {
            fmt.Println("Invalid input, please try again")
            continue
        }

        //TODO: Break down commands
        //TODO: Functionality to print whole list
        if firstQty != "?" && secondQty != "?" {
            firstQtyFloat, _ := strconv.ParseFloat(firstQty, 10)
            secondQtyFloat, _ := strconv.ParseFloat(secondQty, 10)

            _, error := unitList.AddConversion(firstQtyFloat, firstUnit, secondQtyFloat, secondUnit)
            if error != nil {
                fmt.Printf("There was an error adding the conversion: %v\n", error)
            } else {
                fmt.Printf("Added conversion, %v %v = %v %v\n\n", firstQtyFloat, firstUnit, secondQtyFloat, secondUnit)
            }
            continue
        }

        var qty float64
        var fromName, toName string
        var parseError error

        if firstQty == "?" {
            qty, parseError = strconv.ParseFloat(secondQty, 10)
            fromName = secondUnit
            toName = firstUnit
        } else {
            qty, parseError = strconv.ParseFloat(firstQty, 10)
            fromName = firstUnit
            toName = secondUnit
        }

        if parseError != nil {
            fmt.Printf("There was an error parsing your input: %v", parseError)
            continue
        }

        toQty, err := unitList.GetConversion(qty, fromName, toName)

        if err != nil {
            fmt.Printf("There was an error getting the conversion: %v\n", err)
            continue
        }

        fmt.Printf("%v %v = %v %v\n\n", qty, fromName, toQty, toName)
    }
}
