package main 

import ( 
    "fmt"
    "aoc2022/data"
)

var input []string

func main() {
    input = data.Get("data/day06.txt")

    fmt.Println("part1:", part1())
    fmt.Println("part2:", part2())
}


func part1() int {
    signal := input[0]

    count, mem := 0, make(map[rune]bool)
    for i, c := range signal {
        if _, in := mem[c]; in {
            mem = make(map[rune]bool) 
            count = 1
            mem[c] = true
        } else {
            if count >= 4 {
                return i
            }

            count++ 
            mem[c] = true
        }
    }
    return -1
}


func part2() int {
    signal := input[0] 
    
    for i := range signal {
        flag := true
        mem := make(map[byte]bool)
        for j := i; j < i + 14; j++ {
            if _, in := mem[signal[j]]; in {
                flag = false
                break
            }
            mem[signal[j]] = true
        }

        if flag {
            return i + 14
        }
    }
    return -1
}
