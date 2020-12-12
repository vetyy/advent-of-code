package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
)

type Instruction struct {
    Action rune
    Value  int
}

func main() {
    file, err := os.Open("input")
    if err != nil {
        log.Fatalf("failed to open input file: %v", err)
    }
    defer file.Close()

    var instructions []*Instruction
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        value, _ := strconv.Atoi(line[1:])
        instructions = append(instructions, &Instruction{
            Action: rune(line[0]),
            Value:  value,
        })
    }

    part1(instructions)
    part2(instructions)
}

func part1(instructions []*Instruction) {
    directions := []rune{'E', 'S', 'W', 'N'}
    var ewPosition, nsPosition, currentDirection int
    for _, inst := range instructions {
        switch inst.Action {
        case 'N':
            nsPosition += inst.Value
        case 'S':
            nsPosition -= inst.Value
        case 'E':
            ewPosition += inst.Value
        case 'W':
            ewPosition -= inst.Value
        case 'L':
            currentDirection = (currentDirection - inst.Value/90 + 4) % 4
        case 'R':
            currentDirection = (currentDirection + inst.Value/90) % 4
        case 'F':
            switch directions[currentDirection] {
            case 'E':
                ewPosition += inst.Value
            case 'W':
                ewPosition -= inst.Value
            case 'N':
                nsPosition += inst.Value
            case 'S':
                nsPosition -= inst.Value
            }
        }
    }
    fmt.Println("Part 1:", Abs(ewPosition)+Abs(nsPosition))
}

func part2(instructions []*Instruction) {
    var ewPosition, nsPosition int
    wpEW, wpNS := 10, 1
    for _, inst := range instructions {
        switch inst.Action {
        case 'N':
            wpNS += inst.Value
        case 'S':
            wpNS -= inst.Value
        case 'E':
            wpEW += inst.Value
        case 'W':
            wpEW -= inst.Value
        case 'L':
            for i := 0; i < inst.Value/90; i++ {
                wpEW, wpNS = -wpNS, wpEW
            }
        case 'R':
            for i := 0; i < inst.Value/90; i++ {
                wpEW, wpNS = wpNS, -wpEW
            }
        case 'F':
            ewPosition += inst.Value * wpEW
            nsPosition += inst.Value * wpNS
        }
    }
    fmt.Println("Part 2:", Abs(ewPosition)+Abs(nsPosition))
}

func Abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
