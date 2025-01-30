package main 

import (
    "aoc2022/data" 
    "fmt" 
    "strings"
    "strconv"
)

var input []string 


func main() {
    input = data.Get("data/day04.txt")   

    fmt.Println("p1:", part1()) 
    fmt.Println("p2:", part2())
}

func part1() int {
    total := 0 

    for _, p := range input {

        p1, p2 := getNumPair(p)

        if (p1[0] <= p2[0] && p1[1] >= p2[1]) || (p2[0] <= p1[0] && p2[1] >= p1[1]) {
            total++
        }
    }

    return total
}

func part2() int {
    total := 0 

    for _, p := range input {

        p1, p2 := getNumPair(p)

        if (p1[0] <= p2[0] && p1[1] >= p2[0]) || (p2[0] <= p1[0] && p2[1] >= p1[0]) {
            total++
        }
    }

    return total
}

var getNumPair = func(s string) ([]int, []int) {
    pair := strings.Split(s, ",")
    i := make([]int, 2) 
    j := make([]int, 2)
    p1, p2 := strings.Split(pair[0], "-"), strings.Split(pair[1], "-")
    var err error
    i[0], err = strconv.Atoi(p1[0])
    if err != nil {
        panic(err)
    }
    i[1], err = strconv.Atoi(p1[1]) 
    if err != nil {
        panic(err)
    }

    j[0], err = strconv.Atoi(p2[0])
    if err != nil {
        panic(err)
    }
    j[1], err = strconv.Atoi(p2[1]) 
    if err != nil {
        panic(err)
    }
    return i,j
}
