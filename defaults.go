package main

func (n *UnitNode) addDefaultConversions() {
    n.AddConversion(16, "Tablespoon", 1, "cup")
    n.AddConversion(4, "cup", 1, "quart")
    n.AddConversion(4, "quart", 1, "gallon")
}
