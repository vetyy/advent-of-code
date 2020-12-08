package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"
)

type Instruction struct {
    Operation string
    Argument  int
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
        line := strings.Split(scanner.Text(), " ")
        value, _ := strconv.Atoi(line[1])
        instructions = append(instructions, &Instruction{Operation: line[0], Argument: value})
    }

    accumulator, _ := findLoop(instructions)
    fmt.Println("Part 1:", accumulator)

    accumulator, _ = fixBoot(instructions)
    fmt.Println("Part 2:", accumulator)
}

func findLoop(instructions []*Instruction) (accumulator int, ok bool) {
    var position int
    visited := map[int]bool{}

    for {
        visited[position] = true

        instruction := instructions[position]
        if instruction.Operation == "acc" {
            accumulator += instruction.Argument
        } else if instruction.Operation == "jmp" {
            position += instruction.Argument - 1
        }

        position += 1

        if visited[position] {
            return accumulator, true
        }

        if position == len(instructions) {
            return accumulator, false
        }
    }
}

func fixBoot(instructions []*Instruction) (accumulator int, ok bool) {
    for _, instruction := range instructions {
        operation := instruction.Operation
        if operation == "nop" {
            instruction.Operation = "jmp"
        } else if operation == "jmp" {
            instruction.Operation = "nop"
        }

        accumulator, ok := findLoop(instructions)
        if !ok {
            return accumulator, true
        }
        instruction.Operation = operation
    }
    return 0, false
}
