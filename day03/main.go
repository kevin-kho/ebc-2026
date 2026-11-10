package main

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/kevin-kho/aoc-utilities/common"
)

type Iso struct {
	Standard int
	Vote     int
}

func (i Iso) ConvertToDecimal() int {

	var res int
	var pow int

	for i.Standard > 0 && i.Vote > 0 {
		s := i.Standard % 10
		i.Standard /= 10

		v := i.Vote % 10
		i.Vote /= 10

		s = int(math.Pow(float64(s), float64(pow)))

		res += s * v

		pow++
	}
	return res
}

func CreateIsos(data []byte) ([]Iso, error) {
	var res []Iso
	for entry := range bytes.SplitSeq(data, []byte{'\n'}) {
		entryStr := strings.TrimPrefix(string(entry), "ISO-")

		std, val, _ := strings.Cut(entryStr, " ")

		stdInt, err := strconv.Atoi(std)
		if err != nil {
			return res, err
		}

		valInt, err := strconv.Atoi(val)
		if err != nil {
			return res, err
		}

		res = append(res, Iso{
			Standard: stdInt,
			Vote:     valInt,
		})

	}

	return res, nil

}

func SolvePartOne(isos []Iso) int {
	var res int
	for _, i := range isos {
		res += i.ConvertToDecimal()
	}
	return res
}

func main() {
	dataPartOne, err := common.ReadInput("inputPartOne.txt")
	if err != nil {
		log.Fatal(err)
	}
	dataPartOne = common.TrimNewLineSuffix(dataPartOne)

	isos, err := CreateIsos(dataPartOne)
	if err != nil {
		log.Fatal(err)
	}

	resPartOne := SolvePartOne(isos)
	fmt.Println(resPartOne)

}
