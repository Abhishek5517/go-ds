package heap

type Heap[T any] struct {
	data []T
	less func(a, b T) bool
}

func New[T any](less func(a, b T) bool) *Heap[T] {
	return &Heap[T]{less: less}
}

func (h *Heap[T]) Len() int {
	return len(h.data)
}

func (h *Heap[T]) Empty() bool {
	return len(h.data) == 0
}

func (h *Heap[T]) Top() T {
	if len(h.data) == 0 {
		panic("heap is empty")
	}
	return h.data[0]
}

func (h *Heap[T]) Push(x T) {
	h.data = append(h.data, x)
	h.up(len(h.data) - 1)
}

func (h *Heap[T]) Pop() T {
	if len(h.data) == 0 {
		panic("heap is empty")
	}

	n := len(h.data)
	h.data[0], h.data[n-1] = h.data[n-1], h.data[0]
	x := h.data[n-1]
	h.data = h.data[:n-1]

	if len(h.data) > 0 {
		h.down(0)
	}
	return x
}

func (h *Heap[T]) up(i int) {
	for {
		p := (i - 1) / 2
		if i == 0 || !h.less(h.data[i], h.data[p]) {
			break
		}
		h.data[i], h.data[p] = h.data[p], h.data[i]
		i = p
	}
}

func (h *Heap[T]) down(i int) {
	n := len(h.data)
	for {
		l := 2*i + 1
		if l >= n {
			break
		}
		j := l
		r := l + 1
		if r < n && h.less(h.data[r], h.data[l]) {
			j = r
		}
		if !h.less(h.data[j], h.data[i]) {
			break
		}
		h.data[i], h.data[j] = h.data[j], h.data[i]
		i = j
	}
}
