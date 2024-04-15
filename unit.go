package main

import (
    "errors"
    "fmt"
    "golang.org/x/exp/slices"
    "math"
)

type UnitNode struct {
    name *string
    Next *UnitNode
    ScaleToNext float64
}

func (n *UnitNode) AddConversion(fromQty float64, fromName string, toQty float64, toName string) (*UnitNode, error) {
    logger := Logger{}
    logPrefix := fmt.Sprintf("\n[From: %v, To: %v] ", fromName, toName)
    logger.Debug("\n" + logPrefix + "Entered AddConversion")

    smallerQty, smallerName, largerQty, largerName := sortUnits(fromQty, fromName, toQty, toName)
    smallerToLargerScale := smallerQty / largerQty
    if n.Next == nil {
        logger.Debug(logPrefix + "DONE: First mapping in the list")
        n.name = &smallerName
        n.ScaleToNext = smallerToLargerScale
        n.Next = &UnitNode{name: &largerName, ScaleToNext: 1}
        return n, nil
    }

    logger.Debug(logPrefix + "Adding new conversion")

    // TODO: don't try to pull both when happy path is only one existing
    smallerUnit, _ := n.findFirstMatchingNodeByName(smallerName)
    largerUnit, startToLargerScale := n.findFirstMatchingNodeByName(largerName)

    if smallerUnit != nil && largerUnit != nil {
        return nil, errors.New("Both units found in list, updates are not supported")
    }

    if smallerUnit != nil {
        logger.Debug(logPrefix + "DONE: Smaller unit found in list, adding larger unit at target scale.")
        return smallerUnit.insertUnitAtTargetScale(largerName, smallerToLargerScale), nil
    }

    if largerUnit != nil {
        if startToLargerScale < smallerToLargerScale {
            logger.Debug(logPrefix + "DONE: Smaller unit belongs at beginning of the list.")
            newUnit := &UnitNode{name: n.name, ScaleToNext: n.ScaleToNext, Next: n.Next}
            n.name = &smallerName
            n.ScaleToNext = smallerToLargerScale / startToLargerScale
            n.Next = newUnit
            return n, nil
        }

        logger.Debug(logPrefix + "DONE: Larger unit found in list, adding smaller unit at target scale.")
        return n.insertUnitAtTargetScale(smallerName, startToLargerScale / smallerToLargerScale), nil
    }

    return nil, errors.New("Neither unit found in list")
}

func (n *UnitNode) GetConversion(fromQty float64, fromName string, toName string) (float64, error) {
    smallerUnit, _ := n.findFirstMatchingNodeByName(fromName, toName)

    if smallerUnit == nil {
        return 0, errors.New("Neither unit was contained in list.")
    }

    var conversion, scaleSmallerToLarger float64
    var largerUnit *UnitNode

    if *smallerUnit.name == fromName {
        largerUnit, scaleSmallerToLarger = smallerUnit.findFirstMatchingNodeByName(toName)
        conversion = fromQty / scaleSmallerToLarger
    } else {
        largerUnit, scaleSmallerToLarger = smallerUnit.findFirstMatchingNodeByName(fromName)
        conversion = fromQty * scaleSmallerToLarger
    }

    if largerUnit == nil {
        return 0, errors.New("Larger unit was not found in the list")
    }

    return conversion, nil
}

func (n *UnitNode) insertUnitAtTargetScale(newUnitName string, targetScale float64) *UnitNode {
    curr := n
    currScale := float64(1)

    for curr.Next != nil && currScale*curr.ScaleToNext < targetScale {
        currScale *= curr.ScaleToNext
        curr = curr.Next
    }

    proportion := targetScale / currScale
    newScaleToNext := math.Max(1, n.ScaleToNext / proportion)
    newUnit := &UnitNode{name: &newUnitName, ScaleToNext: newScaleToNext, Next: n.Next}
    n.ScaleToNext = proportion
    n.Next = newUnit

    return newUnit
}

func (n *UnitNode) findFirstMatchingNodeByName(names ...string) (unit *UnitNode, scaleFromStart float64) {
    curr := n
    currScale := float64(1)

    for curr != nil {
        if slices.Contains(names, *curr.name) {
            return curr, currScale
        }

        currScale *= curr.ScaleToNext
        curr = curr.Next
    }

    return nil, math.MaxFloat64
}

func sortUnits(fromQty float64, fromName string, toQty float64, toName string) (float64, string, float64, string) {
    if fromQty >= toQty {
        return fromQty, fromName, toQty, toName
    } else {
        return toQty, toName, fromQty, fromName
    }
}

