package main

import (
	"advent_of_code/utils"
	"fmt"
	"strconv"
)

const CHALLENGE string = "../input/lanternfish.input"
const TEST string = "../input/lanternfish.test"
const SINGLE string = "../input/lanternfish.single"

const GESTATION_RATE int = 7
const NEWBORN int = GESTATION_RATE + 2
const DAYS int = 256

func status(day, fish, total int) {
	fmt.Printf("DAY: %d | THERE ARE: %d NEW FISH ~ THERE ARE: %d TOTAL: \n", day, fish, total)
}

func fish_array(adults [7]int, juveniles [9]int) {
	var summary_head string = "[ 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 ] \n"
	var summary string = "[ "
	for i := 0; i < 9; i++ {
		adult := 0
		if i < len(adults) {
			adult = adults[i]
		}
		juvenile := juveniles[i]

		summary += strconv.Itoa(adult+juvenile) + " | "
	}
	summary += " ] \n"
	fmt.Println(summary_head + summary)
}

func main() {
	var x *utils.InputReader = utils.NewInputReader(CHALLENGE)
	x.SplitStringToInts(",")
	x.ConvertToInts()

	var adults [GESTATION_RATE]int
	var newborns [NEWBORN]int

	for _, i := range x.IntData {
		adults[i] += 1
	}

	//TOTAL_ADULTS := utils.Sum_Array(adults[:])
	//TOTAL_NEWBORNS := utils.Sum_Array(newborns[:])
	//status(0, TOTAL_NEWBORNS, TOTAL_ADULTS+TOTAL_NEWBORNS)
	//fish_array(adults, newborns)

	// FOR EACH DAY IN THE PERIOD
	for day := 0; day < DAYS; day++ {
		// FIND THE NEWBORNS MATURING
		maturing := day % NEWBORN
		spawning := day % GESTATION_RATE
		newborns_maturing := newborns[maturing]
		adults_spawning := adults[spawning]

		adults[spawning] += newborns_maturing
		newborns[maturing] += adults_spawning

		TOTAL_ADULTS := utils.Sum_Array(adults[:])
		TOTAL_NEWBORNS := utils.Sum_Array(newborns[:])
		status(day+1, TOTAL_NEWBORNS, TOTAL_ADULTS+TOTAL_NEWBORNS)
		fish_array(adults, newborns)
	}

	// var fishpop []int
	// fishpop = append(fishpop, x.IntData...)

	// for day := 0; day < DAYS; day++ {
	// 	for i, fertility := range fishpop {
	// 		if fertility == 0 {
	// 			fishpop = append(fishpop, 8)
	// 			fishpop[i] = 6
	// 		} else {
	// 			fishpop[i]--
	// 		}
	// 	}
	// 	fmt.Printf("Fishes: %d \n", len(fishpop))
	// }

}
