package main

import (
	"advent_of_code/utils"
	"fmt"
	"strings"
)

const INPUT_BOARDS string = "../input/giant_squid_boards.input"
const TEST_BOARDS string = "../input/giant_squid_boards.test"

const INPUT_NUMBERS string = "../input/giant_squid_order.input"
const TEST_NUMBERS string = "../input/giant_squid_order.test"

type bingo_board struct {
	columns [5][5]int
	rows    [5][5]int

	hitC [5][5]int
	hitR [5][5]int

	bingo       bool
	last_number int

	raw_input []string
}

func (b *bingo_board) construct_board() {
	for i := 0; i < len(b.raw_input); i++ {
		s_raw := strings.Fields(b.raw_input[i])
		copy(b.rows[i][:], utils.ArrStoI(s_raw)[:5])
	}

	for col := 0; col < 5; col++ {
		for row := 0; row < 5; row++ {
			b.columns[col][row] = b.rows[row][col]
		}
	}
}

func (b *bingo_board) mark_board(number int) {
	for row := 0; row < 5; row++ {
		for ele := 0; ele < 5; ele++ {
			if b.rows[row][ele] == number {
				b.hitR[row][ele] = 1
			}
		}
	}

	for col := 0; col < 5; col++ {
		for ele := 0; ele < 5; ele++ {
			if b.columns[col][ele] == number {
				b.hitC[col][ele] = 1
			}
		}
	}

	b.bingo = b.check_win()
}

func (b *bingo_board) check_win() bool {
	for col := range b.hitC {
		if utils.Sum_Array(b.hitC[col][:]) == 5 {
			return true
		}
	}

	for row := range b.hitR {
		if utils.Sum_Array(b.hitR[row][:]) == 5 {
			return true
		}
	}

	return false
}

func (b *bingo_board) solve() int {
	var sum int
	for col := 0; col < 5; col++ {
		for ele := 0; ele < 5; ele++ {
			if b.hitC[col][ele] == 0 {
				sum += b.columns[col][ele]
			}
		}
	}
	return sum * b.last_number
}

func NewBingoBoard(data []string) *bingo_board {
	var b = new(bingo_board)
	b.raw_input = data
	return b
}

type bingo_game struct {
	calling_order []int
	boards        []*bingo_board
}

func (g *bingo_game) play_to_win() *bingo_board {
	for _, c := range g.calling_order {
		for j, b := range g.boards {
			b.mark_board(c)
			if b.bingo {
				b.last_number = c
				fmt.Printf("WE HAVE A WINNER - LUCKY NUMBER: %d \n", j)
				return b
			}
		}
	}
	var failure *bingo_board = new(bingo_board)
	return failure
}

func (g *bingo_game) play_to_lose() *bingo_board {
	for _, c := range g.calling_order {
		if len(g.boards) == 1 {
			g.boards[0].last_number = c
			fmt.Printf("WE HAVE A LOSER - LUCKY NUMBER: %d \n", 0)
			return g.boards[0]
		}

		board_count := len(g.boards)
		for b := 0; b < board_count; b++ {
			g.boards[b].mark_board(c)
			if g.boards[b].bingo {
				g.throw_out_board(b)
				board_count--
				b--
			}
		}
	}
	var failure *bingo_board = new(bingo_board)
	return failure
}

func (g *bingo_game) throw_out_board(i int) {
	g.boards[i] = g.boards[len(g.boards)-1]
	g.boards = g.boards[:len(g.boards)-1]
}

func main() {
	var bingo *bingo_game = new(bingo_game)
	var ir *utils.InputReader = utils.NewInputReader(TEST_NUMBERS)
	ir.SplitStringToInts(",")
	bingo.calling_order = ir.IntData

	var board_ir *utils.InputReader = utils.NewInputReader(TEST_BOARDS)
	//board_count := (board_ir.InputCount + 1) / 6

	for i := 0; i < board_ir.InputCount; i += 6 {
		board := NewBingoBoard(board_ir.StrData[i : i+5])
		board.construct_board()
		bingo.boards = append(bingo.boards, board)
	}

	winner := bingo.play_to_win()
	fmt.Printf("SOLUTION: %d \n", winner.solve())

	loser := bingo.play_to_lose()
	fmt.Printf("SOLUTION: %d \n", loser.solve())

}
