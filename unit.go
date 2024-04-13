package main

import (
	"errors"
	"fmt"
	"math"
    "golang.org/x/exp/slices"
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

    smallerQty, smallerName, largerQty, largerName := sortUnits(fromQty, fromName, toQty, toName)
    smallerToLargerScale := smallerQty / largerQty
    if n.Next == nil {
        logger.Debug(logPrefix + "DONE: First mapping in the list")
        n.name = &smallerName
        n.ScaleToNext = smallerToLargerScale
        n.Next = &UnitNode{name: &largerName, ScaleToNext: 1}
        return
    }

    logger.Debug(logPrefix + "Adding new conversion")

    //TODO don't try to pull both when happy path is only one existing
    smallerUnit, _ := n.findFirstMatchingNodeByName(smallerName)
    largerUnit, startToLargerScale := n.findFirstMatchingNodeByName(largerName)

    if smallerUnit != nil && largerUnit != nil {
        panic("Both units found in list, updates are not supported")
    }

    if largerUnit != nil {
        logger.Debug(logPrefix + "Larger unit found in list, searching for smaller unit's position")
        
        if startToLargerScale < smallerToLargerScale {
            logger.Debug(logPrefix + "DONE: Smaller unit belongs at beginning of the list.")
            newUnit := &UnitNode{name: n.name, ScaleToNext: n.ScaleToNext, Next: n.Next}
            n.name = &smallerName
            n.ScaleToNext = smallerToLargerScale / startToLargerScale
            n.Next = newUnit
            return
        }
        
        targetScaleStartToSmaller := startToLargerScale / smallerToLargerScale
        before, scaleStartToBefore := n.findNodeByTargetScale(targetScaleStartToSmaller)
        logger.Debug(logPrefix + "DONE: Smaller unit belongs after %v, which has a scale from start of %v", *before.name, scaleStartToBefore)

        newBeforeScaleToNext := targetScaleStartToSmaller / scaleStartToBefore
        smallerUnit := &UnitNode{name: &smallerName, ScaleToNext: before.ScaleToNext / newBeforeScaleToNext, Next: before.Next} 
        before.ScaleToNext = newBeforeScaleToNext
        before.Next = smallerUnit
        return
    } 

    if smallerUnit != nil {
        logger.Debug(logPrefix + "Smaller unit found in list, searching for larger unit's position")

        before, scaleSmallerToBefore := smallerUnit.findNodeByTargetScale(smallerToLargerScale)
        logger.Debug(logPrefix + "DONE: Larger unit belongs after %v, which has a scale from smaller of %v", *before.name, scaleSmallerToBefore)
        largerUnit := &UnitNode{name: &largerName, ScaleToNext: float32(math.Max(1, float64(before.ScaleToNext / smallerToLargerScale))), Next: before.Next} 
        before.ScaleToNext = smallerToLargerScale / scaleSmallerToBefore
        before.Next = largerUnit
        return 
    }
    
    panic("Neither unit found in list")
}

func (n *UnitNode) GetConversion(fromQty float32, fromName string, toName string) (float32, error) {
    smallerUnit, scaleStartToSmaller := n.findFirstMatchingNodeByName(fromName, toName)
    
    if smallerUnit == nil {
        return 0, errors.New("Neither unit was contained in list.")
    }

    largerUnit, scaleStartToLarger := smallerUnit.findFirstMatchingNodeByName(fromName, toName)

    if largerUnit == nil {
        return 0, errors.New("Both units were not contained in list.")
    }
    
    proportion := scaleStartToLarger / scaleStartToSmaller
    return fromQty * proportion, nil
}

func (n *UnitNode) findNodeByTargetScale(targetScale float32) (*UnitNode, float32) {
    curr := n
    currScale := float32(1)
    
    for curr.Next != nil && currScale * curr.ScaleToNext < targetScale {
        currScale *= curr.ScaleToNext
        curr = curr.Next
    }
    
    return curr, currScale
}

func (n *UnitNode) findFirstMatchingNodeByName(names ...string) (unit *UnitNode, scaleFromStart float32) {
    curr := n
    currScale := float32(1)
    
    for curr != nil {
        if slices.Contains(names, *curr.name) {
            return curr, currScale
        }

        currScale *= curr.ScaleToNext
        curr = curr.Next
    }

    return nil, math.MaxFloat32
}

func sortUnits(fromQty float32, fromName string, toQty float32, toName string) (float32, string, float32, string) {
    if fromQty >= toQty {
        return fromQty, fromName, toQty, toName
    } else {
        return toQty, toName, fromQty, fromName
    }
}

