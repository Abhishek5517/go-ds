package queue

type Queue[T any] struct {
	data []T
	head int
}

func New[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Push(val T) {
	q.data = append(q.data, val)
}

func (q *Queue[T]) Front() T {
	if q.Empty() {
		panic("queue is empty")
	}
	return q.data[q.head]
}

func (q *Queue[T]) Pop() T {

	if q.Empty() {
		panic("queue is empty")
	}
	x := q.data[q.head]
	var zero T
	q.data[q.head] = zero
	q.head++
	if q.head > 64 && q.head*2 >= len(q.data) {
		q.data = append([]T(nil), q.data[q.head:]...)
		q.head = 0
	}
	return x
}

func (q *Queue[T]) Empty() bool {
	return q.head >= len(q.data)
}

func (q *Queue[T]) Size() int {
	return len(q.data) - q.head
}

func (q *Queue[T]) Clear() {
	q.data = nil
	q.head = 0
}
