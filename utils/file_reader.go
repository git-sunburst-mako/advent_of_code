package utils

import (
	"bufio"
	"os"
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

type InputReader struct {
	FilePath   string
	StrData    []string
	IntData    []int
	InputCount int
}

func NewInputReader(input string) *InputReader {
	i := new(InputReader)
	i.FilePath = input
	i.readLines()
	i.InputCount = len(i.StrData)
	return i
}

func (i *InputReader) readLines() {
	file, _ := os.Open(i.FilePath)

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	i.StrData = lines
}

func (ir *InputReader) ConvertToInts() {
	var ints []int
	for i := range ir.StrData {
		j, _ := strconv.Atoi(ir.StrData[i])
		ints = append(ints, j)
	}
	ir.IntData = ints
}

func (ir *InputReader) ConvertToCommands() CommandList {
	var ret CommandList
	for i := range ir.StrData {
		j := NewCommand(ir.StrData[i])
		ret.Commands = append(ret.Commands, j)
	}
	return ret
}

// return sum of interger array
func Sum_Array(arr []int) int {
	var s int = 0
	for _, i := range arr {
		s += i
	}
	return s
}

// return first > second (bool)
func Arr_Sum_Greater(a, b []int) bool {
	return Sum_Array(a) > Sum_Array(b)
}
