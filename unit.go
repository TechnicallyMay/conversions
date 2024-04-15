package main

import (
    "errors"
    "fmt"
    "golang.org/x/exp/slices"
    "math"
)

type unitNode struct {
    name *string
    Next *unitNode
    ScaleToNext float64
}

func NewList(fromQty float64, fromName string, toQty float64, toName string) *unitNode {
    smallerQty, smallerName, largerQty, largerName := sortUnits(fromQty, fromName, toQty, toName)
    smallerToLargerScale := smallerQty / largerQty

    return &unitNode{name: &smallerName, ScaleToNext: smallerToLargerScale, Next: &unitNode{name: &largerName, ScaleToNext: 1}}
}

func (n *unitNode) AddConversion(fromQty float64, fromName string, toQty float64, toName string) (bool, error) {
    logger := Logger{}
    logPrefix := fmt.Sprintf("\n[From: %v, To: %v] ", fromName, toName)
    logger.Debug("\n" + logPrefix + "Entered AddConversion")

    smallerQty, smallerName, largerQty, largerName := sortUnits(fromQty, fromName, toQty, toName)
    smallerToLargerScale := smallerQty / largerQty

    logger.Debug(logPrefix + "Adding new conversion")

    // TODO: don't try to pull both when happy path is only one existing
    smallerUnit, _ := n.findFirstMatchingNodeByName(smallerName)
    largerUnit, startToLargerScale := n.findFirstMatchingNodeByName(largerName)

    if smallerUnit != nil && largerUnit != nil {
        return false, errors.New("Both units found in list, updates are not supported")
    }

    if smallerUnit != nil {
        logger.Debug(logPrefix + "DONE: Smaller unit found in list, adding larger unit at target scale.")
        return smallerUnit.insertUnitAtTargetScale(largerName, smallerToLargerScale), nil
    }

    if largerUnit != nil {
        if startToLargerScale < smallerToLargerScale {
            logger.Debug(logPrefix + "DONE: Smaller unit belongs at beginning of the list.")
            newUnit := &unitNode{name: n.name, ScaleToNext: n.ScaleToNext, Next: n.Next}
            n.name = &smallerName
            n.ScaleToNext = smallerToLargerScale / startToLargerScale
            n.Next = newUnit
            return true, nil
        }

        logger.Debug(logPrefix + "DONE: Larger unit found in list, adding smaller unit at target scale.")
        return n.insertUnitAtTargetScale(smallerName, startToLargerScale / smallerToLargerScale), nil
    }

    return false, nil
}

func (n *unitNode) GetConversion(fromQty float64, fromName string, toName string) (float64, error) {
    smallerUnit, _ := n.findFirstMatchingNodeByName(fromName, toName)

    if smallerUnit == nil {
        return math.Inf(-1), nil
    }

    var conversion, scaleSmallerToLarger float64
    var largerUnit *unitNode

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

func (n *unitNode) insertUnitAtTargetScale(newUnitName string, targetScale float64) bool {
    curr := n
    currScale := float64(1)

    for curr.Next != nil && currScale*curr.ScaleToNext < targetScale {
        currScale *= curr.ScaleToNext
        curr = curr.Next
    }

    proportion := targetScale / currScale
    newScaleToNext := math.Max(1, n.ScaleToNext / proportion)
    newUnit := &unitNode{name: &newUnitName, ScaleToNext: newScaleToNext, Next: n.Next}
    n.ScaleToNext = proportion
    n.Next = newUnit

    return true
}

func (n *unitNode) findFirstMatchingNodeByName(names ...string) (unit *unitNode, scaleFromStart float64) {
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

