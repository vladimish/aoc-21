package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Status int

const (
	HIGHER  Status = 0
	SMALLER Status = 1
	EQUALS  Status = 2
)

func main() {
	data, err := ioutil.ReadFile("Day_3_Binary_Diagnostic/input")
	if err != nil {
		panic(err)
	}

	strs := strings.Split(string(data), "\n")

	// Part 1

	eps, gam := "", ""
	for i := 0; i < len(strs[0]); i++ {
		zer := 0
		one := 0
		for k := 0; k < len(strs); k++ {
			if strs[k][i]-48 == 0 {
				zer++
			} else {
				one++
			}
		}
		// fmt.Println(zer, one)
		if zer > one {
			eps += "1"
			gam += "0"
		} else {
			eps += "0"
			gam += "1"
		}
	}
	epsRes, _ := strconv.ParseInt(eps, 2, 32)
	gamRes, _ := strconv.ParseInt(gam, 2, 32)
	//fmt.Println(epsRes)
	//fmt.Println(gamRes)
	//fmt.Println(eps)
	//fmt.Println(gam)
	fmt.Println(epsRes * gamRes)

	// Part 2
	st := EQUALS
	backupStrs := make([]string, len(strs))
	copy(backupStrs, strs)
	for i := 0; i < len(strs[0]); i++ {
		if len(strs) != 1 {
			zer := 0
			one := 0
			for k := 0; k < len(strs); k++ {
				if strs[k][i]-48 == 0 {
					zer++
				} else {
					one++
				}
			}
			// fmt.Println(zer, one)
			if zer > one {
				st = HIGHER
			} else if zer < one {
				st = SMALLER
			} else {
				st = EQUALS
			}

			for k := 0; k < len(strs); k++ {
				if strs[k][i]-48 == 0 && st == SMALLER || strs[k][i]-48 == 1 && st == HIGHER || strs[k][i]-48 == 0 && st == EQUALS {
					strs = append(strs[:k], strs[k+1:]...)
					k--
				}
			}
		}

		if len(backupStrs) != 1 {
			zer := 0
			one := 0
			for k := 0; k < len(backupStrs); k++ {
				if backupStrs[k][i]-48 == 0 {
					zer++
				} else {
					one++
				}
			}
			// fmt.Println(zer, one)
			if zer > one {
				st = HIGHER
			} else if zer < one {
				st = SMALLER
			} else {
				st = EQUALS
			}

			for k := 0; k < len(backupStrs); k++ {
				if backupStrs[k][i]-48 == 1 && st == SMALLER || backupStrs[k][i]-48 == 0 && st == HIGHER || backupStrs[k][i]-48 == 1 && st == EQUALS {
					backupStrs = append(backupStrs[:k], backupStrs[k+1:]...)
					k--
				}
			}
		}
	}
	//fmt.Println(strs)
	//fmt.Println(backupStrs)
	a, _ := strconv.ParseInt(strs[0], 2, 32)
	b, _ := strconv.ParseInt(backupStrs[0], 2, 32)
	fmt.Println(a * b)
}
