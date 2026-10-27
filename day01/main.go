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

func SolvePartTwo(intArr []int) int {
	// Follows the pattern: (n * n) - ((n - 2) * (n - 2))
	// For n >= 3
	// Elements in this sequence will always be divisible by 4
	var res int

	for _, num := range intArr {
		// case: make the smallest frame
		if num < 8 {
			res += (8 - num)
			continue
		}

		rem := num % 4
		if rem == 0 {
			continue
		}

		res += (4 - rem)

	}

	return res
}

func main() {
	dataPartOne, err := common.ReadInput("inputPartOne.txt")
	if err != nil {
		log.Fatal(err)
	}
	dataPartOne = common.TrimNewLineSuffix(dataPartOne)

	intArrPartOne, err := common.ParseIntArray(dataPartOne, []byte{','})
	if err != nil {
		log.Fatal(err)
	}

	resPartOne := SolvePartOne(intArrPartOne)
	fmt.Println(resPartOne)

	dataPartTwo, err := common.ReadInput("inputPartTwo.txt")
	if err != nil {
		log.Fatal(err)
	}
	dataPartTwo = common.TrimNewLineSuffix(dataPartTwo)

	intArrPartTwo, err := common.ParseIntArray(dataPartTwo, []byte{','})
	if err != nil {
		log.Fatal(err)
	}

	resPartTwo := SolvePartTwo(intArrPartTwo)
	fmt.Println(resPartTwo)

}
