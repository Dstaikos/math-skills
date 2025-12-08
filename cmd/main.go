package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	mathskills "math-skills"
)

func main() {
	if len(os.Args) != 2 {

		fmt.Println("Usage: go run main.go <input.txt>")
		return
	}

	inputFile := os.Args[1]

	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("Error \n")
		return
	}

	defer file.Close()

	var numbers []int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		num, err := strconv.Atoi(line)
		if err != nil {
			fmt.Printf("Error converting string to int \n")
			return
		}

		numbers = append(numbers, num)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Scanner error:", err)
		return
	}

	fmt.Printf("Average: %d\n", mathskills.Average(numbers))
	fmt.Printf("Median: %d\n", mathskills.Median(numbers))
	fmt.Printf("Variance: %d\n", mathskills.Variance(numbers))
	fmt.Printf("Standard Deviation: %d\n", mathskills.Sdeviation(numbers))
}
