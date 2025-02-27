package main 

import ( 
    "aoc2022/data" 
    "fmt" 
    "slices"
    "strings"
    "strconv"
)

var input []string
var stacks [][]byte

func main() {
    input = data.Get("data/day05.txt")
    stacksraw := make([]string, 0)

    for i, line := range input {
        if line != "" {
            stacksraw = append(stacksraw, line)
        } else {
            input = input[i + 1:]
            break
        }
    } 

    stacksraw = stacksraw[:len(stacksraw) - 1]
    slices.Reverse(stacksraw)

    makeStacks := func() {
        stacks = make([][]byte, 0)
        for i := 1; i < len(stacksraw[0]); i += 4 {
            stack := make([]byte, 0)
            for _, stacklvl := range stacksraw {
                if stacklvl[i] != ' '{
                    stack = append(stack, stacklvl[i])
                }
            }
            stacks = append(stacks, stack)
        }   
    }

    makeStacks()
    fmt.Println("p1:", part1())
    makeStacks()
    fmt.Println("p2:", part2())
}

func part1() string {
    moveStacks(true)
    str := "" 
    for _, c := range stacks {
        str += string(c[len(c) - 1])
    }
    return str
}

func part2() string {
    moveStacks(false)
    str := "" 
    for _, c := range stacks {
        str += string(c[len(c) - 1])
    }
    return str

}

func moveStacks(p1 bool) {
    for _, i := range input {
        inst := getInstruction(i)

        source := stacks[inst[1] - 1] 
        amount := inst[0]

        forklift := source[len(source) - amount:]
        stacks[inst[1] - 1] = source[:len(source) - amount]
        if p1 { slices.Reverse(forklift) }

        stacks[inst[2] - 1] = append(stacks[inst[2] - 1], forklift...)
    }
}

func getInstruction(i string) [3]int {
    var inst [3]int
    splt := strings.Split(i, " ")
    inst[0], _ = strconv.Atoi(splt[1])
    inst[1], _ = strconv.Atoi(splt[3]) 
    inst[2], _ = strconv.Atoi(splt[5]) 
    return inst
}
