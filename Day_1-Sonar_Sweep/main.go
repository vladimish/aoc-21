package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("Day_1-Sonar_Sweep/input")
	if err != nil {
		panic(err)
	}

	strs := strings.Split(string(data), "\n")
	ints := make([]int, len(strs))

	for i := range strs {
		ints[i], err = strconv.Atoi(strs[i])
		if err != nil {
			fmt.Println(ints[i], err)
		}
	}

	// Part 1
	res1 := 0
	for i := 0; i < len(ints)-1; i++ {
		if ints[i] < ints[i+1] {
			res1++
		}
	}

	fmt.Println(res1)

	// Part 2
	res2 := 0
	for i := 2; i < len(ints)-1; i++ {
		if ints[i]+ints[i-1]+ints[i-2] < ints[i+1]+ints[i]+ints[i-1] {
			res2++
		}
	}

	fmt.Println(res2)
}
