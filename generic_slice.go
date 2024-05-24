package mapreduce

type GenericSlice struct {
	items []interface{}
}

func CreateGenericSlice(items []interface{}) *GenericSlice {
	return &GenericSlice{items: items}
}

func (s *GenericSlice) Items() interface{} {
	return s.items
}

func Mergesort(arr []interface{}) []interface{} {
	if len(arr) <= 1 {
		return arr
	}

	arr1 := Mergesort(arr[:len(arr)/2])
	arr2 := Mergesort(arr[len(arr)/2:])

	return merge(arr1, arr2)
}

func merge(arr1 []interface{}, arr2 []interface{}) []interface{} {
	var arr []interface{}

	index1 := 0
	index2 := 0

	for index1 < len(arr1) && index2 < len(arr2) {
		if less(arr1[index1], arr2[index2]) {
			arr = append(arr, arr1[index1])
			index1++
		} else {
			arr = append(arr, arr2[index2])
			index2++
		}
	}

	for ;index1 < len(arr1);index1++ {
		arr = append(arr, arr1[index1])
	}

	for ;index2 < len(arr2);index2++ {
		arr = append(arr, arr2[index2])
	}

	return arr
}

func (s *GenericSlice) Sort() {
	s.items = Mergesort(s.items)
}

func (s *GenericSlice) Map(mapper func(interface{}) interface{}) Collection {
	result := make([]interface{}, len(s.items))
	for i, v := range s.items {
		result[i] = mapper(v)
	}
	return &GenericSlice{items: result}
}

func (s *GenericSlice) Filter(filter func(interface{}) bool) Collection {
	var result []interface{}
	for _, v := range s.items {
		if filter(v) {
			result = append(result, v)
		}
	}
	return &GenericSlice{items: result}
}

func (s *GenericSlice) Reduce(reducer func(interface{}, interface{}) interface{}, initial interface{}) interface{} {
	result := initial
	for _, v := range s.items {
		result = reducer(result, v)
	}
	return result
}
