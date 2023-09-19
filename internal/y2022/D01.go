package y2022

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func findMax(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	if len(numbers) == 1 {
		return numbers[0]
	}

	largestNumber := numbers[0]

	for _, x := range numbers {
		if x > largestNumber {
			largestNumber = x
		}
	}

	return largestNumber
}

func readFile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	return file
}

func PartOne() {
	file := readFile("internal/y2022/D01.data")

	scanner := bufio.NewScanner(file)

	var elf = 0
	var elves []int

	for scanner.Scan() {
		line := scanner.Text()

		if line != "" {
			lineAsNum, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			elf += lineAsNum
		} else {
			elves = append(elves, elf)
			elf = 0
		}
	}

	println(findMax(elves))
}
