package main

import (
    "fmt"
    "io/ioutil"
    "regexp"
    "strings"
)

func parseFile() []map[string]string {
    inputData, _ := ioutil.ReadFile("input")
    var passports []map[string]string
    for _, passportData := range strings.Split(strings.TrimSpace(string(inputData)), "\n\n") {
        passport := map[string]string{}
        for _, line := range strings.Split(passportData, "\n") {
            for _, part := range strings.Split(line, " ") {
                keyval := strings.Split(part, ":")
                key, val := keyval[0], keyval[1]
                passport[key] = val
            }
        }
        passports = append(passports, passport)
    }
    return passports
}

func main() {
    required := map[string]*regexp.Regexp{
        "byr": regexp.MustCompile(`^(19[2-9][0-9]|200[0-2])$`),
        "iyr": regexp.MustCompile(`^(201[0-9]|2020)$`),
        "eyr": regexp.MustCompile(`^(202[0-9]|2030)$`),
        "hgt": regexp.MustCompile(`^(1[5-8][0-9]|19[0-3])cm|(59|6[0-9]|7[0-6])in$`),
        "hcl": regexp.MustCompile(`^#[0-9a-f]{6}$`),
        "ecl": regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`),
        "pid": regexp.MustCompile(`^\d{9}$`),
        // "cid",
    }

    passports := parseFile()
    total := len(passports)
    invalidPart1 := 0
    invalidPart2 := 0
    for _, passport := range passports {
        for reqKey := range required {
            if passport[reqKey] == "" {
                invalidPart1++
                break
            }
        }
        for reqKey, regex := range required {
            if val := passport[reqKey]; !regex.MatchString(val) {
                invalidPart2++
                break
            }
        }
    }
    fmt.Printf("Part 1 Total %d Valid %d\n", total, total-invalidPart1)
    fmt.Printf("Part 2 Total %d Valid %d\n", total, total-invalidPart2)
}
