package mapreduce

import (
	"fmt"
	"testing"
)

func Mapper(i interface{}) interface{} {
	return i.(int) * 2
}

func Filterer(i interface{}) bool {
	return i.(int)%2 == 0
}

func Reducer(i interface{}, j interface{}) interface{} {
	return i.(int) + j.(int)
}

func TestMapReduce(t *testing.T) {
	arr := []interface{}{1, 2, 3, 4, 5}
	fmt.Println(arr) 													// Expected: [1 2 3 4 5]	Actual: [1 2 3 4 5]
	new_arr := Map(Mapper, arr)
	fmt.Println(new_arr)                                                // Expected: [2 4 6 8 10]	Actual: [2 4 6 8 10]
	fmt.Println(Filter(Filterer, arr))                                  // Expected: [2 4]			Actual: [2 4]
	fmt.Println(Reduce(Reducer, arr, 0))                                // Expected: 15				Actual: 15
	fmt.Println(Reduce(Reducer, new_arr, 0))                            // Expected: 30				Actual: 30
	fmt.Println(Reduce(Reducer, Filter(Filterer, new_arr), 0))          // Expected: 30				Actual: 30
	fmt.Println(Reduce(Reducer, Map(Mapper, Filter(Filterer, arr)), 0)) // Expected: 12				Actual: 12
}
