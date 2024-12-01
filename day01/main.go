package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)
	sum := 0

	var firstCol []int
	var secondCol []int

	for fileScanner.Scan() {
		line := fileScanner.Text()
		parts := strings.Fields(line)
		parts0, _ := strconv.Atoi(parts[0])
		parts1, _ := strconv.Atoi(parts[1])
		firstCol = append(firstCol, parts0)
		secondCol = append(secondCol, parts1)

	}

	sort.Ints(firstCol)
	sort.Ints(secondCol)

	for i := 0; i < len(firstCol); i++ {
		diff := firstCol[i] - secondCol[i]
		if diff < 0 {
			diff = -diff
		}
		sum += diff
	}
	fmt.Printf("sum %v\n", sum)
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
