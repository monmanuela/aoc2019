package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

func day3_part1() int {
	file, err := os.Open("day3.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	path, _ := reader.ReadString('\n')
	path2, _ := reader.ReadString('\n')

	passed := map[point]bool{}
	start := point{
		x: 0,
		y: 0,
	}
	for _, movt := range strings.Split(path, ",") {
		step, _ := strconv.ParseInt(string(movt[1:]), 10, 0)
		switch dir := string(movt[0]); dir {
		case "U":
			for i := start.y; i <= start.y + int(step); i++ {
				passed[point{
					x: start.x,
					y: i,
				}] = true
			}
			start = point {
				x: start.x,
				y: start.y + int(step),
			}
		case "R":
			for i := start.x; i <= start.x + int(step); i++ {
				passed[point{
					x: i,
					y: start.y,
				}] = true
			}
			start = point {
				x: start.x + int(step),
				y: start.y,
			}
		case "D":
			for i := start.y; i >= start.y - int(step); i-- {
				passed[point{
					x: start.x,
					y: i,
				}] = true
			}
			start = point {
				x: start.x,
				y: start.y - int(step),
			}
		case "L":
			for i := start.x; i >= start.x - int(step); i-- {
				passed[point{
					x: i,
					y: start.y,
				}] = true
			}
			start = point {
				x: start.x - int(step),
				y: start.y,
			}
		default:
			fmt.Println("Unknown direction")
		}
	}

	start = point{
		x: 0,
		y: 0,
	}
	crossing := map[point]bool{}
	for _, movt := range strings.Split(path2, ",") {
		step, _ := strconv.ParseInt(string(movt[1:]), 10, 0)
		switch dir := string(movt[0]); dir {
		case "U":
			for i := start.y; i <= start.y + int(step); i++ {
				if passed[point{
					x: start.x,
					y: i,
				}] {
					crossing[point{
						x: start.x,
						y: i,
					}] = true
				}
			}
			start = point {
				x: start.x,
				y: start.y + int(step),
			}
		case "R":
			for i := start.x; i <= start.x + int(step); i++ {
				if passed[point{
					x: i,
					y: start.y,
				}] {
					crossing[point{
						x: i,
						y: start.y,
					}] = true
				}
			}
			start = point {
				x: start.x + int(step),
				y: start.y,
			}
		case "D":
			for i := start.y; i >= start.y - int(step); i-- {
				if passed[point{
					x: start.x,
					y: i,
				}] {
					crossing[point{
						x: start.x,
						y: i,
					}] = true
				}
			}
			start = point {
				x: start.x,
				y: start.y - int(step),
			}
		case "L":
			for i := start.x; i >= start.x - int(step); i-- {
				if passed[point{
					x: i,
					y: start.y,
				}] {
					crossing[point{
						x: i,
						y: start.y,
					}] = true
				}
			}
			start = point {
				x: start.x - int(step),
				y: start.y,
			}
		default:
			fmt.Println("Unknown direction")
		}
	}

	closest := math.MaxInt16
	for point, _ := range crossing {
		if point.x == 0 && point.y == 0 {
			continue
		}
		if point.x < 0 {
			point.x = point.x * -1
		}
		if point.y < 0 {
			point.y = point.y * -1
		}
		if (point.x + point.y < closest) {
			closest = point.x + point.y
		}
	}
	return closest
}

func day3_part2() int {
	file, err := os.Open("day3.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	path, _ := reader.ReadString('\n')
	path2, _ := reader.ReadString('\n')

	passed := map[point]int{}
	start := point{
		x: 0,
		y: 0,
	}
	step_count := 1
	for _, movt := range strings.Split(path, ",") {
		step, _ := strconv.ParseInt(string(movt[1:]), 10, 0)
		switch dir := string(movt[0]); dir {
		case "U":
			for i := start.y + 1; i <= start.y + int(step); i++ {
				passed[point{
					x: start.x,
					y: i,
				}] = step_count
				step_count++
			}
			start = point {
				x: start.x,
				y: start.y + int(step),
			}
		case "R":
			for i := start.x + 1; i <= start.x + int(step); i++ {
				passed[point{
					x: i,
					y: start.y,
				}] = step_count
				step_count++
			}
			start = point {
				x: start.x + int(step),
				y: start.y,
			}
		case "D":
			for i := start.y - 1; i >= start.y - int(step); i-- {
				passed[point{
					x: start.x,
					y: i,
				}] = step_count
				step_count++
			}
			start = point {
				x: start.x,
				y: start.y - int(step),
			}
		case "L":
			for i := start.x - 1; i >= start.x - int(step); i-- {
				passed[point{
					x: i,
					y: start.y,
				}] = step_count
				step_count++
			}
			start = point {
				x: start.x - int(step),
				y: start.y,
			}
		default:
			fmt.Println("Unknown direction")
		}
	}

	start = point{
		x: 0,
		y: 0,
	}
	step_count = 1
	crossing := map[point]int{}
	for _, movt := range strings.Split(path2, ",") {
		step, _ := strconv.ParseInt(string(movt[1:]), 10, 0)
		switch dir := string(movt[0]); dir {
		case "U":
			for i := start.y + 1; i <= start.y + int(step); i++ {
				if passed[point{
					x: start.x,
					y: i,
				}] > 0 {
					crossing[point{
						x: start.x,
						y: i,
					}] = passed[point{
						x: start.x,
						y: i,
					}] + step_count
				}
				step_count++
			}
			start = point {
				x: start.x,
				y: start.y + int(step),
			}
		case "R":
			for i := start.x + 1; i <= start.x + int(step); i++ {
				if passed[point{
					x: i,
					y: start.y,
				}] > 0 {
					crossing[point{
						x: i,
						y: start.y,
					}] = passed[point{
						x: i,
						y: start.y,
					}] + step_count
				}
				step_count++
			}
			start = point {
				x: start.x + int(step),
				y: start.y,
			}
		case "D":
			for i := start.y - 1; i >= start.y - int(step); i-- {
				if passed[point{
					x: start.x,
					y: i,
				}] > 0 {
					crossing[point{
						x: start.x,
						y: i,
					}] = passed[point{
						x: start.x,
						y: i,
					}] + step_count
				}
				step_count++
			}
			start = point {
				x: start.x,
				y: start.y - int(step),
			}
		case "L":
			for i := start.x - 1; i >= start.x - int(step); i-- {
				if passed[point{
					x: i,
					y: start.y,
				}] > 0 {
					crossing[point{
						x: i,
						y: start.y,
					}] = passed[point{
						x: i,
						y: start.y,
					}] + step_count
				}
				step_count++
			}
			start = point {
				x: start.x - int(step),
				y: start.y,
			}
		default:
			fmt.Println("Unknown direction")
		}
	}

	closest := math.MaxInt64
	for point, stepFromOrigin := range crossing {
		if point.x == 0 && point.y == 0 {
			continue
		}
		if (stepFromOrigin < closest) {
			closest = stepFromOrigin
		}
	}
	return closest
}

func main() {
	fmt.Println(day3_part2())
}