package modle

import "reflect"

type Stack struct {
	data []interface{}
}

func (s *Stack) Push(e interface{}) {
	s.data = append(s.data, e)
}

func (s *Stack) Pop() (interface{}, bool) {
	if ret := s.Top(); ret != nil {
		s.data = s.data[:len(s.data)-1]
		return ret, true
	}
	return nil, false
}

func (s *Stack) Top() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return s.data[len(s.data) -1]
}

func (s *Stack) Length() int{
	if s.data == nil {
		return 0
	}
	return len(s.data)
}

func (s *Stack) IsEmpty() bool {
	return s.Length() == 0
}

func (s *Stack) Exist(element interface{}) bool {
	for _, v := range s.data {
		isEqual := reflect.DeepEqual(element, v)
		if isEqual {
			return true
		}
	}
	return false
}