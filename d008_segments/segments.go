package main

import (
	"advent_of_code/utils"
	"strings"
)

const CHALLENGE string = "../input/segments.input"
const TEST string = "../input/segments.test"

type signal_outputs struct {
	signal []string
	output []string
}

type digit_complete func() bool

type display struct {
	top_h rune // 0
	mid_h rune // 1
	bot_h rune // 2

	top_l rune // 3
	top_r rune // 4

	bot_l rune // 5
	bot_r rune // 6

	one   digit_complete
	two   digit_complete
	three digit_complete
	four  digit_complete
	five  digit_complete
	six   digit_complete
	seven digit_complete
	eight digit_complete
	nine  digit_complete
	zero  digit_complete
}

func new_display() *display {
	var d *display = new(display)
	d.one = func() bool { return segments_solved([]rune{d.top_r, d.bot_r}) }
	d.two = func() bool { return segments_solved([]rune{d.top_h, d.mid_h, d.bot_h, d.top_r, d.bot_l}) }
	d.three = func() bool { return segments_solved([]rune{d.top_h, d.mid_h, d.bot_h, d.top_r, d.bot_r}) }
	d.four = func() bool { return segments_solved([]rune{d.mid_h, d.top_l, d.top_r, d.bot_r}) }
	d.five = func() bool { return segments_solved([]rune{d.top_h, d.mid_h, d.bot_h, d.top_l, d.bot_r})
	d.six = func() bool { return segments_solved([]rune{d.top_h, d.mid_h, d.bot_h, d.top_l, d.bot_l, d.bot_r})
	d.seven = func() bool { return segments_solved([]rune{d.top_h, d.top_r, d.bot_r })
	d.eight = func() bool { return segments_solved([]rune{d.top_h, d.mid_h, d.bot_h, d.top_l, d.top_r, d.bot_l, d.bot_r})
	d.nine = func() bool { return segments_solved([]rune{d.top_h, d.mid_h, d.top_l, d.top_r, d.bot_r}) }
	d.zero = func() bool { return segments_solved([]rune{d.top_h, d.bot_h, d.top_l, d.top_r, d.bot_l, d.bot_r})

	//d.two = func() bool { return d.top_h != 0 &&}
}

func segments_solved(r []rune) bool {
	for _, v := range r {
		if v == 0 {
			return false
		}
	}
	return true
}

func map_signal_to_display(signals []string) {
	var d display
	var solved bool
	for {
		for i := 0; i < len(signals); i++ {

		}
		if solved {
			break
		}
	}

}

func new_signal_ouput(s []string, o []string) *signal_outputs {
	var so *signal_outputs = new(signal_outputs)
	so.signal = s
	so.output = o
	return so
}

var signals []signal_outputs

func main() {
	var ir *utils.InputReader = utils.NewInputReader(TEST)
	for _, v := range ir.StrData {
		values := strings.Split(v, " ")
		signals = append(signals, *new_signal_ouput(values[:10], values[11:15]))
	}

}
