package main

import (
    "fmt"
    "github.com/matthewrcheng/mapreduce/collection"
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

	// Map and Reduce
	mapreduced := gs.Map(func(i interface{}) interface{} {
		return i.(int) * 2
	}).Reduce(func(acc interface{}, i interface{}) interface{} {
		return acc.(int) + i.(int)
	}, 0)
	fmt.Println("MapReduced:", mapreduced) // Output: 30
}