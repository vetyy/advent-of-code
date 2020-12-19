package main

import (
    "fmt"
    "io/ioutil"
    "strconv"
    "strings"
)

type EvalFunc func(characters []string) int

func main() {
    inputData, _ := ioutil.ReadFile("input")
    data := strings.Split(strings.TrimSpace(string(inputData)), "\n")
    fmt.Println("Part 1:", Solve(data, evaluatePart1))
    fmt.Println("Part 2:", Solve(data, evaluatePart2))
}

func Solve(data []string, eval EvalFunc) (out int) {
    for _, line := range data {
        line = strings.ReplaceAll(line, " ", "")
        var stack []string
        for i := 0; i < len(line); i++ {
            if line[i] == ')' {
                for j := len(stack) - 1; j >= 0; j-- {
                    if stack[j] == "(" {
                        tmpStack := stack[j:]
                        stack = stack[:j]
                        stack = append(stack, strconv.Itoa(eval(tmpStack)))
                        break
                    }
                }
            } else {
                stack = append(stack, string(line[i]))
            }
        }
        out += eval(stack)
    }
    return out
}

func Peek(stack []string) string {
    if len(stack) == 0 {
        return ""
    }
    return stack[len(stack)-1]
}

func evaluatePart1(characters []string) (result int) {
    var stack []string
    for _, ch := range characters {
        if ch == "+" || ch == "*" {
            stack = append(stack, ch)
        } else {
            value, _ := strconv.Atoi(ch)
            switch Peek(stack) {
            case "+":
                result += value
            case "*":
                result *= value
            default:
                result = value
            }
        }
    }
    return result
}

func evaluatePart2(characters []string) int {
    var prev string
    var stack []string
    for _, ch := range characters {
        stack = append(stack, ch)
        if prev == "+" {
            tmpStack := stack[len(stack)-3:]
            stack = stack[:len(stack)-3]
            stack = append(stack, strconv.Itoa(evaluatePart1(tmpStack)))
        }
        prev = ch
    }
    return evaluatePart1(stack)
}
