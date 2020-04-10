package main

import (
	"fmt"
	"strconv"
)

//It is a six-digit number.
//The value is within the range given in your puzzle input.
//Two adjacent digits are the same (like 22 in 122345).
//Going from left to right, the digits never decrease; they only ever increase or stay the same (like 111123 or 135679).
func day4_part1() {
	const start = 245318;
	const end = 765747;
	count := 0;
	for n := start; n <= end; n++ {
		if hasAdjacentDigits(n) && isIncreasing(n) {
			count++;
		}
	}
	fmt.Println("Total: ", count);
}

func hasAdjacentDigits(n int) bool {
	str := strconv.Itoa(n);
	result := false;
	for i := 0; i < len(str) - 1; i++ {
		if str[i] == str[i + 1] {
			result = true;
			break;
		}
	}
	return result;
}

func isIncreasing(n int) bool {
	str := strconv.Itoa(n);
	result := true;
	for i := 0; i < len(str) - 1; i++ {
		if str[i] > str[i + 1] {
			result = false;
			break;
		}
	}
	return result;
}

func day4_part2() {
	const start = 245318;
	const end = 765747;
	count := 0;
	for n := start; n <= end; n++ {
		if hasAdjacentDigitsModified(n) && isIncreasing(n) {
			count++;
		}
	}
	fmt.Println("Total: ", count);
}

func hasAdjacentDigitsModified(n int) bool {
	str := strconv.Itoa(n);
	result := false;
	for i := 0; i < len(str) - 1; i++ {
		if str[i] == str[i + 1] {
			result = true;
			//check left and right
			if (i > 0 && str[i - 1] == str[i]) || (i < len(str) - 2 && str[i + 2] == str[i + 1]) {
				result = false;
			}
			if result {
				break;
			}
		}
	}
	return result;
}
