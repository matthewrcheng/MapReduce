package mapreduce

import (
	"fmt"
	"testing"
)

func TestGenericSlice(t *testing.T) {
	items := []interface{}{5, 2, 4, 1, 3}

	gs := CreateGenericSlice(items)

	// Test Map
	mapped := gs.Map(func(i interface{}) interface{} {
		return i.(int) * 2
	})
	fmt.Println("Mapped:", mapped.Items()) // Output: [10 4 8 2 6]

	// Test Filter
	filtered := gs.Filter(func(i interface{}) bool {
		return i.(int)%2 == 0
	})
	fmt.Println("Filtered:", filtered.Items()) // Output: [2 4]

	// Test Reduce
	reduced := gs.Reduce(func(acc interface{}, i interface{}) interface{} {
		return acc.(int) + i.(int)
	}, 0)
	fmt.Println("Reduced:", reduced) // Output: 15

	// Test Chaining Map and Reduce
	mapreduced := gs.Map(func(i interface{}) interface{} {
		return i.(int) * 2
	}).Reduce(func(acc interface{}, i interface{}) interface{} {
		return acc.(int) + i.(int)
	}, 0)
	fmt.Println("MapReduced:", mapreduced) // Output: 30

	gs.Sort()
	fmt.Println("Sorted:", gs.items) // Output: [1 2 3 4 5]
}
