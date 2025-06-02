package main 


import (
	"aoc2022/data" 
	"fmt"
	"container/heap"
	"aoc2022/util"
	"strconv"
)


var input []string 


func main() {
	input = data.Get("data/day12.txt")

	fmt.Println("p1:", part1(nil))
	fmt.Println("p2:", part2())
}

func part1(start *coord) int {
	comp := func(a, b coord) bool {
		return a.score < b.score
	}

	Q := util.MakePrio(comp)
	seen := make(map[string]int)

	if start == nil {
		for y := range input {
			for x, val := range input[y] {
				if val == 'S' {
					start = &coord{0, 'a', x, y}
					break
				}
			}
			if Q.Len() > 0 {
				break 
			}
		}
	}


	heap.Push(Q, *start)

	for Q.Len() > 0 {
		c := heap.Pop(Q).(coord)
		if input[c.y][c.x] == 'E' {
			return c.score
		}


		for _, dir := range dirs {
			nX, nY := c.x + dir.x, c.y + dir.y 

			if nY < 0 || nY >= len(input) || nX < 0 || nX >= len(input[nY]) { continue }

			nVal := func(d byte) byte {
				if d == 'E' {
					return 'z'
				} else if d == 'S' {
					return 'a'
				}
				return d
			}(input[nY][nX])

			if nVal > c.val + 1 { continue }

			if s, has := seen[strconv.Itoa(nX) + " " + strconv.Itoa(nY)]; (!has || (c.score + 1 < s)) {
				nC := coord{c.score + 1, nVal, nX, nY}
				heap.Push(Q, nC)
				seen[strconv.Itoa(nX) + " " + strconv.Itoa(nY)] = nC.score
			}
		}
	}

	return -1
}

// I can think of a lot of cool ways to filter out certain 'a' points from the search but 
// idk this takes less than a quarter of a second
func part2() int {
	path := 2000000
	for y := range input {
		for x, val := range input[y] {
			if val == 'a' || val =='S' {
				start := &coord{0, 'a', x, y}
				score := part1(start)
				if score < path && score > 0 {
					path = score
				}
			}
		}
	}
	return path
}

type state struct {
	x,y int
}

type coord struct {
	score int
	val byte 
	x, y int
}

var dirs = []coord {
	{0, 0, 0, 1},
	{0, 0, 1, 0}, 
	{0, 0, 0, -1}, 
	{0, 0, -1, 0},
}
