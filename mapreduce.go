package mapreduce

func Map(mapper func(interface{}) interface{}, arr []interface{}) []interface{} {
	res := make([]interface{}, len(arr))
	for idx,item := range arr {
		res[idx] = mapper(item)
	}
	return res
}

func Filter(filterer func(interface{}) bool, arr []interface{}) []interface{} {
	var res []interface{}
	for _,item := range arr {
		if filterer(item) {
			res = append(res, item)
		}
	}
	return res
}

func Reduce(reducer func(interface{}, interface{}) interface{}, arr []interface{}, res interface{}) interface{} {
	for _,item := range arr {
		res = reducer(item, res)
	}
	return res
}