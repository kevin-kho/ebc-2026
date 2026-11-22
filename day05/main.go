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

func main() {
	dataPartOne, err := common.ReadInput("inputPartOne.txt")
	if err != nil {
		log.Fatal(err)
	}
	dataPartOne = common.TrimNewLineSuffix(dataPartOne)

	boardPartOne := CreateBoard(dataPartOne)

	resPartOne := SolvePartOne(boardPartOne)
	fmt.Println(resPartOne)

}
