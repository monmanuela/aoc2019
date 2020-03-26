package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func day1_part1() {
	file, err := os.Open("day1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, _ := strconv.ParseInt(scanner.Text(), 10, 0)
		sum += int(num) / 3 - 2
	}
	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func recursiveFuelSum(init int) int {
	sum := 0
	start := init / 3 - 2
	for start > 0 {
		sum += start
		start = start / 3 - 2
	}
	return sum
}

func main() {
	file, err := os.Open("day1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, _ := strconv.ParseInt(scanner.Text(), 10, 0)
		sum += recursiveFuelSum(int(num))
	}
	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}