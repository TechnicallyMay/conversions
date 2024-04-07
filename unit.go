package main

import "fmt"

type UnitNode struct {
    name *string
    Next *UnitNode
    ScaleToNext float32
}

func (n *UnitNode) AddConversion(fromQty float32, fromName string, toQty float32, toName string) {
    logPrefix := fmt.Sprintf("\n[From: %v, To: %v] ", fromName, toName)
    fmt.Printf("\n" + logPrefix + "Entered AddConversion")

    var largerUnitName, smallerUnitName string 
    var largerUnitQty, smallerUnitQty float32

    if fromQty >= toQty {
        fmt.Printf(logPrefix + "From unit is smaller than To unit")
        smallerUnitName, largerUnitName = fromName, toName 
        smallerUnitQty, largerUnitQty = fromQty, toQty 
    } else {
        fmt.Printf(logPrefix + "To unit is smaller than From unit")
        smallerUnitName, largerUnitName = toName, fromName  
        smallerUnitQty, largerUnitQty = toQty, fromQty  
    }

    smallerToLargerScale := smallerUnitQty / largerUnitQty
    if n.name == nil {
        //TODO: Can probably remove this if block and just have the algo naturally handle this.
        fmt.Printf(logPrefix + "DONE: First mapping in the list")
        n.name = &smallerUnitName
        n.ScaleToNext = smallerToLargerScale
        n.Next = &UnitNode{name: &largerUnitName}
        return
    }

    fmt.Printf(logPrefix + "Adding new conversion")

    existingLargerUnit := n
    var smallestToLargerScale float32 = 1
    for *existingLargerUnit.name != largerUnitName {
        if (existingLargerUnit.Next == nil) {
            existingLargerUnit = nil
            break
        }
        smallestToLargerScale *= existingLargerUnit.ScaleToNext
        existingLargerUnit = existingLargerUnit.Next
    }

    if existingLargerUnit != nil {
        fmt.Printf(logPrefix + "Found larger unit in the list")
    } else {
        fmt.Printf(logPrefix + "Did not find larger unit in list")
    }
    
    existingSmallerUnit := n
    existingToLargerScale := smallestToLargerScale
    for *existingSmallerUnit.name != smallerUnitName {
        //TODO: In here somewhere we ned to handle adding unit before beginning
        if existingSmallerUnit.Next == nil {
            existingSmallerUnit = nil
            break
        }

        existingToSmallerScale := existingToLargerScale / smallerToLargerScale
        existingToLargerScale /= existingSmallerUnit.ScaleToNext
        if existingLargerUnit != nil && smallerToLargerScale > existingToLargerScale {
            fmt.Printf(logPrefix + "DONE: Larger unit is already in list, and found where smaller unit belongs")
            smallerUnit := &UnitNode{name: &smallerUnitName, ScaleToNext: existingSmallerUnit.ScaleToNext / existingToSmallerScale, Next: existingSmallerUnit.Next} 
            existingSmallerUnit.ScaleToNext = existingToSmallerScale
            existingSmallerUnit.Next = smallerUnit
            return
        }

        existingSmallerUnit = existingSmallerUnit.Next
    }

    if existingLargerUnit != nil && existingSmallerUnit != nil {
        panic(logPrefix + "Found both units already in list, updates are not supported.")
    }

    next := existingSmallerUnit
    nextToSmallerScale := smallerToLargerScale
    for next.Next != nil && nextToSmallerScale > next.ScaleToNext {
        smallerToLargerScale /= next.ScaleToNext
        next = next.Next
    }

    fmt.Printf(logPrefix + "DONE: Larger unit going after %v", *next.name)
    newUnit := &UnitNode{name: &largerUnitName, ScaleToNext: existingToLargerScale / smallerToLargerScale, Next: next.Next}
    next.Next = newUnit
    next.ScaleToNext = nextToSmallerScale
}

