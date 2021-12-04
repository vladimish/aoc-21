package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("Day_4_Giant_Squid/input")
	if err != nil {
		panic(err)
	}

	strs := strings.Split(string(data), "\n")

	firstInts := strings.Split(strs[0], ",")
	numbers := make([]int, len(firstInts))
	for i := range firstInts {
		numbers[i], err = strconv.Atoi(firstInts[i])
		if err != nil {
			fmt.Println(err)
		}
	}

	mat := make([][][]int, len(strs)/6)
	for i := 0; i < len(strs)/6; i++ {
		mat[i] = make([][]int, 5)
		for k := 0; k < len(mat[0]); k++ {
			mat[i][k] = make([]int, 5)
		}
	}

	for i := 0; i < len(mat); i++ {
		for k := 0; k < len(mat[0]); k++ {
			current := strs[2+i*5+i+k]
			//fmt.Println(current)
			currentSplit := strings.Split(current, " ")
			for j := 0; j < len(currentSplit); j++ {
				if currentSplit[j] == "" {
					currentSplit = append(currentSplit[:j], currentSplit[j+1:]...)
				}
				mat[i][k][j], err = strconv.Atoi(currentSplit[j])
				if err != nil {
					panic(err)
				}
			}
		}
		//fmt.Println("===")
	}

	ban := make(map[int]int)

	for n := 0; n < len(numbers); n++ {
		for i := 0; i < len(mat); i++ {
			for k := 0; k < len(mat[0]); k++ {
				for j := 0; j < len(mat[0][0]); j++ {
					if numbers[n] == mat[i][k][j] {
						mat[i][k][j] = -1
					}
				}
			}
		}

		for i := 0; i < len(mat); i++ {
			_, ok := ban[i]
			if !ok {
				for k := 0; k < len(mat[0]); k++ {
					check1, check2 := true, true
					for j := 0; j < len(mat[0][0]); j++ {
						if mat[i][k][j] != -1 {
							check1 = false
						}
					}
					for j := 0; j < len(mat[0][0]); j++ {
						if mat[i][j][k] != -1 {
							check2 = false
						}
					}
					if check1 || check2 {
						res := 0
						for l := 0; l < len(mat[0]); l++ {
							for m := 0; m < len(mat[0][0]); m++ {
								if mat[i][l][m] != -1 {
									res += mat[i][l][m]
								}
							}
						}
						fmt.Println(i, n)
						fmt.Println(res * numbers[n])
						ban[i] = 1
						break
					}
				}
			}
		}
	}
}
