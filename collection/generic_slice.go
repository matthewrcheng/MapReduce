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
