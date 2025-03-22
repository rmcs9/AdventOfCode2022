package main 

import (
    "aoc2022/data" 
    "fmt"
    "strings"
    "strconv"
)

var input []string

func main() {
    input = data.Get("data/day09.txt")

    fmt.Println("p1:", part1()) 
    fmt.Println("p2:", part2()) 
}

func part1() int {
    seen := map[coord]bool{ {0, 0}: true, }
    H, T := coord{0, 0}, coord{0, 0}

    for _, ins := range input {
        spl := strings.Split(ins, " ")
        num, _ := strconv.Atoi(spl[1])

        for range num {
            H.x += direction[spl[0]].x 
            H.y += direction[spl[0]].y 

            for hloc, mv := range tailChase {
                if T.x + hloc.x == H.x && T.y + hloc.y == H.y {
                    T.x += mv.x 
                    T.y += mv.y 

                    seen[T] = true
                    break
                }
            }
        }
    }
    return len(seen)
}

func part2() int {
    seen := map[coord]bool{ {0,0}: true, } 
    knots := make([]coord, 10) 

    for _, ins := range input {
        spl := strings.Split(ins, " ") 
        num, _ := strconv.Atoi(spl[1])

        for range num {
            knots[0].x += direction[spl[0]].x 
            knots[0].y += direction[spl[0]].y 

            for i := 1; i < len(knots); i++ {
                for hloc, mv := range tailChase {
                    if knots[i].x + hloc.x == knots[i-1].x && knots[i].y + hloc.y == knots[i-1].y {
                        knots[i].x += mv.x 
                        knots[i].y += mv.y 
                        
                        if i == 9 {
                            seen[knots[i]] = true
                        }
                        break
                    }
                }
            }
        }
    }
    return len(seen)
}

type coord struct {
    x int 
    y int
}

var direction = map[string]coord {
    "U": { 0, -1},
    "D": { 0,  1}, 
    "R": { 1,  0}, 
    "L": {-1,  0},
}

var tailChase = map[coord]coord {
    //2 up
    {0, -2}: {0, -1}, 
    //2 down
    {0,  2}: {0,  1}, 
    //2 right
    {2,  0}: { 1, 0}, 
    //2 left
    {-2, 0}: {-1, 0}, 
    //2 up 1 right
    {1, -2}: {1, -1}, 
    //2 up 1 left
    {-1, -2}: {-1, -1}, 
    //2 down, 1 right
    {1, 2} : {1, 1},
    //2 down 1 left
    {-1, 2} : {-1, 1},
    //2 right 1 up 
    {2, -1} : {1, -1}, 
    //2 right 1 down 
    {2, 1}: {1, 1},
    //2left 1 up 
    {-2, -1}: {-1, -1},
    //2 left 1 down 
    {-2, 1}: {-1, 1}, 
    //p2 directions
    //right 2 up 2 
    {2, -2}: {1, -1}, 
    //right 2 down 2 
    {2, 2}: {1, 1}, 
    //left 2 up 2 
    {-2, -2}: {-1, -1}, 
    //left 2 down 2 
    {-2, 2}: {-1, 1},
}
