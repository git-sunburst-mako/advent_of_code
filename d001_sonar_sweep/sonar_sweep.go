package main

import (
	"advent_of_code/utils"
	"fmt"
)

const INPUT_FILE string = "../input/sonar_sweep.input"
const TEST_INPUT string = "../input/sonar_sweep.test"

func main() {
	var x *utils.InputReader = utils.NewInputReader(INPUT_FILE)
	x.ConvertToInts()

	// PART 1
	d1, i1 := count_depth_increases(x)
	fmt.Printf("pt1: decrease %d | increase %d | total readings %d \n", d1, i1, d1+i1)

	// PART 2
	d2, i2 := sliding_window(x, 3)
	fmt.Printf("pt2: decrease %d | increase %d | total readings %d \n", d2, i2, d2+i2)

}

// return decrease, increase
func count_depth_increases(ir *utils.InputReader) (int, int) {
	var decrease, increase int
	for i := 1; i < ir.InputCount; i++ {
		if ir.IntData[i] < ir.IntData[i-1] {
			decrease++
		} else if ir.IntData[i] > ir.IntData[i-1] {
			increase++
		}
	}
	return decrease, increase
}

// return decrease, increase
func sliding_window(ir *utils.InputReader, window int) (int, int) {
	var decrease, increase int
	for i := 0; i < ir.InputCount-window; i++ {
		j := i + window
		a := ir.IntData[i:j]
		b := ir.IntData[i+1 : j+1]
		if utils.Arr_Sum_Greater(a, b) {
			decrease++
		} else if utils.Arr_Sum_Greater(b, a) {
			increase++
		}
	}
	return decrease, increase
}
