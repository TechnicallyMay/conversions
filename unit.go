package main

// import "fmt"

type unit struct {
    name string 
}

type UnitNode struct {
    Unit *unit
    Prev *UnitNode
    Next *UnitNode
    ScaleToNext float32
}

func (n *UnitNode) AddConversion(fromQty float32, fromName string, toQty float32, toName string) {
    curr := n
    next := n.Next

    // Ensure from is the smaller of the two, to simplify the later algorithms
    if !(fromQty >= toQty) {
        oldFromQty, oldFromName := fromQty, fromName
        fromQty, fromName = toQty, toName
        toQty, toName = oldFromQty, oldFromName
    }

    scale := fromQty / toQty
    //TODO: Clean this mess up :) 
    if curr.Unit == nil {
        // First element in list
        curr.Unit = &unit{name: fromName}
        curr.Next = &UnitNode{Unit: &unit{name: toName}, Prev: curr}
        curr.ScaleToNext = scale
    } else if curr.Unit.name == fromName {
        if (next == nil) {
            // Current unit is from, and next unit should be to
            curr.Next = &UnitNode{Unit: &unit{name: toName}, Prev: curr}
            curr.ScaleToNext = scale
        } else {    
            if scale > curr.ScaleToNext {
                // Current unit is from, and "to" should be after next
                next.AddConversion(scale / curr.ScaleToNext, next.Unit.name, toQty, toName)
            } else {
                // Current unit is from, and "to" should be between curr and next
                newUnit := &UnitNode{Unit: &unit{name: toName}, Prev: curr, Next: next, ScaleToNext: curr.ScaleToNext / scale}
                next.Prev = newUnit
                curr.Next = newUnit
                curr.ScaleToNext = scale
            }

            // TODO: Calculate appropriate scale 
            // next.AddConversion()
        }
    } else if curr.Unit.name == toName {
    } else if next != nil {
        next.AddConversion(fromQty, fromName, toQty, toName)
    } else {
        // TODO: New list? 
    }
    // Walk the list. 
    // If I find "from" in the list (by value)
    //      Until aggregate scale is > current scale, walk the list
    //          Add toUnit with aggregated scale
    // If I find "to" in the list (by value)
    //      Same steps as above, in reverse
    // If I don't find "from" or "to" in the list? New list? 

    // newNode := &Unit{Value: toAdd}

    // last := n
    // for last.Next != nil {
    //     last = last.Next
    // }

    // last.Next = newNode
}

