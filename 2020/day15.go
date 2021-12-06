package main

import "fmt"

func main() {
    fmt.Println("Part 1:", memoryGame([]int{1, 20, 11, 6, 12, 0}, 2020))
    fmt.Println("Part 2:", memoryGame([]int{1, 20, 11, 6, 12, 0}, 30000000))
}

func memoryGame(input []int, nthNumber int) (num int) {
    seen := map[int][]int{}
    for i, n := range input {
        seen[n] = append(seen[n], i+1)
    }

    for i := len(seen) + 1; i <= nthNumber; i++ {
        prev := seen[num]
        if len(prev) <= 1 {
            num = 0
        } else {
            seen[num] = prev[1:]
            num = prev[1] - prev[0]
        }
        seen[num] = append(seen[num], i)
    }
    return num
}
