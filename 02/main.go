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
	
	intCodes := make([]int, len(strCodes))

	for i := 0; i < len(strCodes); i++ {
		intCodes[i], _ = strconv.Atoi(strCodes[i])
	}

	intCodes[1] = 12
	intCodes[2] = 2

	for i := 0; i < len(intCodes) ; i+=4 {
		if intCodes[i] == 1 {
			intCodes[intCodes[i+3]] = intCodes[intCodes[i+1]] + intCodes[intCodes[i+2]]
		}
		if intCodes[i] == 2 {
			intCodes[intCodes[i+3]] = intCodes[intCodes[i+1]] * intCodes[intCodes[i+2]]
		}
		if intCodes[i] == 99 {
			break
		}
	}
	fmt.Println(intCodes)
}

