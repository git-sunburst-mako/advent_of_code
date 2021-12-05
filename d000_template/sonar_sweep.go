package main

import (
	"advent_of_code/utils"
)

const INPUT_FILE string = "../input/x.input"
const TEST_INPUT string = "../input/x.test"

func main() {
	var x *utils.InputReader = utils.NewInputReader(INPUT_FILE)

	x.ConvertToInts()

}
