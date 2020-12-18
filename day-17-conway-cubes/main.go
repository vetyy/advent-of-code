package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

type Coordinates struct {
    X, Y, Z, W int
}

func Neighbours(part2 bool) (final []Coordinates) {
    for x := -1; x <= 1; x++ {
        for y := -1; y <= 1; y++ {
            for z := -1; z <= 1; z++ {
                if part2 {
                    for w := -1; w <= 1; w++ {
                        final = append(final, Coordinates{x, y, z, w})
                    }
                } else {
                    final = append(final, Coordinates{x, y, z, 0})
                }
            }
        }
    }
    return final
}

func main() {
    file, err := os.Open("input")
    if err != nil {
        log.Fatalf("failed to open input file: %v", err)
    }
    defer file.Close()

    var x int
    cubes := map[Coordinates]struct{}{}
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        for y, ch := range line {
            if ch == '#' {
                cubes[Coordinates{x, y, 0, 0}] = struct{}{}
            }
        }
        x += 1
    }

    fmt.Println("Part 1:", solve(cubes, 6, false))
    fmt.Println("Part 2:", solve(cubes, 6, true))
}

func solve(cubes map[Coordinates]struct{}, cycles int, part2 bool) int {
    for i := 0; i < cycles; i++ {
        neighbourCubes := map[Coordinates]int{}
        for cube := range cubes {
            for _, c := range Neighbours(part2) {
                if c.X == 0 && c.Y == 0 && c.Z == 0 && c.W == 0 {
                    continue
                }
                nCoordinates := Coordinates{cube.X + c.X, cube.Y + c.Y, cube.Z + c.Z, cube.W + c.W}
                neighbourCubes[nCoordinates] += 1
            }
        }

        newCubes := map[Coordinates]struct{}{}
        for neighbour, count := range neighbourCubes {
            _, ok := cubes[neighbour]
            if count == 3 || (count == 2 && ok) {
                newCubes[neighbour] = struct{}{}
            }
        }
        cubes = newCubes
    }
    return len(cubes)
}
