package main

import (
    "fmt"
    "io/ioutil"
    "strconv"
    "strings"
)

func main() {
    inputData, _ := ioutil.ReadFile("input")
    data := strings.Split(strings.TrimSpace(string(inputData)), "\n\n")

    edgesByTitle := map[int]map[string]bool{}
    for _, block := range data {
        blockParts := strings.Split(block, "\n")
        piece := blockParts[1:]
        title := strings.Split(blockParts[0], " ")[1]
        titleID, _ := strconv.Atoi(title[:len(title)-1])

        edgesByTitle[titleID] = map[string]bool{
            getTopEdge(piece):                    true,
            getBottomEdge(piece):                 true,
            getLeftEdge(piece):                   true,
            getRightEdge(piece):                  true,
            reversedString(getTopEdge(piece)):    true,
            reversedString(getBottomEdge(piece)): true,
            reversedString(getLeftEdge(piece)):   true,
            reversedString(getRightEdge(piece)):  true,
        }
    }

    product := 1
    for title, edges := range edgesByTitle {
        counter := 0
        for title2, edges2 := range edgesByTitle {
            if title == title2 {
                continue
            }
            edgeCount := 0
            for edge := range edges {
                if edges2[edge] {
                    edgeCount += 1
                }
            }
            if edgeCount == 2 {
                counter += 1
            }
        }
        if counter == 2 {
            product *= title
        }
    }
    fmt.Println(product)
}

func getLeftEdge(block []string) string {
    var tmp []uint8
    for _, part := range block {
        tmp = append(tmp, part[0])
    }
    return string(tmp)
}

func getRightEdge(block []string) string {
    var tmp []uint8
    for _, part := range block {
        tmp = append(tmp, part[len(part)-1])
    }
    return string(tmp)
}

func getTopEdge(block []string) string {
    return block[0]
}

func getBottomEdge(block []string) string {
    return block[len(block)-1]
}

func reversedString(s string) string {
    tmp := ""
    for _, ch := range s {
        tmp = string(ch) + tmp
    }
    return tmp
}
