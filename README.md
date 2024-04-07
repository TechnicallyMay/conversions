# Overview
Golang application that will allow for arbitrary, user-provided unit conversions to be added and queried. 
This is being built as an opportunity to learn Go while also focusing on writing performant algorithms. 

When a new unit conversion is added to the list, it will be sorted among the other units already known by the 
application, thus building a linked list of units ordered by relative scale. This approach will provide 
* O(n) runtime complexity for lookups
* O(n) runtime complexity for adding new unit conversions
* O(n) memory complexity

See `Alternatives Considered` for some details on why this approach was chosen over others.

For example, if a user were to provide the conversion `16 Tablespoons = 1 cup`, and subsequently provided the 
conversion `16 cups = 1 gallon`, then the application would build the following conversion tree:

Tablespoon -16-> Cup -16-> Gallon

So, if we ask the application how many tablespoons are in a gallon, we get

`Tablespoon * 16 * 16 = Gallon`

or

`256 Tablespoon = 1 gallon`

# Running
Install [Go](https://go.dev/)

`cd conversions`
`go run .`

*To run tests:*
`go test`

# Alternatives Considered
### Using a Graph
TODO

### Using a map
If this were a program actually intended for use, a map may actually be a more practical approach. Though the 
complexity for generating mappings between all units would be high, lookup speed would be O(1). 

Another small drawback of using a map is the memory complexity. Due to storing an entry from each unit to each other 
unit, the memory complexity grows to O(n<sup>2</sup>) compared to the O(1) from the linked list. 

