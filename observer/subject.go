package observer

import "container/list"

type Subject struct {
	observers *list.List
}

func (s *Subject) Init() {
	s.observers = list.New()
}

func (s *Subject) registerObserver(observer Observer) {
	s.observers.PushBack(observer)
}

func (s *Subject) unregisterObserver(observer Observer) {
	for temp := s.observers.Front(); temp != nil; temp = temp.Next() {
		if temp.Value == observer {
			s.observers.Remove(temp)
			break
		}
	}
}

func (s *Subject) Notify() {
	for temp := s.observers.Front(); temp != nil; temp = temp.Next() {
		temp.Value.(Observer).Update(s)
	}
}
