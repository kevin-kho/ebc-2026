package main

import (
	"bytes"
	"fmt"
	"log"
	"maps"

	"github.com/kevin-kho/aoc-utilities/common"
	"github.com/kevin-kho/aoc-utilities/models"
)

type Board struct {
	X      int
	Y      int
	Values map[models.Pos]byte
}

func (b Board) PrintGrid() {
	var grid [][]byte
	for range b.Y {
		row := make([]byte, b.X)
		grid = append(grid, row)
	}

	for pos, val := range b.Values {
		grid[pos.Y][pos.X] = val
	}

	for _, row := range grid {
		fmt.Println(string(row))
	}

}

func CreateBoard(data []byte) Board {
	res := Board{
		Values: map[models.Pos]byte{},
	}

	arr := bytes.Split(data, []byte{'\n'})
	res.Y = len(arr)
	res.X = len(arr[0])

	for y, row := range bytes.Split(data, []byte{'\n'}) {
		for x, val := range row {
			pos := models.Pos{
				X: x,
				Y: y,
			}
			res.Values[pos] = val
		}
	}

	return res
}

func GetAllDir() []models.Pos {
	res := models.GetDeltas()
	res = append(res, models.Pos{X: -1, Y: -1})
	res = append(res, models.Pos{X: -1, Y: 1})
	res = append(res, models.Pos{X: 1, Y: 1})
	res = append(res, models.Pos{X: 1, Y: -1})

	return res

}

func GetTreeDirs() []models.Pos {
	var res []models.Pos
	for x := -2; x < 3; x++ {
		for y := -2; y < 3; y++ {
			if x == 0 && y == 0 {
				continue
			}
			if common.IntAbs(x) == 2 && common.IntAbs(y) == 2 {
				continue
			}
			res = append(res, models.Pos{X: x, Y: y})
		}
	}

	return res
}

// Might change it to receiver function
func MapOutLandAndWater(board Board) Board {

	start := models.Pos{X: board.X / 2, Y: board.Y / 2}
	board.Values[start] = '#'

	// set of all land
	// will use to check if any land changes
	mp := map[models.Pos]bool{start: true}

	var queue []models.Pos // should only contain non #

	for true {
		newMp := maps.Clone(mp)

		// populate queue with 1 and 0s
		for pos := range mp {
			for _, d := range GetAllDir() {
				newX := pos.X + d.X
				newY := pos.Y + d.Y

				// skip out of bounds
				if !(0 <= newX && newX < board.X) || !(0 <= newY && newY < board.Y) {
					continue
				}
				adjPos := models.Pos{X: newX, Y: newY}

				// skip adj #
				if board.Values[adjPos] == '#' {
					continue
				}
				queue = append(queue, adjPos)
			}
		}

		// BFS the queue
		for range len(queue) {
			pos := queue[0]
			queue = queue[1:]
			var ct int

			for _, d := range GetAllDir() {
				newX := pos.X + d.X
				newY := pos.Y + d.Y

				// skip out of bounds
				if !(0 <= newX && newX < board.X) || !(0 <= newY && newY < board.Y) {
					continue
				}

				newPos := models.Pos{X: newX, Y: newY}

				if board.Values[newPos] == '#' || board.Values[newPos] == '1' {
					ct++
				}
			}

			if ct >= 5 {
				newMp[pos] = true
				board.Values[pos] = '#'
			}
		}

		// No new land
		if maps.Equal(mp, newMp) {
			break
		}

		mp = newMp

	}

	// Turn the rest into water
	for pos, val := range board.Values {
		if val == '#' {
			continue
		}
		board.Values[pos] = '.'
	}

	return board

}

