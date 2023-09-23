package system

import (
	"sync"
)

type ConditionCallback[T any] func(item *T) bool

type SafeList[T any] struct {
	mu    sync.Mutex
	array []T
}

func (s *SafeList[T]) Append(item T) {
	s.mu.Lock()
	s.array = append(s.array, item)
	s.mu.Unlock()
}

func (s *SafeList[T]) Remove(index int) {
	s.mu.Lock()
	s.array[index] = s.array[len(s.array)-1]
	s.array = s.array[:len(s.array)-1]
	s.mu.Unlock()
}

func (s *SafeList[T]) Override(array []T) {
	s.mu.Lock()
	s.array = array
	s.mu.Unlock()
}

func (s *SafeList[T]) Array() *[]T {
	return &s.array
}

func (s *SafeList[T]) Exists(condition ConditionCallback[T]) bool {
	for _, v := range s.array {
		if condition(&v) {
			return true
		}
	}
	return false
}

func (s *SafeList[T]) Find(condition ConditionCallback[T]) *T {
	for _, v := range s.array {
		if condition(&v) {
			return &v
		}
	}
	return nil
}

func NewSafeList[T any]() *SafeList[T] {
	c := SafeList[T]{
		array: []T{},
	}
	return &c
}
