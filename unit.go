package main

import (
	"errors"
	"fmt"
)

type UnitNode struct {
    name *string
    Next *UnitNode
    ScaleToNext float32
}

func (n *UnitNode) AddConversion(fromQty float32, fromName string, toQty float32, toName string) {
    logger := Logger{}
    logPrefix := fmt.Sprintf("\n[From: %v, To: %v] ", fromName, toName)
    logger.Debug("\n" + logPrefix + "Entered AddConversion")

    var largerUnitName, smallerUnitName string 
    var largerUnitQty, smallerUnitQty float32

    if fromQty >= toQty {
        logger.Debug(logPrefix + "From unit is smaller than To unit")
        smallerUnitName, largerUnitName = fromName, toName 
        smallerUnitQty, largerUnitQty = fromQty, toQty 
    } else {
        logger.Debug(logPrefix + "To unit is smaller than From unit")
        smallerUnitName, largerUnitName = toName, fromName  
        smallerUnitQty, largerUnitQty = toQty, fromQty  
    }

    smallerToLargerScale := smallerUnitQty / largerUnitQty
    if n.name == nil {
        //TODO: Can probably remove this if block and just have the algo naturally handle this.
        logger.Debug(logPrefix + "DONE: First mapping in the list")
        n.name = &smallerUnitName
        n.ScaleToNext = smallerToLargerScale
        n.Next = &UnitNode{name: &largerUnitName}
        return
    }

    logger.Debug(logPrefix + "Adding new conversion")

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
        logger.Debug(logPrefix + "Found larger unit in the list")

        if smallestToLargerScale < smallerToLargerScale {
            // smallest belongs at beginning. Copy first node to second, then replace details of first with smaller
            newUnit := &UnitNode{name: n.name, Next: n.Next, ScaleToNext: n.ScaleToNext}
            n.ScaleToNext = smallerToLargerScale / smallestToLargerScale;
            n.name = &smallerUnitName
            n.Next = newUnit
            return
        }
    } else {
        logger.Debug(logPrefix + "Did not find larger unit in list")
    }
    
    existingSmallerUnit := n
    existingToLargerScale := smallestToLargerScale
    for *existingSmallerUnit.name != smallerUnitName {
        if existingSmallerUnit.Next == nil {
            existingSmallerUnit = nil
            break
        }

        existingToSmallerScale := existingToLargerScale / smallerToLargerScale
        existingToLargerScale /= existingSmallerUnit.ScaleToNext
        if existingLargerUnit != nil && smallerToLargerScale > existingToLargerScale {
            logger.Debug(logPrefix + "DONE: Larger unit is already in list, and found where smaller unit belongs")
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

    logger.Debug(logPrefix + "DONE: Larger unit going after %v", *next.name)
    newUnit := &UnitNode{name: &largerUnitName, ScaleToNext: existingToLargerScale / smallerToLargerScale, Next: next.Next}
    next.Next = newUnit
    next.ScaleToNext = nextToSmallerScale
}

func (n *UnitNode) GetConversion(fromQty float32, fromName string, toName string) (float32, error) {
    curr := n

    for *curr.name != fromName && *curr.name != toName {
        if curr.Next == nil {
            return 0, errors.New("Neither unit was contained in the list")
        }

        curr = curr.Next
    }

    foundUnit := *curr.name
    var toFind string
    multiply := true
    if *curr.name == fromName {
        toFind = toName
        multiply = false
    } else {
        toFind = fromName
    }

    //TODO: Conversion when going from smaller -> larger unit
    // i.e. 1 teaspoon = ? gallon
    var conversionRate float32 = curr.ScaleToNext
    for *curr.Next.name != toFind {
        if curr.Next == nil {
            return 0, errors.New(fmt.Sprintf("Only one unit, %v,  was contained in the list", foundUnit))
        }

        curr = curr.Next
        conversionRate *= curr.ScaleToNext
    }

    fmt.Printf("\n\nconversionRate: %v, fromQty: %v\n\n", conversionRate, fromQty)

    if multiply {
        return fromQty * conversionRate, nil
    } else {
        return fromQty / conversionRate , nil
    }
}
