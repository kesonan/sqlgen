package set

import "container/list"

type ListSet struct {
	list *list.List
	m    map[interface{}]*list.Element
}

func From(data ...any) *ListSet {
	set := &ListSet{
		list: list.New(),
		m:    make(map[interface{}]*list.Element),
	}
	set.Add(data...)
	return set
}

func FromString(data ...string) *ListSet {
	return From(string2Any(data...)...)
}

func (s *ListSet) Add(data ...any) {
	for _, v := range data {
		if _, ok := s.m[v]; ok {
			continue
		}

		var el = s.list.PushBack(v)
		s.m[v] = el
	}
}

func string2Any(data ...string) []any {
	var ret []any
	for _, v := range data {
		ret = append(ret, v)
	}

	return ret
}
func (s *ListSet) AddStringList(data []string) {
	s.Add(string2Any(data...)...)
}

func (s *ListSet) Remove(data any) {
	el, ok := s.m[data]
	if !ok {
		return
	}

	delete(s.m, data)
	s.list.Remove(el)
}

func (s *ListSet) Exists(data any) bool {
	_, ok := s.m[data]
	return ok
}

func (s *ListSet) String() []string {
	var ret []string
	s.Range(func(v interface{}) {
		s, ok := v.(string)
		if ok {
			ret = append(ret, s)
		}
	})

	return ret
}

func (s *ListSet) Int() []int {
	var ret []int
	s.Range(func(v interface{}) {
		i, ok := v.(int)
		if ok {
			ret = append(ret, i)
		}
	})

	return ret
}

func (s *ListSet) Int32() []int32 {
	var ret []int32
	s.Range(func(v interface{}) {
		i, ok := v.(int32)
		if ok {
			ret = append(ret, i)
		}
	})

	return ret
}

func (s *ListSet) Int64() []int64 {
	var ret []int64
	s.Range(func(v interface{}) {
		i, ok := v.(int64)
		if ok {
			ret = append(ret, i)
		}
	})

	return ret
}

func (s *ListSet) Init() {
	s.list = list.New()
	s.m = make(map[interface{}]*list.Element)
}

func (s *ListSet) Range(fn func(v interface{})) {
	var next = s.list.Front()
	if next == nil {
		return
	}
	for {
		if next == nil {
			return
		}
		fn(next.Value)
		next = next.Next()
	}
}
