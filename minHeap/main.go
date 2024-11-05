package main

import (
	"container/heap"
	"fmt"
)

type MinHeap []int

func (h MinHeap) Len() int { return len(h) }

func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }

func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

/*
Push() 實際邏輯:
1. 將新元素插入到完全樹的葉節點 (末尾)
2. 進行上浮, 調整樹的結構

這裡只需要實作 (1)
"container/heap" 已經定義好 (2)
*/

func (h *MinHeap) Push(x interface{}) { // O(log n)
	*h = append(*h, x.(int))
}

/*
Pop() 實際邏輯:
 1. 把根結點的最小值 <-> 葉節點的最大值 Swap
 2. 把現在在葉節點的最小值 Pop
 3. 把在根節點的最大值下沉, 將當前完全樹的最小值上浮的根節點

這邊只需要實作 (2)
"container/heap" 已經定義好 (1), (3)
*/

func (h *MinHeap) Pop() interface{} { // O(log n)
	old := *h
	n := len(old)
	min := old[n-1]
	*h = old[:n-1]

	return min
}

func main() {
	h := &MinHeap{3, 1, 4, 1, 5, 9, 2}
	heap.Init(h)
	heap.Push(h, 3)
	heap.Push(h, 4)

	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
}
