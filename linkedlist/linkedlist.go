package linkedlist

import "fmt"

// Implementation of a LinkedList - elements keep track of the next element in the list
type LinkedListNode[T comparable] struct {
    Value T
    Next *LinkedListNode[T]
}

// Adds an element to the end of the linked list
func (n *LinkedListNode[T]) Add(toAdd T) {
    newNode := &LinkedListNode[T]{Value: toAdd}

    last := n
    for last.Next != nil {
        last = last.Next
    }

    last.Next = newNode
}

// Removes the first matching element from the list
func (n *LinkedListNode[T]) Remove(toRemove T) {
    var prev *LinkedListNode[T] = nil
    var next *LinkedListNode[T] = n.Next

    for next != nil && next.Value != toRemove {
        prev = next
        next = prev.Next
    }

    if next == nil {
        fmt.Printf("Tried to remove a value that wasn't in the list: %v\n", toRemove)
        panic(3)
    }

    prev.Next = next.Next
}

