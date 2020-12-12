package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

const Empty = 'L'
const Occupied = '#'
const Floor = '.'

func main() {
    fmt.Println("Part 1:", solve(false))
    fmt.Println("Part 2:", solve(true))
}

func solve(part2 bool) int {
    file, err := os.Open("input")
    if err != nil {
        log.Fatalf("failed to open input file: %v", err)
    }
    defer file.Close()

    var seats [][]rune
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        row := scanner.Text()
        var seatRow []rune
        for _, seat := range row {
            seatRow = append(seatRow, seat)
        }
        seats = append(seats, seatRow)
    }

    seatCount := 0
    for {
        seats = walk(seats, part2)
        tempCount := countOccupiedSeats(seats)
        if seatCount == tempCount {
            return seatCount
        }
        seatCount = tempCount
    }
}

func isOutOfBounds(x, y, rows, columns int) bool {
    return x < 0 || y < 0 || x >= rows || y >= columns
}

func walk(seats [][]rune, part2 bool) (newSeats [][]rune) {
    tolerance := 4
    if part2 {
        tolerance = 5
    }

    for i, row := range seats {
        var tempRow []rune
        for j, seat := range row {
            occupiedCount := 0
            for dx := -1; dx <= 1; dx++ {
                for dy := -1; dy <= 1; dy++ {
                    if dx != 0 || dy != 0 {
                        x := i + dx
                        y := j + dy
                        if part2 {
                            for !isOutOfBounds(x, y, len(seats), len(row)) && seats[x][y] == Floor {
                                x += dx
                                y += dy
                            }
                        }
                        if isOutOfBounds(x, y, len(seats), len(row)) {
                            continue
                        }
                        if seats[x][y] == Occupied {
                            occupiedCount += 1
                        }
                    }
                }
            }

            if seat == Empty && occupiedCount == 0 {
                tempRow = append(tempRow, Occupied)
            } else if seat == Occupied && occupiedCount >= tolerance {
                tempRow = append(tempRow, Empty)
            } else {
                tempRow = append(tempRow, seat)
            }
        }
        newSeats = append(newSeats, tempRow)
    }
    return newSeats
}

func countOccupiedSeats(seats [][]rune) (count int) {
    for _, row := range seats {
        for _, s := range row {
            if s == Occupied {
                count += 1
            }
        }
    }
    return count
}
