package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/kevin-kho/aoc-utilities/common"
	"github.com/kevin-kho/aoc-utilities/models"
)

type Board struct {
	X      int
	Y      int
	Values map[models.Pos]byte
}

func CreateBoard(data []byte) Board {
	res := Board{
		Values: map[models.Pos]byte{},
	}

	arr := bytes.Split(data, []byte{'\n'})
	res.X = len(arr)
	res.Y = len(arr[0])

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

func main() {
	dataPartOne, err := common.ReadInput("inputPartOne.txt")
	if err != nil {
		log.Fatal(err)
	}
	dataPartOne = common.TrimNewLineSuffix(dataPartOne)

	boardPartOne := CreateBoard(dataPartOne)
	fmt.Println(boardPartOne)

}
