package main 


import ( 
    "aoc2022/data" 
    "fmt" 
)

var input []string

func main() {
    input = data.Get("data/day03.txt")

    fmt.Println("p1:", part1())
    fmt.Println("p2:", part2())
}

func part1() int {
    
    total := 0 

    for _, ruck := range input {
        sack := make(map[rune]bool)
        comp1 := ruck[:len(ruck) / 2] 
        comp2 := ruck[len(ruck) / 2:]

        for _, b := range comp1 {
            sack[b] = true   
        }

        found := make(map[rune]bool) 
        for _, b := range comp2 {
            if _, has := sack[b]; has {
                if _, f := found[b]; f { continue }

                if b < 96 {
                    total += int(b) - 38
                } else {
                    total += int(b) - 96
                }
                found[b] = true
            }
        }
    }

    return total
}

func part2() int {
    
    total := 0

    for len(input) != 0 {
        elfs := make([]string, 3)
        for i := range 3 {
            elfs[i] = input[0]
            input = input[1:]
        }

        groupSeen := make(map[rune]int)

        for i := range 3 {
            total += func() int {
                seen := make(map[rune]bool)
                for _, b := range elfs[i] {
                    if _, s := seen[b]; s { continue } 

                    if _, s := groupSeen[b]; !s {
                        groupSeen[b] = 1 
                    } else if groupSeen[b] == 2 {
                        if b < 96 {
                            return int(b) - 38
                        } else {
                            return int(b) - 96
                        }
                    } else {
                        groupSeen[b] = groupSeen[b] + 1
                    }
                    seen[b] = true
                }
                return 0
            }()
        }
    }

    return total
}
