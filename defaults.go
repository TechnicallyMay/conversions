package main

func getDefaultConversions() []*unitNode {
    volumetric := NewList(16, "tablespoon", 1, "cup")
    volumetric.AddConversion(4, "cup", 1, "quart")
    volumetric.AddConversion(4, "quart", 1, "gallon")
    volumetric.AddConversion(3, "teaspoon", 1, "tablespoon")

    length := NewList(12, "inch", 1, "foot")
    length.AddConversion(5280, "foot", 1, "mile")
    length.AddConversion(1.609, "kilometer", 1, "mile")
    length.AddConversion(9_460_730_472_580.8, "kilometer", 1, "lightyear")
    length.AddConversion(1000, "meter", 1, "kilometer")
    length.AddConversion(149_597_870_700, "meter", 1, "au")

    weight := NewList(16, "ounce", 1, "pound")
    weight.AddConversion(28.35, "gram", 1, "ounce")
    weight.AddConversion(2000, "pound", 1, "ton")
    weight.AddConversion(1000, "gram", 1, "kilogram")

    return []*unitNode{volumetric, length, weight}
}

