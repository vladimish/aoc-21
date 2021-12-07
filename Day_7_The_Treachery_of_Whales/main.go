package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("Day_7_The_Treachery_of_Whales/input")
	if err != nil {
		panic(err)
	}

	strs := strings.Split(string(data), ",")

	input := make([]int, len(strs))
	for i := range strs {
		input[i], err = strconv.Atoi(strs[i])
		if err != nil {
			fmt.Println(input[i], err)
		}
	}

	// Task 1
	min := math.MaxInt
	max := 0

	for i := range input {
		if input[i] > max {
			max = input[i]
		}
	}

	for i := 0; i < max; i++ {
		t := 0
		for k := range input {
			t += int(math.Abs(float64(i) - float64(input[k])))
		}
		if t < min {
			min = t
		}
		//fmt.Println(t)
	}

	fmt.Println(min)

	// Task 2
	min = math.MaxInt
	for i := 0; i < max; i++ {
		t := 0
		for k := range input {
			smth := int(math.Abs(float64(i)-float64(input[k]))) + 1
			smth = smth * (smth - 1) / 2
			t += smth
		}
		if t < min {
			min = t
		}
		//fmt.Println(t)
	}

	fmt.Println(min)
}
