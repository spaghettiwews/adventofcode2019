package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	var strCodes []string
	
	input, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		strCodes = strings.Split(scanner.Text(), ",")
	}
	
	for n := 0; n <= 99; n++ {
		for v := 0; v <= 99; v++ {
			if theProgram(makeInts(strCodes), n, v) == 19690720 {
				fmt.Printf("noun: %v verb: %v\n", n, v)
				fmt.Println(100 * n + v)
				break
			}
		}
	}
}

func theProgram(intCodes []int, noun int, verb int) int {
	
	intCodes[1] = noun
	intCodes[2] = verb

	for address := 0; address < len(intCodes) ; address+=4 {
		if intCodes[address] == 1 {
			intCodes[intCodes[address+3]] = intCodes[intCodes[address+1]] + intCodes[intCodes[address+2]]
		}
		if intCodes[address] == 2 {
			intCodes[intCodes[address+3]] = intCodes[intCodes[address+1]] * intCodes[intCodes[address+2]]
		}
		if intCodes[address] == 99 {
			break
		}
	}
	return intCodes[0]
}

func makeInts(in []string) []int {
	out := make([]int, len(in))

	for i := 0; i < len(in); i++ {
		out[i], _ = strconv.Atoi(in[i])
	}

	return out
}