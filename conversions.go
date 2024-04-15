package main

import (
    "errors"
    "fmt"
    "math"
    "strconv"
)

func main() {
    nodes := getDefaultConversions()

    for {
        fmt.Println("Type your command in the format `<firstQuantity> <firstUnit> = <secondQuantity> <secondUnit>`")
        fmt.Println("For example, to add a new conversion: `4 cup = 1 quart`")
        fmt.Println("Or, to get a conversion: `10 cup = ? quart`")

        var firstQty, firstUnit, secondQty, secondUnit string
        fmt.Println()
        fmt.Scanf("%s %s = %s %s",  &firstQty, &firstUnit, &secondQty, &secondUnit)

        if firstQty == "print" {
            runPrintNodes(nodes)
        } else if firstQty != "?" && secondQty != "?" {
            runGetConversion(nodes, firstQty, secondQty, firstUnit, secondUnit)
        } else if firstQty == "" || firstUnit == "" || secondQty == "" || secondUnit == "" {
            fmt.Println("Invalid input, please try again")
        } else {
            runAddConversion(nodes, firstQty, secondQty, firstUnit, secondUnit)
        }

        fmt.Println()
        fmt.Println()
    }
}

func runPrintNodes(nodes []*unitNode) {
    for i, node := range nodes {
        var curr *unitNode = node

        fmt.Printf("\n\nPrinting list %v\n", i)
        var j = 0
        for curr != nil {
            fmt.Printf("\n[List %v] Unit at position %v is %v. ScaleToNext is %v", i, j, *curr.name, curr.ScaleToNext)
            curr = curr.Next
            j++
        }
    }
}

func runGetConversion(nodes []*unitNode, firstQty string, secondQty string, firstUnit string, secondUnit string) {
    firstQtyFloat, _ := strconv.ParseFloat(firstQty, 10)
    secondQtyFloat, _ := strconv.ParseFloat(secondQty, 10)

    var error error
    nodes, error = addConversion(nodes, firstQtyFloat, firstUnit, secondQtyFloat, secondUnit)
    if error != nil {
        fmt.Printf("There was an error adding the conversion: %v\n", error)
    } else {
        fmt.Printf("Added conversion, %v %v = %v %v\n\n", firstQtyFloat, firstUnit, secondQtyFloat, secondUnit)
    }
}

func runAddConversion(nodes []*unitNode, firstQty string, secondQty string, firstUnit string, secondUnit string) {
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
        return
    }

    toQty, err := getConversion(nodes, qty, fromName, toName)

    if err != nil {
        fmt.Printf("There was an error getting the conversion: %v\n", err)
        return
    }

    fmt.Printf("%v %v = %v %v\n\n", qty, fromName, toQty, toName)
}

func getConversion(nodes []*unitNode, qty float64, fromName string, toName string) (float64, error) {
    for _, node := range nodes {
        toQty, err := node.GetConversion(qty, fromName, toName)
        if err != nil {
            return math.Inf(-1), err
        }
        if toQty > 0 {
            return toQty, nil
        }
    }

    return math.Inf(-1), errors.New("Didn't find either unit to convert")
}

func addConversion(nodes []*unitNode, fromQty float64, fromName string, toQty float64, toName string) ([]*unitNode, error) {
    for _, node := range nodes {
        added, err := node.AddConversion(fromQty, fromName, toQty, toName)
        if err != nil {
            return nodes, err
        }
        if added {
            return nodes, nil
        }
    }

    return append(nodes, NewList(fromQty, fromName, toQty, toName)), nil
}

