package heap

type Heap struct {
	arr []int
}

func New() Heap {
	return Heap{arr: make([]int, 0)}
}

func (h Heap) GetMin() int {
	return h.arr[0]
}

func (h Heap) Size() int {
	return len(h.arr)
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func (h *Heap) InsertKey(k int) {
	h.arr = append(h.arr, k)
	i := h.Size() - 1

	for i != 0 && h.arr[parent(i)] > h.arr[i] {
		h.arr[parent(i)], h.arr[i] = h.arr[i], h.arr[parent(i)]
		i = parent(i)
	}
}

func (h *Heap) DecreaseKey(i, newVal int) {
	h.arr[i] = newVal

	for i != 0 && h.arr[parent(i)] > h.arr[i] {
		h.arr[parent(i)], h.arr[i] = h.arr[i], h.arr[parent(i)]
		i = parent(i)
	}
}

const MaxInt = 100000000

func (h *Heap) ExtractMin() int {
	if h.Size() == 0 {
		return MaxInt
	}

	if h.Size() == 1 {
		root := h.arr[0]
		h.arr = h.arr[:h.Size()-1]

		return root
	}

	root := h.arr[0]
	h.arr[0] = h.arr[h.Size()-1]
	h.arr = h.arr[:h.Size()-1]
	h.MinHeapify(0)

	return root
}

func (h *Heap) DeleteKey(i int) {
	h.DecreaseKey(i, -MaxInt)
	h.ExtractMin()
}

func (h *Heap) MinHeapify(i int) {
	left := left(i)
	right := right(i)
	smallest := i

	if left < h.Size() && h.arr[left] < h.arr[i] {
		smallest = left
	}

	if right < h.Size() && h.arr[right] < h.arr[smallest] {
		smallest = right
	}

	if smallest != i {
		h.arr[i], h.arr[smallest] = h.arr[smallest], h.arr[i]
		h.MinHeapify(smallest)
	}
}
