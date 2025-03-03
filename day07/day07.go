package main 


import (
    "fmt"
    "aoc2022/data" 
    "strings"
    "strconv"
)


var input []string 

func main() {
    input = data.Get("data/day07.txt")   

    fmt.Println("p1:", part1()) 
    fmt.Println("p2:", part2())
}


var currentDir *dir 
var ptr int = 1

func part1() int {

    currentDir = new(dir) 
    currentDir.name = "/" 

    rootDir := currentDir

    for ptr < len(input) {
        readInput()
    }

    currentDir = rootDir

    return findSize(rootDir)
}

func findSize(r *dir) int {
    res := 0
    for _, dir := range r.childDirs {
        res += findSize(dir)
    }
    if r.size <= 100000 {
        res += r.size
    }
    if r.parentDir != nil {
        r.parentDir.size += r.size
    }
    return res
}

func readInput() {
    cmd := strings.Split(input[ptr], " ")

    if cmd[0] != "$" {
        panic(input[ptr])
    }

    if cmd[1] == "ls" {
        ptr++ 
        readOutput(currentDir)
    } else if cmd[1] == "cd" {
        if cmd[2] == ".." {
            currentDir = currentDir.parentDir
        } else {
            nDir := new(dir)
            nDir.parentDir = currentDir
            nDir.name = cmd[2]

            currentDir.childDirs = append(currentDir.childDirs, nDir)
            currentDir = nDir
        }
        ptr++
    } else {
        panic(input[ptr])
    }
}

func readOutput(cDir *dir) {
    cmd := strings.Split(input[ptr], " ") 

    for ptr < len(input) && cmd[0] != "$" {
        if cmd[0] != "dir" {
            num, err := strconv.Atoi(cmd[0])
            if err != nil {
                panic(err)
            }

            cDir.size += num
        }
        ptr++
        if ptr < len(input) { cmd = strings.Split(input[ptr], " ") }
    }
}

var smallest int = 70000001
var target int

func part2() int {
    target = 70000000 - currentDir.size 
    target = 30000000 - target
    findSmallest(currentDir)
    return smallest
}

func findSmallest(cDir *dir) {
    for _, dir := range cDir.childDirs {
        findSmallest(dir)
    }

    if cDir.size > target && cDir.size < smallest {
        smallest = cDir.size
    }
}


type dir struct {
    name string 
    size int 
    childDirs []*dir
    parentDir *dir
}
