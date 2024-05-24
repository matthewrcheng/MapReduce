package mapreduce

type Collection interface {
	Items() interface{}
	Map(mapper func(interface{}) interface{}) Collection
	Filter(filter func(interface{}) bool) Collection
	Reduce(reducer func(interface{}, interface{}) interface{}, initial interface{}) interface{}
}
