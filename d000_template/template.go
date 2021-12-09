package main

import (
	"advent_of_code/utils"
)

const CHALLENGE string = "../input/x.input"
const TEST string = "../input/x.test"

func main() {
	var x *utils.InputReader = utils.NewInputReader(TEST)

	x.ConvertToInts()

}
