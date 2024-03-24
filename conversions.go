package main

import (
    "fmt"
    "github.com/TechnicallyMay/conversions/linkedlist"
)

func main() {
    list := linkedlist.LinkedListNode[string]{Value: "hello"}
    list.Add("world")
    list.Add("124")
    list.Add("125")
    list.Add("126")
    list.Add("127")
    list.Add("128")
    list.Add("129")

    list.Remove("128")

    var curr *linkedlist.LinkedListNode[string] = &list

    for curr != nil {
        fmt.Println(curr.Value)
        curr = curr.Next
    }
}
