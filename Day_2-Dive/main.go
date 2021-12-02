package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type direction string

const (
	UP      direction = "up"
	DOWN    direction = "down"
	FORWARD direction = "forward"
)

type Command struct {
	D direction
	V int
}

func main() {
	data, err := ioutil.ReadFile("Day_2-Dive/input")
	if err != nil {
		panic(err)
	}

	strs := strings.Split(string(data), "\n")
	cmds := make([]Command, len(strs))
	for i := range strs {
		t := strings.Split(strs[i], " ")
		if len(t) != 2 {
			continue
		}
		tv, _ := strconv.Atoi(t[1])
		cmds[i] = Command{
			D: direction(t[0]),
			V: tv,
		}
	}

	// Task 1
	x, y := 0, 0
	for i := range cmds {
		switch cmds[i].D {
		case DOWN:
			y += cmds[i].V
			break
		case UP:
			y -= cmds[i].V
			break
		case FORWARD:
			x += cmds[i].V
			break
		}
	}

	fmt.Println(x * y)

	// Task 2
	aim, x2, y2 := 0, 0, 0
	for i := range cmds {
		switch cmds[i].D {
		case DOWN:
			aim += cmds[i].V
			break
		case UP:
			aim -= cmds[i].V
			break
		case FORWARD:
			x2 += cmds[i].V
			y2 += aim * cmds[i].V
			break
		}
	}

	fmt.Println(x2 * y2)
}
