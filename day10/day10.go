package main 

import (
    "fmt"
    "aoc2022/data"
    "strings"
    "strconv"
)

var input []string

func main() {
    input = data.Get("data/day10.txt")


    fmt.Println("p1:", part1())
    fmt.Println("p2:")
    part2()
}

func part1() int {
    regX := 1
    cycleCount := 0
    targetCyc := 20
    instructions := input

    result := 0
    for targetCyc <= 220 {
        instruction := instructions[0]
        instructions = instructions[1:]
        var num int

        if instruction[0] == 'a' {
            splt := strings.Split(instruction, " ")
            num, _ = strconv.Atoi(splt[1])

            cycleCount += 2
        } else {
            cycleCount++
            num = 0
        }

        if cycleCount >= targetCyc {
            result += targetCyc * regX 
            targetCyc += 40
        }
        regX += num
    }
    return result
}

func part2() {
    CRT := make([][]byte, 6) 
    regX := 1
    counter = 0
    inst = "noop"

    for y := range len(CRT) {
        CRT[y] = make([]byte, 40)

        for x := range len(CRT[y]) {
            regX = update(regX) 
            CRT[y][x] = '.'
            for i := -1; i <= 1; i++ {
                if regX + i == x { CRT[y][x] = '#' }
            }
        }
    }

    for _, line := range CRT {
        fmt.Println(string(line))
    }
}

var counter int 
var inst string
func update(regX int) int {
    //if the counter is reset, fetch the next instruction
    if counter <= 0 {
        if inst[0] == 'a' {
            splt := strings.Split(inst, " ") 
            num, _ := strconv.Atoi(splt[1])
            regX += num
        }

        inst = input[0]
        input = input[1:]

        if inst[0] == 'a' {
            counter = 1
        } else {
            counter = 0
        }
    } else {
        counter--
    }
    return regX
}
