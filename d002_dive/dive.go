package main

import (
	"advent_of_code/utils"
	"fmt"
)

const INPUT_FILE string = "../input/dive.input"
const TEST_INPUT string = "../input/dive.test"

const FORWARD string = "forward"
const UP string = "up"
const DOWN string = "down"

type submarine struct {
	horizontal int
	depth      int
	aim        int
}

func (s *submarine) simple_move(c utils.Command) {
	switch c.Direction {
	case FORWARD:
		s.horizontal += c.Value
	case UP:
		s.depth -= c.Value
	case DOWN:
		s.depth += c.Value
	}
}

func (s *submarine) tilt_move(c utils.Command) {
	switch c.Direction {
	case FORWARD:
		s.horizontal += c.Value
		s.depth += s.dive(c.Value)
	case UP:
		s.aim -= c.Value
	case DOWN:
		s.aim += c.Value
	}
}

func (s *submarine) dive(dist int) int {
	return dist * s.aim
}

func (s *submarine) solution() int {
	return s.horizontal * s.depth
}

func (s *submarine) follow_instructions(c utils.CommandList, part int) {
	for _, v := range c.Commands {
		switch part {
		case 1:
			s.simple_move(v)
		case 2:
			s.tilt_move(v)
		}
	}
}

func main() {
	var x *utils.InputReader = utils.NewInputReader(INPUT_FILE)
	commands := x.ConvertToCommands()
	var sub_1 submarine

	sub_1.follow_instructions(commands, 1)
	pt1 := sub_1.solution()
	fmt.Printf("pt1: horizontal %d | depth %d | solution %d \n", sub_1.horizontal, sub_1.depth, pt1)

	var sub_2 submarine

	sub_2.follow_instructions(commands, 2)
	pt2 := sub_2.solution()
	fmt.Printf("pt2: horizontal %d | depth %d | solution %d \n", sub_2.horizontal, sub_2.depth, pt2)
}
