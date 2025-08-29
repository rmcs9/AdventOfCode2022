package main

import (
	"fmt"
	"aoc2022/data"
	"strings"
	"strconv"
)

var input []string 
var stoneMesh map[coord]bool = make(map[coord]bool)
var lowestY int = -1

type coord struct {
	x, y int
}

func main() {
	input = data.Get("data/day14.txt")

	fillMesh()

	fmt.Println("p1:", part1())
	fmt.Println("p2:", part2())
}

func part1() int {
	s := 0

	for true {
		if !dropSand(false) {
			s++
		} else {
			break
		}
	}
	return s
}

func dropSand(p2 bool) (done bool) {
	sand := coord{500, 0}
	down := coord{0, 1} 
	left := coord{-1, 1} 
	right := coord{1, 1}
	check := func(vec coord) bool {
		if p2 {
			return !stoneMesh[coord{sand.x + vec.x, sand.y + vec.y}] && sand.y + vec.y < lowestY + 2
		} else {
			return !stoneMesh[coord{sand.x + vec.x, sand.y + vec.y}]
		}
	}
	for true {
		if check(down) {
			sand.x += down.x 
			sand.y += down.y
			if !p2 {
				if sand.y > lowestY + 1 {
					return true
				}
			}
			continue
		} else if check(left) {
			sand.x += left.x 
			sand.y += left.y 
			continue
		} else if check(right) {
			sand.x += right.x 
			sand.y += right.y
			continue
		} else {
			stoneMesh[sand] = true
			return false
		}
	}
	return true
}

func part2() int {
	stoneMesh = make(map[coord]bool)
	fillMesh()
	s := 0
	for !stoneMesh[coord{500, 0}] {
		dropSand(true)
		s++
	}

	return s
}

func fillMesh() {
	for _, line := range input {
		coords := strings.Split(line, " -> ")

		for i := range len(coords) -1 {
			start, end := func() (coord, coord) {
				startStrings, endStrings := strings.Split(coords[i], ","), strings.Split(coords[i + 1], ",") 

				startX, err := strconv.Atoi(startStrings[0])
				if err != nil {
					panic(err)
				}
				startY, err := strconv.Atoi(startStrings[1])
				if err != nil {
					panic(err)
				}
				endX, err := strconv.Atoi(endStrings[0])
				if err != nil {
					panic(err)
				}
				endY, err := strconv.Atoi(endStrings[1])
				if err != nil {
					panic(err)
				}

				return coord{startX, startY}, coord{endX, endY}
			}()

			if start.x == end.x {
				if start.y < end.y {
					for start.y <= end.y {
						stoneMesh[start] = true
						start.y++
					}
				} else {
					for start.y >= end.y {
						stoneMesh[start] = true 
						start.y--
					}
				}
			} else if start.y == end.y {
				if start.x < end.x {
					for start.x <= end.x {
						stoneMesh[start] = true 
						start.x++
					}
				} else {
					for start.x >= end.x {
						stoneMesh[start] = true 
						start.x--
					}
				}
			} else {
				panic("mismatch")
			}
			if end.y > lowestY {
				lowestY = end.y 
			}
		}
	}
}
