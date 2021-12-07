package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("Day_6_Lanternfish/input")
	if err != nil {
		panic(err)
	}

	strs := strings.Split(string(data), ",")

	amount := map[int]int{0: 0, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0}
	for i := range strs {
		res, err := strconv.Atoi(strs[i])
		if err != nil {
			fmt.Println(res, err)
		}
		amount[res]++
	}

	days := 256

	for i := 0; i < days; i++ {
		init := amount[0]
		for j := 1; j <= 9; j++ {
			amount[j-1] = amount[j]
		}
		amount[6] += init
		amount[8] += init

		//fmt.Println(amount)
	}

	s := 0
	for i := 0; i <= 8; i++ {
		s += amount[i]
	}

	fmt.Println(s)
}
