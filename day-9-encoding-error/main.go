package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
)

const PreambleSize = 25

func main() {
    file, err := os.Open("input")
    if err != nil {
        log.Fatalf("failed to open input file: %v", err)
    }
    defer file.Close()

    var i, invalidNumber int
    var index []int
    slidingIndex := make([]int, PreambleSize)

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        value, _ := strconv.Atoi(scanner.Text())
        if i >= PreambleSize && !findInvalidNumber(slidingIndex, value) {
            invalidNumber = value
            fmt.Println("Part 1:", value)
            break
        }
        i += 1
        slidingIndex[i%PreambleSize] = value
        index = append(index, value)
    }

    for i, n1 := range index {
        min, max := n1, n1
        for _, n2 := range index[i+1:] {
            n1 += n2
            if n2 > max {
                max = n2
            }
            if min > n2 {
                min = n2
            }
            if n1 == invalidNumber {
                fmt.Println("Part 2:", min+max)
                return
            }
        }
    }

}

func findInvalidNumber(index []int, invalidValue int) bool {
    for i, val := range index {
        for _, val2 := range index[i+1:] {
            if invalidValue-val-val2 == 0 {
                return true
            }
        }
    }
    return false
}
