func lastStoneWeight(stones []int) int {
	h := (*IntHeap)(&stones)
	heap.Init(h)

	for h.Len() > 1 {
		left := heap.Pop(h).(int) - heap.Pop(h).(int)
		if left > 0 {
			heap.Push(h, left)
		}
	}
	if h.Len() == 1 {
		return heap.Pop(h).(int)
	}
	return 0
}

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}