package utils

import "strconv"

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

// convert 1/0 int to bool
func ItoB(i int) bool {
	if i == 0 {
		return false
	} else {
		return true
	}
}

// convert string of bits to binary int64
func SBinarytoI64(s string) int64 {
	if i, err := strconv.ParseInt(s, 2, 64); err != nil {
		return -1
	} else {
		return i
	}
}

func ArrBintoS(b []bool) string {
	var ret string
	for _, v := range b {
		if v {
			ret += "1"
		} else {
			ret += "0"
		}
	}
	return ret
}

func BitFlipper(s string) string {
	var ret string
	for _, i := range s {
		if i == '1' {
			ret += "0"
		} else if i == '0' {
			ret += "1"
		} else {
			ret += "X"
		}
	}
	return ret
}

func ArrStoI(s []string) []int {
	var ints []int
	for i := range s {
		j, _ := strconv.Atoi(s[i])
		ints = append(ints, j)
	}
	return ints
}

func MinMaxI(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

func LessByVal(i, j int) bool { return i < j }
