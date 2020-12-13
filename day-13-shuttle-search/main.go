package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "strconv"
    "strings"
)

func main() {
    inputData, err := ioutil.ReadFile("input")
    if err != nil {
        log.Fatalf("failed to read input file: %v", err)
    }
    data := strings.Split(string(inputData), "\n")
    timestamp, _ := strconv.Atoi(data[0])

    var busIDs, gaps []int
    for i, busID := range strings.Split(data[1], ",") {
        if busID == "x" {
            continue
        }
        gaps = append(gaps, i)
        value, _ := strconv.Atoi(busID)
        busIDs = append(busIDs, value)
    }

    var smallestBusID int
    smallestDiff := len(busIDs)
    for i, busID := range busIDs {
        diff := busID - (timestamp % busID)
        if diff < smallestDiff {
            smallestDiff = diff
            smallestBusID = i
        }
    }
    fmt.Println(busIDs[smallestBusID] * smallestDiff)

    var increment int
    min := busIDs[0]
    for i, id := range busIDs[1:] {
        for {
            increment += min
            if (increment+gaps[i+1])%id == 0 {
                break
            }
        }
        min *= id
    }
    fmt.Println(increment)
}
