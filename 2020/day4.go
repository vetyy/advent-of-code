package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	allowedEyeColors := map[string]bool{
		"amb": true,
		"blu": true,
		"brn": true,
		"gry": true,
		"grn": true,
		"hzl": true,
		"oth": true,
	}

	inputData, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatalf("failed to read input file: %v", err)
	}

	validPassportsCounterPart1 := 0
	validPassportsCounterPart2 := 0
	validFieldsCounterPart1 := 0
	validFieldsCounterPart2 := 0
	for _, passport := range strings.Split(string(inputData), "\n\n") {
		for _, line := range strings.Split(passport, "\n") {
			if len(line) == 0 {
				break
			}

			for _, field := range strings.Split(line, " ") {
				fieldParts := strings.Split(field, ":")
				fieldID, fieldVal := fieldParts[0], fieldParts[1]

				switch fieldID {
				case "byr":
					val := Atoi(fieldVal)
					if val >= 1920 && val <= 2002 {
						validFieldsCounterPart2 += 1
					}
					validFieldsCounterPart1 += 1
				case "iyr":
					val := Atoi(fieldVal)
					if val >= 2010 && val <= 2020 {
						validFieldsCounterPart2 += 1
					}
					validFieldsCounterPart1 += 1
				case "eyr":
					val := Atoi(fieldVal)
					if val >= 2020 && val <= 2030 {
						validFieldsCounterPart2 += 1
					}
					validFieldsCounterPart1 += 1
				case "hgt":
					if strings.HasSuffix(fieldVal, "cm") {
						val := Atoi(strings.TrimRight(fieldVal, "cm"))
						if val >= 150 && val <= 193 {
							validFieldsCounterPart2 += 1
						}
					} else if strings.HasSuffix(fieldVal, "in") {
						val := Atoi(strings.TrimRight(fieldVal, "in"))
						if val >= 59 && val <= 76 {
							validFieldsCounterPart2 += 1
						}
					}
					validFieldsCounterPart1 += 1
				case "hcl":
					matched, err := regexp.MatchString(`^#[0-9a-f]{6}$`, fieldVal)
					if err != nil {
						log.Fatalf("failed to compile regexp: %v", err)
					}
					if matched {
						validFieldsCounterPart2 += 1
					}
					validFieldsCounterPart1 += 1
				case "ecl":
					if allowedEyeColors[fieldVal] {
						validFieldsCounterPart2 += 1
					}
					validFieldsCounterPart1 += 1
				case "pid":
					matched, err := regexp.MatchString(`^[0-9]{9}$`, fieldVal)
					if err != nil {
						log.Fatalf("failed to compile regexp: %v", err)
					}
					if matched {
						validFieldsCounterPart2 += 1
					}
					validFieldsCounterPart1 += 1
				}
			}
		}
		if validFieldsCounterPart1 == 7 {
			validPassportsCounterPart1 += 1
		}
		if validFieldsCounterPart2 == 7 {
			validPassportsCounterPart2 += 1
		}
		validFieldsCounterPart1 = 0
		validFieldsCounterPart2 = 0
	}

	fmt.Println("Valid passports found in part 1:", validPassportsCounterPart1)
	fmt.Println("Valid passports found in part 2:", validPassportsCounterPart2)
}

func Atoi(strVal string) int {
	val, err := strconv.Atoi(strVal)
	if err != nil {
		log.Fatalf("invalid integer value provided: %v", err)
	}
	return val
}
