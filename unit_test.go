package main

import "testing"

func TestAddShouldInsertUnitAfter(t *testing.T) {
    list := &unitNode{}

    list.AddConversion(16, "Tablespoon", 1, "cup")
    list.AddConversion(16, "cup", 1, "gallon")

    list.assertEquals(t, "Tablespoon", 16)
    list.Next.assertEquals(t, "cup", 16)
    list.Next.Next.assertEquals(t, "gallon", 1)
}

func TestAddShouldInsertUnitAtBeginning(t *testing.T) {
    list := &unitNode{}

    list.AddConversion(16, "cup", 1, "gallon")
    // Even though tablespoon is added second, it should become the first element in the list.
    list.AddConversion(16, "Tablespoon", 1, "cup")

    list.assertEquals(t, "Tablespoon", 16)
    list.Next.assertEquals(t, "cup", 16)
    list.Next.Next.assertEquals(t, "gallon", 1)
}

func TestAddShouldInsertToUnitBetween(t *testing.T) {
    list := &unitNode{}

    list.AddConversion(16, "Tablespoon", 1, "cup")
    list.AddConversion(16, "cup", 1, "gallon")
    // Even though it's added after gallon, quart should go between cup and gallon
    list.AddConversion(4, "cup", 1, "quart")

    list.assertEquals(t, "Tablespoon", 16)
    list.Next.assertEquals(t, "cup", 4)
    list.Next.Next.assertEquals(t, "quart", 4)
    list.Next.Next.Next.assertEquals(t, "gallon", 1)
}

func TestAddShouldInsertFromUnitBetween(t *testing.T) {
    list := &unitNode{}
    list.AddConversion(48, "teaspoon", 1, "cup")
    list.AddConversion(4, "cup", 1, "quart")
    // Even though it's added last, Tablespoon should go between teaspoon and cup
    list.AddConversion(64, "Tablespoon", 1, "quart")

    list.assertEquals(t, "teaspoon", 3)
    list.Next.assertEquals(t, "Tablespoon", 16)
    list.Next.Next.assertEquals(t, "cup", 4)
    list.Next.Next.Next.assertEquals(t, "quart", 1)
}
//TODO: Equal units ( 1 = 1 )
//TODO: Math (2 x = 3 y)
//TODO: Conflicting conversions?

func (node *unitNode) assertEquals(t *testing.T, expectedName string, scaleToNext float64) {
    if *node.name != expectedName {
        t.Errorf("Expected unit name to be '%v', was '%v'", expectedName, *node.name)
    }

    if node.ScaleToNext != scaleToNext {
        t.Errorf("Expected unit with name '%v' ScaleToNext to be '%v', was '%v'",
            *node.name, scaleToNext, node.ScaleToNext)
    }
}

