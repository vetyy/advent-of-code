package main

import (
    "bufio"
    "fmt"
    "log"
    "math"
    "os"
    "sort"
)

func main() {
    file, err := os.Open("input")
    if err != nil {
        log.Fatalf("failed to open input file: %v", err)
    }
    defer file.Close()

    maxSeatID := 0.0
    var seats []float64

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        x := 63.
        row := 127.
        for _, ch := range line[:7] {
            row = row / 2
            if ch == 'F' {
                x -= row / 2
            } else {
                x += row / 2
            }
        }

        y := 3.
        column := 7.
        for _, ch := range line[7:] {
            column = column / 2
            if ch == 'L' {
                y -= column / 2
            } else {
                y += column / 2
            }
        }

        seatID := math.Ceil(x)*8 + math.Ceil(y)
        if maxSeatID < seatID {
            maxSeatID = seatID
        }
        seats = append(seats, seatID)
    }

    fmt.Println("Biggest seat ID is: ", maxSeatID)

    sort.Float64s(seats)
    for i, seatID := range seats {
        if seatID+1 != seats[i+1] {
            fmt.Println("Your seat ID is: ", seatID+1)
            break
        }
    }
}
