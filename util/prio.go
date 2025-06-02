package util 

import ( 
	"container/heap"
)

type Prio[K any] struct {
	comp func(K, K) bool
	Heap []*K
}

func MakePrio[K any] (c func(K, K) bool) *Prio[K] {
	p := new(Prio[K])
	heap.Init(p)
	p.comp = c
	return p
}

func (pq Prio[K]) Len() int { return len(pq.Heap) }

func (pq Prio[K]) Less(i, j int) bool {
	return pq.comp(*pq.Heap[i], *pq.Heap[j])
}

func (pq Prio[K]) Swap(i, j int) {
    pq.Heap[i], pq.Heap[j] = pq.Heap[j], pq.Heap[i]
}

func (pq *Prio[K]) Push(x any) {
    item := x.(K)
    pq.Heap = append(pq.Heap, &item)
}

func (pq *Prio[K]) Pop() any {
    old := *pq 
    n := len(old.Heap) 
    x := old.Heap[n-1]
    old.Heap[n-1] = nil
    (*pq).Heap = old.Heap[0 : n-1]
    return *x
}
