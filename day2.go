package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func day2_part1(noun, verb int) int {
	file, err := os.Open("day2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	input, _ := reader.ReadString('\n')
	var numbers []int
	for _, num := range strings.Split(input, ",") {
		num, _ := strconv.ParseInt(num, 10, 0)
		numbers = append(numbers, int(num))
	}
	numbers[1] = noun
	numbers[2] = verb

	for i := 0; i < len(numbers); i += 4 {
		if numbers[i] == 99 {
			break
		} else if numbers[i] == 1 {
			numbers[numbers[i + 3]] = numbers[numbers[i + 1]] + numbers[numbers[i + 2]]
		} else if numbers[i] == 2 {
			numbers[numbers[i + 3]] = numbers[numbers[i + 1]] * numbers[numbers[i + 2]]
		} else {
			fmt.Println("Unexpected opcode")
		}
	}

	return numbers[0]
}

func main() {
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			res := day2_part1(noun, verb)
			if res == 19690720 {
				fmt.Println(noun, ", ", verb)
				fmt.Println(100 * noun + verb)
				break
			}
		}
	}
}