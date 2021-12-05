package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const SIZE = 1000

type Path struct {
	StartX int
	StartY int
	EndX   int
	EndY   int
}

func main() {
	file, err := ioutil.ReadFile("Day_5_Hydrothermal_Venture/input")
	if err != nil {
		panic(err)
	}

	strs := strings.Split(string(file), "\n")

	paths := make([]Path, len(strs))

	// Task 1
	for i := 0; i < len(strs); i++ {
		blocks := strings.Split(strs[i], "->")
		start := strings.Split(blocks[0], ",")
		end := strings.Split(blocks[1], ",")
		start[1] = strings.Replace(start[1], " ", "", 1)
		end[0] = strings.Replace(end[0], " ", "", 1)

		paths[i].StartX, err = strconv.Atoi(start[0])
		if err != nil {
			panic(err)
		}

		paths[i].StartY, err = strconv.Atoi(start[1])
		if err != nil {
			panic(err)
		}

		paths[i].EndX, err = strconv.Atoi(end[0])
		if err != nil {
			panic(err)
		}

		paths[i].EndY, err = strconv.Atoi(end[1])
		if err != nil {
			panic(err)
		}
	}

	//fmt.Println(paths)

	data1 := make([][]int, SIZE)
	rows1 := make([]int, SIZE*SIZE)
	for i := 0; i < SIZE; i++ {
		data1[i] = rows1[i*SIZE : (i+1)*SIZE]
	}

	data2 := make([][]int, SIZE)
	rows2 := make([]int, SIZE*SIZE)
	for i := 0; i < SIZE; i++ {
		data2[i] = rows2[i*SIZE : (i+1)*SIZE]
	}

	for i := range paths {
		smallerX, biggerX, smallerY, biggerY := 0, 0, 0, 0

		if paths[i].StartX > paths[i].EndX {
			biggerX = paths[i].StartX
			smallerX = paths[i].EndX
		} else {
			smallerX = paths[i].StartX
			biggerX = paths[i].EndX
		}

		if paths[i].StartY > paths[i].EndY {
			biggerY = paths[i].StartY
			smallerY = paths[i].EndY
		} else {
			smallerY = paths[i].StartY
			biggerY = paths[i].EndY
		}

		if paths[i].StartX == paths[i].EndX || paths[i].StartY == paths[i].EndY {
			for j := smallerX; j <= biggerX; j++ {
				for k := smallerY; k <= biggerY; k++ {
					data1[j][k]++
				}
			}
		}
	}

	res1 := 0
	for i := range data1 {
		for j := range data1 {
			if data1[i][j] >= 2 {
				res1++
			}
		}
	}

	fmt.Println(res1)

	// Task 2
	copy(data2, data1)
	for i := range paths {
		// DOWN RIGHT
		if paths[i].StartX < paths[i].EndX && paths[i].StartY < paths[i].EndY {
			if paths[i].EndX-paths[i].StartX == paths[i].EndY-paths[i].StartY {
				k := paths[i].StartY
				for j := paths[i].StartX; j <= paths[i].EndX; j++ {
					data2[j][k]++
					k++
				}
			}
			// UP RIGHT
		} else if paths[i].StartX < paths[i].EndX && paths[i].StartY > paths[i].EndY {
			if paths[i].EndX-paths[i].StartX == paths[i].StartY-paths[i].EndY {
				k := paths[i].StartY
				for j := paths[i].StartX; j <= paths[i].EndX; j++ {
					data2[j][k]++
					k--
				}
			}
			// DOWN LEFT
		} else if paths[i].StartX > paths[i].EndX && paths[i].StartY < paths[i].EndY {
			if paths[i].EndX-paths[i].StartX == paths[i].StartY-paths[i].EndY {
				k := paths[i].StartY
				for j := paths[i].StartX; j >= paths[i].EndX; j-- {
					data2[j][k]++
					k++
				}
			}
			// UP LEFT
		} else if paths[i].StartX > paths[i].EndX && paths[i].StartY > paths[i].EndY {
			if paths[i].StartX-paths[i].EndX == paths[i].StartY-paths[i].EndY {
				k := paths[i].StartY
				for j := paths[i].StartX; j >= paths[i].EndX; j-- {
					data2[j][k]++
					k--
				}
			}
		} else {
			fmt.Println(paths[i], "unreachable")
		}
	}

	res2 := 0
	for i := range data2 {
		for j := range data2 {
			if data2[i][j] >= 2 {
				res2++
			}
		}
	}

	fmt.Println(res2)
}
