package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "strconv"
    "strings"
)

const BagColor = "shiny gold"

func main() {
    inputData, err := ioutil.ReadFile("input")
    if err != nil {
        log.Fatalf("failed to read input file: %v", err)
    }

    colorIndex := map[string]map[string]int{}
    for _, line := range strings.Split(strings.TrimSpace(string(inputData)), "\n") {
        lineParts := strings.Split(strings.Trim(line, "."), " bags contain ")

        bagColorKey := lineParts[0]
        colorIndex[bagColorKey] = map[string]int{}
        for _, color := range strings.Split(lineParts[1], ",") {
            colorParts := strings.Split(strings.TrimSpace(color), " ")
            colorKey := fmt.Sprintf("%s %s", colorParts[1], colorParts[2])
            colorCount, _ := strconv.Atoi(colorParts[0])
            colorIndex[bagColorKey][colorKey] = colorCount
        }
    }

    countPart1 := 0
    for name := range colorIndex {
        countPart1 += findBagByColor(colorIndex, name)
    }
    fmt.Println(countPart1)

    countPart2 := countBags(colorIndex, BagColor)
    fmt.Println(countPart2 - 1)
}

func findBagByColor(rules map[string]map[string]int, color string) int {
    for nextColor := range rules[color] {
        if nextColor == BagColor || findBagByColor(rules, nextColor) == 1 {
            return 1
        }
    }
    return 0
}

func countBags(colorIndex map[string]map[string]int, color string) int {
    cnt := 1
    for nextColor, count := range colorIndex[color] {
        cnt += count * countBags(colorIndex, nextColor)
    }
    return cnt
}
