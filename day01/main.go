package main

import (
	"fmt"
	"log"
	"math"

	"github.com/kevin-kho/aoc-utilities/common"
)

func SolvePartOne(intArr []int) int {

	var res int

	for _, num := range intArr {
		sq := math.Sqrt(float64(num))
		sq = math.Ceil(sq)

		nearest := int(sq) * int(sq)

		res += (nearest - num)

	}

	return res

}

func main() {
	dataPartOne, err := common.ReadInput("inputPartOne.txt")
	if err != nil {
		log.Fatal(err)
	}
	dataPartOne = common.TrimNewLineSuffix(dataPartOne)

	intArr, err := common.ParseIntArray(dataPartOne, []byte{','})
	if err != nil {
		log.Fatal(err)
	}

	resPartOne := SolvePartOne(intArr)
	fmt.Println(resPartOne)

}
