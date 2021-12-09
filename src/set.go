package set

type Set[T comparable] struct {
	m map[T]struct{}
}

func NewSet[T comparable](elements ...T) Set[T] {
	result := Set[T]{
		make(map[T]struct{}),
	}
	for _, e := range elements {
		Add(&result, e)
	}
	return result
}

func Add[T comparable](s *Set[T], e T) {
	s.m[e] = struct{}{}
}

func Delete[T comparable](s *Set[T], e T) {
	delete(s.m, e)
}

func Contains[T comparable](s *Set[T], e T) bool {
	_, ok := s.m[e]
	return ok
}

func Union[T comparable](l *Set[T], r *Set[T]) Set[T] {
	result := NewSet[T]()
	for i := range l.m {
		Add(&result, i)
	}
	for i := range r.m {
		Add(&result, i)
	}
	return result
}

func Intersection[T comparable](l *Set[T], r *Set[T]) Set[T] {
	result := NewSet[T]()
	x, y := l, r
	if len(x.m) > len(y.m) {
		x, y = y, x
	}
	for i := range x.m {
		if Contains(y, i) {
			Add(&result, i)
		}
	}
	return result
}

func ToSlice[T comparable](s *Set[T]) []T {
	result := make([]T, 0, len(s.m))
	for i := range s.m {
		result = append(result, i)
	}
	return result
}

func Difference[T comparable](l *Set[T], r *Set[T]) Set[T] {
	result := NewSet[T]()
	for i := range l.m {
		if !Contains(r, i) {
			Add(&result, i)
		}
	}
	return result
}

func ContainsSet[T comparable](l *Set[T], r *Set[T]) bool {
	for i := range l.m {
		if !Contains(r, i) {
			return false
		}
	}
	return true
}

func Size[T comparable](s *Set[T]) int {
	return len(s.m)
}
