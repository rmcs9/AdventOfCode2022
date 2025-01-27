package main 

import (
    "aoc2022/data"
    "fmt"
    "strconv"
    "container/heap"
)
var input []string

func main() {
    input = data.Get("data/day01.txt")

    fmt.Println("p1:", part1())
    fmt.Println("p2:", part2())
}


func part1() int {
    
    leader := 0 
    current := 0 

    for _, str := range input {
        if str != "" {
            num, err := strconv.Atoi(str)
            if err != nil {
                panic(err)
            }

            current += num
        } else {
            if leader < current {
                leader = current
            }
            current = 0
        }
    }

    return leader
}

func part2() int {
    var elfs = &IntHeap{} 

    heap.Init(elfs)

    current := 0
    for _, str := range input {
        if str != "" {
            num, err := strconv.Atoi(str)
            if err != nil {
                panic(err)
            }
            current += num
        } else {
            heap.Push(elfs, current)
            current = 0
        }
    }

    return heap.Pop(elfs).(int) + heap.Pop(elfs).(int) + heap.Pop(elfs).(int)
    
}


type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
