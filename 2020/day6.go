package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "strings"
)

func main() {
    inputData, err := ioutil.ReadFile("input")
    if err != nil {
        log.Fatalf("failed to read input file: %v", err)
    }

    resultPart1 := 0
    resultPart2 := 0
    for _, group := range strings.Split(string(inputData), "\n\n") {
        counter := map[rune]int{}
        peopleInGroup := strings.Split(strings.TrimSpace(group), "\n")
        for _, person := range peopleInGroup {
            for _, answer := range person {
                counter[answer] += 1
            }
        }
        for _, answerCount := range counter {
            if answerCount == len(peopleInGroup) {
                resultPart2 += 1
            }
        }
        resultPart1 += len(counter)
    }
    fmt.Println("Result of part 1:", resultPart1)
    fmt.Println("Result of part 2:", resultPart2)
}
