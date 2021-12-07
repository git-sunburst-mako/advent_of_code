package main

import (
	"advent_of_code/utils"
	"fmt"
	"strings"
)

const CHALLENGE string = "../input/hydrothermal.input"
const TEST string = "../input/hydrothermal.test"

type coords struct {
	x_val int
	y_val int
}

func new_coords_xy(x, y int) coords {
	c := coords{
		x_val: x,
		y_val: y,
	}
	return c
}

func new_coords_arr(a []int) coords {
	c := coords{
		x_val: a[0],
		y_val: a[1],
	}
	return c
}

type orientation int

const ( // iota is reset to 0
	HORIZONTAL orientation = iota // c0 == 0
	VERTICAL               = iota // c1 == 1
	DIAGONAL               = iota // c2 == 2
)

type vent struct {
	start coords
	end   coords

	path             []coords
	vent_orientation orientation
}

func new_vent(start, end coords) vent {
	var v vent
	v.start = start
	v.end = end

	v.generate_vent()
	v.generate_path()

	return v
}

func (v *vent) generate_vent() {
	if v.start.x_val == v.end.x_val {
		v.vent_orientation = VERTICAL
	} else if v.start.y_val == v.end.y_val {
		v.vent_orientation = HORIZONTAL
	} else {
		v.vent_orientation = DIAGONAL
	}
}

func (v *vent) generate_path() {
	switch v.vent_orientation {
	case VERTICAL:
		y1, y2 := utils.MinMaxI(v.start.y_val, v.end.y_val)
		for s := y1; s <= y2; s++ {
			v.path = append(v.path, new_coords_xy(v.start.x_val, s))
		}
	case HORIZONTAL:
		x1, x2 := utils.MinMaxI(v.start.x_val, v.end.x_val)
		for s := x1; s <= x2; s++ {
			v.path = append(v.path, new_coords_xy(s, v.start.y_val))
		}
	case DIAGONAL:
		var x_increase, y_increase int = 1, 1
		if v.start.x_val > v.end.x_val {
			x_increase = -1
		}
		if v.start.y_val > v.end.y_val {
			y_increase = -1
		}
		var next coords = v.start
		for next != v.end {
			v.path = append(v.path, next)
			next.x_val += x_increase
			next.y_val += y_increase
		}
		v.path = append(v.path, v.end)
	}
}

type seafloor struct {
	vents      []vent
	topography map[coords]int
}

func (s *seafloor) analyse_vents() {
	s.topography = make(map[coords]int)
	for _, v := range s.vents {
		for _, p := range v.path {
			s.topography[p] += 1
		}
	}
}

func (s *seafloor) count_intersections() int {
	var x int
	for _, v := range s.topography {
		if v > 1 {
			x++
		}
	}
	return x
}

func process_coords_input(line string) (coords, coords) {
	raw := strings.Split(line, " -> ")
	j := strings.Split(raw[0], ",")
	k := strings.Split(raw[1], ",")

	start := utils.ArrStoI(j)
	end := utils.ArrStoI(k)

	return new_coords_arr(start), new_coords_arr(end)
}

func main() {
	var x *utils.InputReader = utils.NewInputReader(CHALLENGE)
	var s seafloor

	for _, r := range x.StrData {
		v := new_vent(process_coords_input(r))
		s.vents = append(s.vents, v)
	}

	s.analyse_vents() // build topography
	pt1 := s.count_intersections()

	fmt.Printf("Intersections: %d", pt1)

}
