package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var totalFuel int64
	input, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		mass, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		totalFuel = totalFuel + fuelCalculator(mass) 
	}
	fmt.Println(totalFuel)
}

func fuelCalculator(mass int64) int64 {
	var fuel int64
	for currentFuel := mass/3 - 2; currentFuel > 0; {
		fuel = fuel + currentFuel
		currentFuel = currentFuel/3 - 2
	}
	return fuel
}