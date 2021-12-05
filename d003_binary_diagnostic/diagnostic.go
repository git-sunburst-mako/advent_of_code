package main

import (
	"advent_of_code/utils"
	"fmt"
	"strings"
)

const INPUT_FILE string = "../input/diagnostic.input"
const TEST_INPUT string = "../input/diagnostic.test"

const OXYGEN = true // most common
const CO2 = false   // least common

type rate_generator struct {
	reduced_input     string
	gamma             int64
	epsilon           int64
	power_consumption int64
}

// builds an int array of the most frequent bit
// or X on a tie
func (r *rate_generator) gamma_epsilon(on int, off int) {
	if on > off {
		r.reduced_input += "1"
	} else if off > on {
		r.reduced_input += "0"
	} else {
		r.reduced_input += "X"
	}
}

func (r *rate_generator) calc_gamma() {
	g := strings.Replace(r.reduced_input, "X", "1", -1)
	r.gamma = utils.SBinarytoI64(g)
}

func (r *rate_generator) calc_epsilon() {
	e := utils.BitFlipper(strings.Replace(r.reduced_input, "X", "0", -1))

	r.epsilon = utils.SBinarytoI64(e)
}

func (r *rate_generator) calc_consumption() {
	r.power_consumption = r.gamma * r.epsilon
}

func (r *rate_generator) reading_filter(readings [][]bool, pos int, system bool) string {
	var zeroes, ones [][]bool
	if len(readings) == 1 {
		return utils.ArrBintoS(readings[0])
	}

	for i := 0; i < len(readings); i++ {
		if readings[i][pos] == false {
			zeroes = append(zeroes, readings[i])
		} else if readings[i][pos] == true {
			ones = append(ones, readings[i])
		}
	}

	if system == OXYGEN {
		if len(ones) >= len(zeroes) {
			return r.reading_filter(ones, pos+1, OXYGEN)
		} else {
			return r.reading_filter(zeroes, pos+1, OXYGEN)
		}
	} else if system == CO2 {
		if len(zeroes) <= len(ones) {
			return r.reading_filter(zeroes, pos+1, CO2)
		} else {
			return r.reading_filter(ones, pos+1, CO2)
		}
	} else {
		panic("SYSTEM MALFUNCTION")
	}
}

func solve_part_one(r *rate_generator, x *utils.InputReader) {
	x.ConvertToBinary()
	for l := 0; l < x.BinaryLength(); l++ {
		var on, off int = 0, 0
		for _, d := range x.BinData {
			if d[l] {
				on++
			} else {
				off++
			}
		}
		r.gamma_epsilon(on, off)
	}

	r.calc_gamma()
	r.calc_epsilon()
	r.calc_consumption()
}

func main() {
	var x *utils.InputReader = utils.NewInputReader(INPUT_FILE)
	var r *rate_generator = new(rate_generator)

	solve_part_one(r, x)

	fmt.Printf("Reduced Input: %s | Gamma: %d | Epsilon: %d | Consumption: %d \n",
		r.reduced_input, r.gamma, r.epsilon, r.power_consumption)

	oxygen := r.reading_filter(x.BinData, 0, OXYGEN)
	co2 := r.reading_filter(x.BinData, 0, CO2)
	lifesupport := utils.SBinarytoI64(oxygen) * utils.SBinarytoI64(co2)

	fmt.Printf("Oxygen: %s = %d | CO2: %s = %d | LifeSupport: %d \n",
		oxygen, utils.SBinarytoI64(oxygen), co2, utils.SBinarytoI64(co2), lifesupport)

}
