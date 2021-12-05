package utils

import (
	"strconv"
	"strings"
)

type Command struct {
	Direction string
	Value     int
}

type CommandList struct {
	Commands []Command
}

func NewCommand(input string) Command {
	i := strings.Split(input, " ")
	v, _ := strconv.Atoi(i[1])
	c := Command{Direction: i[0], Value: v}
	return c
}
