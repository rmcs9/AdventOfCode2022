package main 


import (

    "aoc2022/data" 
    "fmt"
    // "strings"
)

var input []string 

func main() {
    input = data.Get("data/day02.txt")


    fmt.Println("p1:", part1())
    fmt.Println("p2:", part2())
}


func part1() int {
    
    total := 0
    
    for _, round := range input {
        total += result[round]
    }

    return total
}

func part2() int {
    total := 0

    for _, round := range input {
        total += result2[round]
    }

    return total
}

var result = map[string]int{  
    "A X" : 4,
    "A Y" : 8,
    "A Z" : 3, 
    "B X" : 1, 
    "B Y" : 5,
    "B Z" : 9,
    "C X" : 7,
    "C Y" : 2,
    "C Z" : 6,
}

var result2 = map[string]int{  
    "A X" : 3,
    "A Y" : 4,
    "A Z" : 8, 
    "B X" : 1, 
    "B Y" : 5,
    "B Z" : 9,
    "C X" : 2,
    "C Y" : 6,
    "C Z" : 7,
}
