package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var total = 0
	var result []int
	for scanner.Scan() {

		for scanner.Scan() {
			x, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("ERROR CONVERTING INPUT")
			}
			result = append(result, x)
		}
	}

	for i := 1; i < len(result); i++ {
		if result[i] > result[i-1] {
			total++
		}
	}

	fmt.Println("TOTAL PART 1: ", total)

	var total2 = 0
	for i := 3; i < len(result); i++ {
		if (result[i] + result[i-1] + result[i-2]) > (result[i-1] + result[i-2] + result[i-3]) {
			total2++
		}
	}
	fmt.Println("TOTAL PART 2: ", total2)
}
