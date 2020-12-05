package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

func main() {
    file, err := os.Open("input")
    if err != nil {
        log.Fatalf("failed to open input file: %v", err)
    }
    defer file.Close()

    maxSeatID := 0
    mySeatID := 0

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()

        row := 0
        for _, ch := range line[:7] {
            row *= 2
            if ch == 'B' {
                row += 1
            }
        }

        column := 0
        for _, ch := range line[7:] {
            column *= 2
            if ch == 'R' {
                column += 1

            }
        }

        seatID := row*8 + column
        if maxSeatID < seatID {
            maxSeatID = seatID
        }

        mySeatID ^= seatID
    }

    fmt.Println("Biggest seat ID is:", maxSeatID)
    fmt.Println("My seat ID is:", mySeatID-1)
}
