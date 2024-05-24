# MapReduce Package

A simple Go package providing generic `Map`, `Filter`, and `Reduce` functions for collections. This package is designed to work with collections of any type using the `interface{}` type, making it highly flexible and reusable.

## Features

- **Map**: Applies a function to each element in the collection.
- **Filter**: Selects elements from the collection that satisfy a predicate function.
- **Reduce**: Aggregates elements of the collection into a single value using a reducer function.
- **ChainSort**: Sorts the collection in ascending order.
- **Sort**: Sorts the collection in ascending order in place.

## Installation

To install the package, use the following command:

```sh
go get github.com/matthewrcheng/mapreduce
```

## Usage
Here is an example of how to use the MapReduce package in your Go project:

```go
package main

import (
    "fmt"
    "github.com/matthewrcheng/mapreduce"
)

func main() {
    items := []interface{}{1, 2, 3, 4, 5}
    gs := collection.NewGenericSlice(items)

    // Map
    mapped := gs.Map(func(i interface{}) interface{} {
        return i.(int) * 2
    })
    fmt.Println("Mapped:", mapped.Items()) // Output: [2 4 6 8 10]

    // Filter
    filtered := gs.Filter(func(i interface{}) bool {
        return i.(int)%2 == 0
    })
    fmt.Println("Filtered:", filtered.Items()) // Output: [2 4]

    // Reduce
    reduced := gs.Reduce(func(acc interface{}, i interface{}) interface{} {
        return acc.(int) + i.(int)
    }, 0)
    fmt.Println("Reduced:", reduced) // Output: 15
}
```

## Chaining Operations
One of the powerful features of this package is the ability to chain Map, Filter, and Reduce operations. Here is an example:

```go
package main

import (
    "fmt"
    "github.com/matthewrcheng/mapreduce"
)

func main() {
    items := []interface{}{1, 2, 3, 4, 5}
    gs := collection.NewGenericSlice(items)

    result := gs.
        Map(func(i interface{}) interface{} {
            return i.(int) * 2
        }).
        Filter(func(i interface{}) bool {
            return i.(int) > 4
        }).
        Reduce(func(acc interface{}, i interface{}) interface{} {
            return acc.(int) + i.(int)
        }, 0)

    fmt.Println("Chained result:", result) // Output: 24
}
```

In this example, the Map, Filter, and Reduce operations are chained together. First, the Map function doubles each number, then the Filter function selects numbers greater than 4, and finally, the Reduce function sums the filtered numbers.

## API
### `Collection` Interface

```go
type Collection interface {
	Items() interface{}
	Map(mapper func(interface{}) interface{}) Collection
	Filter(filter func(interface{}) bool) Collection
	Reduce(reducer func(interface{}, interface{}) interface{}, initial interface{}) interface{}
	ChainSort() Collection
	Sort()
}
```

### `GenericSlice` Type
`NewGenericSlice(items []interface{}) *GenericSlice`
Creates a new `GenericSlice` with the given items.

`func (s *GenericSlice) Items() []interface{}`
Returns the items in the `GenericSlice`.

`func (s *GenericSlice) Map(mapper func(interface{}) interface{}) Collection`
Applies the mapper function to each element in the collection and returns a new `GenericSlice` with the results.

`func (s *GenericSlice) Filter(predicate func(interface{}) bool) Collection`
Filters the elements in the collection using the predicate function and returns a new `GenericSlice` with the elements that satisfy the predicate.

`func (s *GenericSlice) Reduce(reducer func(interface{}, interface{}) interface{}, initial interface{}) interface{}`
Aggregates the elements in the collection using the reducer function and an initial value.

`func (s *GenericSlice) ChainSort() Collection`
Sorts the elements of the collection and returns a new `GenericSlice` with the sorted elements.

`func (s *GenericSlice) Sort()`
Sorts the elements of the collection in place.

## Running Tests

To run the tests for this package, navigate to the project directory and use the following command:

```sh
go test .
```

For more detailed output, you can use:

```sh
go test -v .
```