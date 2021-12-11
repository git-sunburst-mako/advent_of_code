package main

import (
	"advent_of_code/utils"
	"fmt"
	"math"
	"sort"
)

const CHALLENGE string = "../input/whales.input"
const TEST string = "../input/whales.test"

type crab_armada struct {
	positions     map[int]int
	fuel_required int
	count         int
	ordered_list  []int
	mean          map_calculation
	median        arr_calculation
}

type map_calculation func(map[int]int, int) int
type arr_calculation func([]int, int) int

func calc_mean(input map[int]int, count int) int {
	var sum int
	for k, v := range input {
		sum += k * v
	}
	return sum / count
}

func calc_median(input []int, count int) int {
	mid := count / 2
	return input[mid]
}

func (c *crab_armada) calculate_min_fuel(convergence int) int {
	var fuel int
	for k, v := range c.positions {
		dist := int(math.Sqrt(math.Pow(float64(k-convergence), 2)))
		mag := dist * v

		fuel += mag
	}
	return fuel
}

func (c *crab_armada) calculate_mean_fuel(convergence int) int {
	var fuel int
	for k, v := range c.positions {
		dist := int(math.Sqrt(math.Pow(float64(k-convergence), 2)))
		weighted_dist := dist * (dist + 1) / 2
		mag := weighted_dist * v

		fuel += mag
	}
	return fuel
}

func new_crab_aramada(crabs []int) *crab_armada {
	c := new(crab_armada)
	c.positions = make(map[int]int)
	for i, v := range crabs {
		println(i)
		c.positions[v] += 1
	}
	c.count = len(crabs)
	sort.Ints(crabs)
	c.ordered_list = crabs
	c.mean = calc_mean
	c.median = calc_median
	return c
}

func main() {
	var x *utils.InputReader = utils.NewInputReader(CHALLENGE)
	x.SplitStringToInts(",")

	armada := new_crab_aramada(x.IntData)
	mn := armada.mean(armada.positions, armada.count)
	med := armada.median(armada.ordered_list, armada.count)

	fuel_mn := armada.calculate_mean_fuel(mn)
	fmt.Println(armada.fuel_required, fuel_mn)

	fuel_med_min := armada.calculate_min_fuel(med + 2)
	fuel_med := armada.calculate_min_fuel(med)
	fuel_med_max := armada.calculate_min_fuel(med + 1)

	fmt.Println(fuel_med_min, fuel_med, fuel_med_max)

}
