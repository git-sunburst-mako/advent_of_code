package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type InputReader struct {
	FilePath   string
	StrData    []string
	IntData    []int
	BinData    [][]bool
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
	ir.IntData = ArrStoI(ir.StrData)
}

func (ir *InputReader) SplitStringToInts(delim string) {
	nums := strings.Split(ir.StrData[0], delim)
	ir.StrData = nums
	ir.ConvertToInts()
}

func (ir *InputReader) ConvertToCommands() CommandList {
	var ret CommandList
	for i := range ir.StrData {
		j := NewCommand(ir.StrData[i])
		ret.Commands = append(ret.Commands, j)
	}
	return ret
}

func (ir *InputReader) ConvertToBinary() {
	var data [][]bool
	ir.ConvertToInts()
	for _, i := range ir.StrData {
		bits := strings.Split(i, "")
		var datum []bool
		for _, j := range bits {
			b, _ := strconv.Atoi(j)
			datum = append(datum, ItoB(b))
		}
		data = append(data, datum)
	}
	ir.BinData = data
}

func (ir *InputReader) BinaryLength() int {
	return len(ir.BinData[0])
}