func MapOutShore(board Board) Board {
	dirs := GetAllDir()

	// set initial shore
	for pos, val := range board.Values {
		if val != '#' {
			continue
		}
		var ct int
		for _, d := range dirs {
			newX := pos.X + d.X
			newY := pos.Y + d.Y
			if !(0 <= newX && newX < board.X) || !(0 <= newY && newY < board.Y) || board.Values[models.Pos{X: newX, Y: newY}] == '.' {
				ct++
			}
		}
		if ct > 0 {
			board.Values[pos] = '='
		}
	}

	for true {
		newMp := maps.Clone(board.Values)

		for pos, val := range board.Values {
			if val != '#' {
				continue
			}

			var ct int
			for _, d := range dirs {
				newX := pos.X + d.X
				newY := pos.Y + d.Y
				if !(0 <= newX && newX < board.X) || !(0 <= newY && newY < board.Y) {
					continue
				}
				if board.Values[models.Pos{X: newX, Y: newY}] == '=' {
					ct++
				}
			}

			if ct >= 5 {
				newMp[pos] = '='
			}
		}

		if maps.Equal(board.Values, newMp) {
			break
		}
		board.Values = newMp
	}

	return board
}

func MapOutTrees(board Board) Board {
	// dirs := GetAllDir()

	// set initial forest
	for pos, val := range board.Values {
		if val != '#' {
			continue
		}

		valid := true
		for _, d := range GetTreeDirs() {
			newX := pos.X + d.X
			newY := pos.Y + d.Y

			if !(0 <= newX && newX < board.X) || !(0 <= newY && newY < board.Y) {
				valid = false
				continue
			}
			nei := board.Values[models.Pos{X: newX, Y: newY}]
			if nei == '.' || nei == '=' {
				valid = false
			}
		}
		if valid {
			board.Values[pos] = '@'
		}

	}

	// grow forest
	for range 2 {
		newMp := maps.Clone(board.Values)
		for pos, val := range board.Values {
			if val != '#' {
				continue
			}
			var ct int
			var sand bool

			for _, d := range GetAllDir() {
				newX := pos.X + d.X
				newY := pos.Y + d.Y

				if !(0 <= newX && newX < board.X) || !(0 <= newY && newY < board.Y) {
					continue
				}
				nei := board.Values[models.Pos{X: newX, Y: newY}]
				switch nei {
				case '=':
					sand = true
				case '@':
					ct++
				}
			}
			if ct >= 4 && !sand {
				newMp[pos] = '@'
			}
		}
		board.Values = newMp
	}

	return board
}

func SolvePartOne(board Board) int {

	board = MapOutLandAndWater(board)
	var ct int
	for _, val := range board.Values {
		if val == '#' {
			ct++
		}
	}

	board.PrintGrid()

	return ct

}

func SolvePartTwo(board Board) int {
	board = MapOutLandAndWater(board)
	board = MapOutShore(board)

	var ct int
	for _, val := range board.Values {
		if val == '=' {
			ct++
		}
	}
	board.PrintGrid()
	return ct

}

func SolvePartThree(board Board) int {
	board = MapOutLandAndWater(board)
	board = MapOutShore(board)
	board = MapOutTrees(board)

	board.PrintGrid()
	var ct int
	for _, val := range board.Values {
		if val == '@' {
			ct++
		}
	}

	return ct
}

func main() {
	dataPartOne, err := common.ReadInput("inputPartOne.txt")
	if err != nil {
		log.Fatal(err)
	}
	dataPartOne = common.TrimNewLineSuffix(dataPartOne)

	boardPartOne := CreateBoard(dataPartOne)

	resPartOne := SolvePartOne(boardPartOne)
	fmt.Println(resPartOne)

	dataPartTwo, err := common.ReadInput("inputPartTwo.txt")
	if err != nil {
		log.Fatal(err)
	}
	dataPartTwo = common.TrimNewLineSuffix(dataPartTwo)
	boardPartTwo := CreateBoard(dataPartTwo)

	resPartTwo := SolvePartTwo(boardPartTwo)
	fmt.Println(resPartTwo)

	dataPartThree, err := common.ReadInput("inputPartThree.txt")
	if err != nil {
		log.Fatal(err)
	}
	dataPartThree = common.TrimNewLineSuffix(dataPartThree)
	boardPartThree := CreateBoard(dataPartThree)

	resPartThree := SolvePartThree(boardPartThree)
	fmt.Println(resPartThree)

}
