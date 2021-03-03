package main

import (
	"container/heap"
	"sort"
)

var s []int

/*
* 排序数组
 */

// Insert ..
func Insert(num int) {
	s = append(s, num)
	sort.Ints(s)
}

// GetMedian ..
func GetMedian() float64 {
	n := len(s)
	var res float64
	if n%2 == 0 {
		res = float64(s[n/2]+s[n/2-1]) / 2.0
	} else {
		res = float64(s[n/2])
	}
	return res
}

/*
* 最大堆最小堆
* 最大堆所有元素都比最小堆小
* 如果要插入最大堆的元素比最小堆大，弹出最小堆栈顶，插入最大堆，将新元素插入最小堆
* 反之亦然
* 奇数时插入最大堆，偶数插入最小堆
* n为奇数则返回最大堆数字，否则返回最小堆和最大堆的/2
 */

var n = 1
var maxHeap = NewMaxHeap([]int{})
var minHeap = NewMinHeap([]int{})

// Insert1 ..
func Insert1(num int) {
	if n == 1 {
		n = 0
		heap.Push(minHeap, num)
		heap.Push(maxHeap, heap.Pop(minHeap))
	} else {
		n = 1
		heap.Push(maxHeap, num)
		heap.Push(minHeap, heap.Pop(maxHeap))
	}
}

// GetMedian1 ..
func GetMedian1() float64 {
	var (
		left  int
		right int
	)
	if maxHeap.Len() > 0 {
		left = (*maxHeap)[0]
	}
	if minHeap.Len() > 0 {
		right = (*minHeap)[0]
	}

	if n == 1 {
		val := float64(left + right)
		return val / 2
	}
	return float64(left)
}

// IntMaxHeap max heap
type IntMaxHeap []int

func (h IntMaxHeap) Len() int {
	return len(h)
}

func (h IntMaxHeap) Less(i int, j int) bool {
	return !(h[i] < h[j])
}

func (h IntMaxHeap) Swap(i int, j int) {
	if i < 0 || i >= h.Len() || j < 0 || j >= h.Len() {
		return
	}
	h[i], h[j] = h[j], h[i]
}

// Push ..
func (h *IntMaxHeap) Push(x interface{}) {
	v := x.(int)
	*h = append(*h, v)
}

// Pop ..
func (h *IntMaxHeap) Pop() interface{} {
	if h.Len() == 0 {
		return 0
	}
	old := *h
	v := old[old.Len()-1]
	*h = old[:old.Len()-1]
	return v
}

// NewMaxHeap new max heap
func NewMaxHeap(v []int) *IntMaxHeap {
	mh := &IntMaxHeap{}
	heap.Init(mh)
	for _, k := range v {
		heap.Push(mh, k)
	}

	return mh
}

// IntMinHeap ..
type IntMinHeap []int

func (h IntMinHeap) Len() int {
	return len(h)
}

func (h IntMinHeap) Less(i int, j int) bool {
	return h[i] < h[j]
}

func (h IntMinHeap) Swap(i int, j int) {
	if i < 0 || i >= h.Len() || j < 0 || j >= h.Len() {
		return
	}
	h[i], h[j] = h[j], h[i]
}

// Push ..
func (h *IntMinHeap) Push(x interface{}) {
	v := x.(int)
	*h = append(*h, v)
}

// Pop ..
func (h *IntMinHeap) Pop() interface{} {
	if h.Len() == 0 {
		return 0
	}
	old := *h
	v := old[old.Len()-1]
	*h = old[:old.Len()-1]
	return v

}

// NewMinHeap ..
func NewMinHeap(v []int) *IntMinHeap {
	mh := &IntMinHeap{}
	heap.Init(mh)
	for _, k := range v {
		heap.Push(mh, k)
	}

	return mh
}
