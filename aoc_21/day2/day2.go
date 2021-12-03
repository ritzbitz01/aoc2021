package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var horizontal = 0
	var depth = 0
	var aim = 0

	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		switch s[0] {
		case "forward":
			x, err := strconv.Atoi(s[1])
			if err != nil {
				fmt.Println("ERROR CONVERTING TO INT: ", s[0], s[1])
			}
			horizontal = horizontal + x
			depth = depth + (aim * x)
		case "down":
			x, err := strconv.Atoi(s[1])
			if err != nil {
				fmt.Println("ERROR CONVERTING TO INT: ", s[0], s[1])
			}
			// depth = depth + x
			aim = aim + x
		case "up":
			x, err := strconv.Atoi(s[1])
			if err != nil {
				fmt.Println("ERROR CONVERTING TO INT: ", s[0], s[1])
			}
			// depth = depth - x
			aim = aim - x
		}
	}

	fmt.Println("HORIZONTAL: ", horizontal)
	fmt.Println("DEPTH: ", depth)
	fmt.Println("TOTAL: ", horizontal*depth)

}
