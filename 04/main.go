package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	matches := make(map[int]int)
	for index := 347312; index <= 805915; index++ {
		s := strings.Split(strconv.Itoa(index), "")
		tmp, pow, dif := 0, 6, 0
		for _, v := range(s) {
			pow--
			if a, _ := strconv.Atoi(v); a == tmp {
				matches[index]=index
			}
			if x, _ := strconv.Atoi(v); x < tmp {
				delete(matches, index)
				dif = tmp - x
				index = index + dif*int(math.Pow10(pow)) - 1
				break
			}
			tmp, _ = strconv.Atoi(v)
		}
	}
	fmt.Println(len(matches))
}
