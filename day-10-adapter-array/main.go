package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "sort"
    "strconv"
)

func main() {
    file, err := os.Open("input")
    if err != nil {
        log.Fatalf("failed to open input file: %v", err)
    }
    defer file.Close()

    adapters := []int{0}
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        value, _ := strconv.Atoi(scanner.Text())
        adapters = append(adapters, value)
    }
    sort.Ints(adapters)
    adapters = append(adapters, adapters[len(adapters)-1]+3)

    var ones, threes int
    for i, adapter := range adapters[1:] {
        x := adapter - adapters[i]
        if x == 1 {
            ones += 1
        } else if x == 3 {
            threes += 1
        }
    }
    fmt.Println("Part 1:", ones*threes)

    index := map[int]int{0: 1}
    for _, adapter := range adapters[1:] {
        for i := 1; i <= 3; i++ {
            diff := adapter - i
            if _, ok := index[diff]; ok {
                index[adapter] += index[diff]
            }
        }
    }
    fmt.Println("Part 2:", index[adapters[len(adapters)-1]])
}
