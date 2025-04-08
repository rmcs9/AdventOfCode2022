package main 

import (
    "aoc2022/data" 
    "fmt"
	"strings"
	"strconv"
	"slices"
)

var input []string 

var monkeys []*monkey

type monkey struct {
	items []int
	op func(int) int 
	test func(int)
}

func main() {
    input = data.Get("data/day11.txt")

	var div int = 1
	for i := 0; i < len(input); i += 7 {
		div *= parseMonkey(input[i:i+6])
	}
    fmt.Println("p1:", part1()) 
    fmt.Println("p2:", part2(div))
}

func part1() int {
	inspections := make([]int, len(monkeys))
	for range 20 {
		for i, monkey := range monkeys {
			for _, item := range monkey.items {
				item = monkey.op(item)
				item = item / 3
				monkey.test(item)
				inspections[i]++
			}
			monkey.items = make([]int, 0)
		}
	}

	slices.Sort(inspections)
	return inspections[len(inspections) - 1] * inspections[len(inspections) - 2]
}

func part2(div int) int {

	monkeys = make([]*monkey, 0)
	for i := 0; i < len(input); i += 7 {
		parseMonkey(input[i:i+6])
	}
	inspections := make([]int, len(monkeys))
	for range 10000 {
		for i, monkey := range monkeys {
			for _, item := range monkey.items {
				item = monkey.op(item)
				item = item % div
				monkey.test(item)
				inspections[i]++
			}
			monkey.items = make([]int, 0)
		}
	}

	slices.Sort(inspections)
	return inspections[len(inspections) - 1] * inspections[len(inspections) - 2]
}

func parseMonkey(mon []string) int {
	monkeyPTR := new(monkey)
	monkeyPTR.items = make([]int, 0)
	itemsLine := mon[1] 
	items := strings.Split(strings.Split(itemsLine, ":")[1], ",")

	for _, item := range items {
		item = strings.Trim(item, " ")
		num, err := strconv.Atoi(item)
		if err != nil {
			panic(err)
		}
		monkeyPTR.items = append(monkeyPTR.items, num)
	}

	opLine := strings.Split(mon[2], "=")[1]
	monkeyPTR.op = createOp(opLine)
	div, _ := strconv.Atoi(strings.Trim(strings.Split(mon[3], "by")[1], " "))
	m1, _ := strconv.Atoi(strings.Trim(strings.Split(mon[4], "monkey")[1], " "))
	m2, _ := strconv.Atoi(strings.Trim(strings.Split(mon[5], "monkey")[1], " "))
	monkeyPTR.test = createTest(div, m1, m2)
	monkeys = append(monkeys, monkeyPTR)
	return div
}

func createOp(input string) func(int) int {
	opSplit := strings.Split(input, " ") 
	switch opSplit[2] {
		case "+":
		if opSplit[3] == "old" {
			return func(old int) int { return old + old } 
		}
		num, _ := strconv.Atoi(opSplit[3])
		return func(old int) int { return old + num }
		case "*":
		if opSplit[3] == "old" {
			return func(old int) int { return old * old }
		}
		num, _ := strconv.Atoi(opSplit[3]) 
		return func(old int) int { return old * num } 
		default: 
		panic("invalid operator in op function: " + opSplit[1]) 
	}
}

func createTest(div int, m1, m2 int) func(int) {
	return func(val int) {
		if val % div == 0 {
			monkeys[m1].items = append(monkeys[m1].items, val)
		} else {
			monkeys[m2].items = append(monkeys[m2].items, val)
		}
	}
}
