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

	var masterList []string

	for scanner.Scan() {
		fullNum := scanner.Text()
		masterList = append(masterList, fullNum)
	}

	partOne(masterList)
	partTwo(masterList)

}

func partOne(numberList []string) {
	type Key struct {
		X, Y int
	}

	m := map[Key]int{}

	for _, s := range numberList {
		s := strings.Split(s, "")
		for i, c := range s {
			cval, _ := strconv.Atoi(c)
			m[Key{i, cval}] = m[Key{i, cval}] + 1
		}
	}

	var gammarate = ""
	var epsilonrate = ""
	for i := 0; i < 12; i++ {
		if m[Key{i, 0}] > m[Key{i, 1}] {
			gammarate = gammarate + "0"
			epsilonrate = epsilonrate + "1"
		} else {
			gammarate = gammarate + "1"
			epsilonrate = epsilonrate + "0"
		}
	}

	output, _ := strconv.ParseInt(gammarate, 2, 64)
	output2, _ := strconv.ParseInt(epsilonrate, 2, 64)
	fmt.Println("GAMMARATE: ", output)
	fmt.Println("EPSILONRATE: ", output2)

	powerconsumption := output * output2
	fmt.Println("POWER CONSUMPTION: ", powerconsumption)
}

func partTwo(numberList []string) {

	highList := numberList
	var i = 0
	for {
		highList = newList(highList, i, true)
		if len(highList) == 1 {
			break
		}
		i++
	}

	output, _ := strconv.ParseInt(highList[0], 2, 64)

	lowList := numberList
	var j = 0
	for {
		lowList = newList(lowList, j, false)
		if len(lowList) == 1 {
			break
		}
		j++
	}

	output2, _ := strconv.ParseInt(lowList[0], 2, 64)

	fmt.Println("LIFESUPPORT: ", output*output2)
}

func newList(fullList []string, place int, returnLow bool) []string {
	var zeroList []string
	var oneList []string

	for _, s := range fullList {
		chars := strings.Split(s, "")
		cval, _ := strconv.Atoi(chars[place])
		if cval == 0 {
			zeroList = append(zeroList, s)
		} else {
			oneList = append(oneList, s)
		}
	}

	if len(zeroList) > len(oneList) {
		if returnLow {
			return oneList
		} else {
			return zeroList
		}
	} else {
		if returnLow {
			return zeroList
		} else {
			return oneList
		}
	}
}
