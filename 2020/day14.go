package main

import (
    "fmt"
    "io/ioutil"
    "strconv"
    "strings"
)

func Pow(base, exponent int) int {
    result := 1
    for i := 0; i < exponent; i++ {
        result *= base
    }
    return result
}

func iterateFloating(mem map[string]int, intValue int, floating []int, value int) {
    if len(floating) == 0 {
        mem[strconv.Itoa(value)] = intValue
    } else {
        iterateFloating(mem, intValue, floating[1:], value)
        iterateFloating(mem, intValue, floating[1:], value+Pow(2, floating[0]))
    }
}

func doPart2(mem map[string]int, intValue int, mask string, value int) {
    var newVal int
    var floating []int
    for i := 0; i < len(mask); i++ {
        switch mask[len(mask)-i-1] {
        case 'X':
            floating = append(floating, i)
        case '0':
            newVal += value & Pow(2, i)
        case '1':
            newVal += Pow(2, i)
        }
    }
    iterateFloating(mem, intValue, floating, newVal)
}

func doPart1(mask string, intValue int) (newValue int) {
    for i := 0; i < len(mask); i++ {
        switch mask[len(mask)-i-1] {
        case 'X':
            newValue += intValue & Pow(2, i)
        case '1':
            newValue += Pow(2, i)
        }
    }
    return newValue
}

func solve(data []string, part2 bool) (sum int) {
    var mask string
    memory := map[string]int{}
    for _, line := range data {
        lineParts := strings.Split(line, " = ")
        key, value := lineParts[0], lineParts[1]
        if key == "mask" {
            mask = value
            continue
        }

        address := key[4 : len(key)-1]
        intValue, _ := strconv.Atoi(value)

        if !part2 {
            memory[address] = doPart1(mask, intValue)
        } else {
            intAddress, _ := strconv.Atoi(address)
            doPart2(memory, intValue, mask, intAddress)
        }
    }

    for _, val := range memory {
        sum += val
    }
    return sum
}

func main() {
    inputData, _ := ioutil.ReadFile("input")
    data := strings.Split(strings.TrimSpace(string(inputData)), "\n")

    fmt.Println("Part 1:", solve(data, false))
    fmt.Println("Part 2:", solve(data, true))
}
