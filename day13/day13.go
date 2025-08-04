package main

import (
	"aoc2022/data"
	"fmt"
	"strconv"
	"sort"
	"slices"
)

var input []string
var lefts, rights []*list

type list struct {
	items []*list
	num int 
	isList bool
}

func main() {
	input = data.Get("data/day13.txt")

	lefts, rights = make([]*list, 0), make([]*list, 0)

	parseLists()
	fmt.Println("p1:", part1())
	fmt.Println("p2:", part2()) 
}


func part1() int {
	total := 0 

	for i := range lefts {
		good, _ := compare(lefts[i], rights[i]) 
		if good {
			total += i + 1
		}
	}
	return total 
}

func part2() int {
	allLists := append(lefts, rights...)
	extra1, _ := parseList("[[2]]")
	extra2, _ := parseList("[[6]]")


	allLists = append(allLists, extra1, extra2)

	sort.Slice(allLists, func(a, b int) bool {
		good, _ := compare(allLists[a], allLists[b])
		return good
	})

	i1 := slices.IndexFunc(allLists, func(l *list) bool {
		return l == extra1
	})	

	i2 := slices.IndexFunc(allLists, func(l *list) bool {
		return l == extra2
	})	

	return (i1 + 1) * (i2 + 1)
}

func parseLists() {
	for i := 0; i < len(input); i += 3 {
		left := input[i]
		right := input[i + 1] 

		leftItem, _ := parseList(left)
		rightItem, _ := parseList(right) 
		
		lefts = append(lefts, leftItem) 
		rights = append(rights, rightItem)
	}
}

func parseList(s string) (*list, string) {
	s = s[1:] 

	this := new(list) 
	this.isList = true 

	for s[0] != ']' {
		if s[0] == '[' {
			subList, newS := parseList(s)
			this.items = append(this.items, subList) 
			s = newS
		} else {
			subNum, newS := parseNumber(s)
			this.items = append(this.items, subNum) 
			s = newS
		}
	}
	//remove closer 
	s = s[1:]
	if len(s) > 0 && s[0] == ',' { s = s[1:] }

	return this, s
}

func parseNumber(s string) (*list, string) {
	numString := ""

	for s[0] != ',' && s[0] != ']' {
		numString += string(s[0]) 
		s = s[1:]
	}
	//remove comma
	if s[0] == ',' { s = s[1:] }

	list := new(list) 
	
	num, err := strconv.Atoi(numString)
	if err != nil {
		panic(err)
	}

	list.num = num
	return list, s
}

func compare(left, right *list) (badCompare, keepComparing bool) {
	if left.isList && right.isList {
		return compareLists(left, right)
	} else if left.isList || right.isList {
		if !left.isList {
			newLeft := new(list) 
			newLeft.isList = true
			newLeft.items = append(newLeft.items, left)
			left = newLeft
		} else {
			newRight := new(list) 
			newRight.isList = true 
			newRight.items = append(newRight.items, right)
			right = newRight
		}
		return compareLists(left, right)

	} else {
		return left.num < right.num, left.num == right.num
	}
}

func compareLists(left, right *list) (badCompare, keepComparing bool) {
	for i := range min(len(left.items), len(right.items)) {
		bc, kc := compare(left.items[i], right.items[i]) 
		if !kc {
			return bc, kc
		}
	}

	return len(left.items) < len(right.items), len(left.items) == len(right.items)
}
