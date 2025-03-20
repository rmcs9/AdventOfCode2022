package main 


import ( 
    "fmt" 
    "aoc2022/data" 
    "strconv"
)

var input []string

func main() {
    input = data.Get("data/day08.txt")


    fmt.Println("p1:", part1())
    fmt.Println("p2:", part2())
}

func part1() int {
    res := 2 * (len(input) + len(input[0])) - 4

    for y := 1; y < len(input) - 1; y++ {
        for x := 1; x < len(input[0]) - 1; x++ {
            for _, dir := range dirs {
                if val, _ := strconv.Atoi((string)(input[y][x])); visibleCheck(x, y, val, dir) {
                    res += 1 
                    break
                }
            }
        }
    }
    return res
}

func part2() int {
    spot := -1

    for y := 1; y < len(input) - 1; y++ {
        for x := 1; x < len(input[0]) - 1; x++ {
            thisspot := 1
            for _, dir := range dirs {
                val, _ := strconv.Atoi((string)(input[y][x]))
                thisspot *= viewDistance(x, y, val, dir)
            }

            if thisspot > spot {
                spot = thisspot
            }
        }
    }

    return spot
}

func viewDistance(x, y, val int, d dir) int {
    dist := 0 
    x += d.dx 
    y += d.dy

    for inBounds(x, y) {
        if num, _ := strconv.Atoi((string)(input[y][x])); val <= num { 
            // count the tree that blocked u :P
            dist++
            break
        }
        dist++ 
        x += d.dx 
        y += d.dy 
    }

    return dist
}

func visibleCheck(x, y, val int, d dir) bool {
    x += d.dx
    y += d.dy 

    for inBounds(x, y) {
        if num, _ := strconv.Atoi((string)(input[y][x])); val <= num {
            return false
        }
        x += d.dx 
        y += d.dy
    }
    return true
}

func inBounds(x, y int) bool {
    return x < len(input[0]) && y < len(input) && x >= 0 && y >= 0
}

type dir struct {
    dy int 
    dx int
}

type coord struct {
    y int 
    x int
}

var dirs = []dir {
    //N 
    { -1, 0 },
    //S 
    {  1, 0 },
    //E 
    {  0, 1 },
    //W
    {  0,-1 }, 
}
