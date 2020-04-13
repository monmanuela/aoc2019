package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("day6.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	parent_child_map := map[string][]string{}
	child_parent_map := map[string]string {}
	for scanner.Scan() {
		orbits := scanner.Text()
		orbit := strings.Split(orbits, ")")
		parent := orbit[0]
		child := orbit[1]
		parent_child_map[parent] = append(parent_child_map[parent], child)
		child_parent_map[child] = parent
	}

	fmt.Println(countMinOrbitalTransfers(parent_child_map, child_parent_map))
}

func countMinOrbitalTransfers(parent_child_map map[string][]string, child_parent_map map[string]string) int {
	root := "COM"
	you := "YOU"
	san := "SAN"
	var you_to_root []string
	var san_to_root []string
	you_to_root = append(you_to_root, you)
	san_to_root = append(san_to_root, san)

	parent := child_parent_map[you]
	for parent != root {
		you_to_root = append(you_to_root, parent)
		parent = child_parent_map[parent]
	}
	parent = child_parent_map[san]
	for parent != root {
		san_to_root = append(san_to_root, parent)
		parent = child_parent_map[parent]
	}

	// Trace from the back to find the branching point
	pt_san := len(san_to_root) - 1
	pt_you := len(you_to_root) - 1
	for san_to_root[pt_san] == you_to_root[pt_you] {
		pt_san--
		pt_you--
	}

	return pt_san + pt_you
}

func countOrbit(parent_child_map map[string][]string, child_parent_map map[string]string) int {
	root := "COM"

	orbit_count := map[string]int {}
	orbit_count[root] = 0
	queue := list.New()
	for _, child := range parent_child_map[root] {
		queue.PushBack(child)
	}

	for queue.Len() > 0 {
		next := queue.Front()
		str := fmt.Sprintf("%v", next.Value)
		orbit_count[str] = orbit_count[child_parent_map[str]] + 1
		queue.Remove(next)
		for _, child := range parent_child_map[str] {
			queue.PushBack(child)
		}
	}

	count := 0
	for _, ct := range orbit_count {
		count += ct
	}
	return count
}
