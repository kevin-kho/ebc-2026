package main

import (
	"cmp"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/kevin-kho/aoc-utilities/common"
)

type Entry struct {
	Id    int
	Quest int
	Part  int
	Time  int
}

type Participant struct {
	Id    int
	Times []int
}

func (p Participant) TotalizeTime() int {
	var res int
	for _, t := range p.Times {
		res += t
	}

	return res
}

func GetEntries(data []byte) ([]Entry, error) {
	var res []Entry

	for entry := range strings.SplitSeq(string(data), "\n") {
		entryArr := strings.Split(entry, " ")
		id := entryArr[0]
		id = strings.TrimPrefix(id, "id=")

		quest := entryArr[1]
		quest = strings.TrimPrefix(quest, "quest=")

		part := entryArr[2]
		part = strings.TrimPrefix(part, "part=")

		time := entryArr[3]
		time = strings.TrimPrefix(time, "time=")

		idInt, err := strconv.Atoi(id)
		if err != nil {
			return res, err
		}

		questInt, err := strconv.Atoi(quest)
		if err != nil {
			return res, err
		}

		partInt, err := strconv.Atoi(part)
		if err != nil {
			return res, err
		}

		timeInt, err := strconv.Atoi(time)
		if err != nil {
			return res, err
		}

		res = append(res, Entry{
			Id:    idInt,
			Quest: questInt,
			Part:  partInt,
			Time:  timeInt,
		})

	}

	return res, nil

}

func SolvePartOne(entries []Entry) int {
	// Entries are all quest 1
	mp := make(map[int]Participant) // key: id, value: sum of times
	for _, e := range entries {
		if _, ok := mp[e.Id]; !ok {
			mp[e.Id] = Participant{Id: e.Id}
		}

		p := mp[e.Id]
		p.Times = append(p.Times, e.Time)
		mp[e.Id] = p

	}

	var participants []Participant
	for _, p := range mp {
		participants = append(participants, p)
	}

	slices.SortFunc(participants, func(a, b Participant) int {
		return cmp.Or(
			cmp.Compare(len(b.Times), len(a.Times)),
			cmp.Compare(a.TotalizeTime(), b.TotalizeTime()),
		)
		// return cmp.Compare(a.TotalizeTime(), b.TotalizeTime())
	})

	var res int
	i := 1
	for _, p := range participants {
		res += p.TotalizeTime() * i
		i++
	}

	return res

}

func main() {
	dataPartOne, err := common.ReadInput("inputPartOne.txt")
	if err != nil {
		log.Fatal(err)
	}
	dataPartOne = common.TrimNewLineSuffix(dataPartOne)

	entriesPartOne, err := GetEntries(dataPartOne)
	if err != nil {
		log.Fatal(err)
	}

	resPartOne := SolvePartOne(entriesPartOne)
	fmt.Println(resPartOne)

	dataPartTwo, err := common.ReadInput("inputPartTwo.txt")
	if err != nil {
		log.Fatal(err)
	}
	dataPartTwo = common.TrimNewLineSuffix(dataPartTwo)

	entriesPartTwo, err := GetEntries(dataPartTwo)
	if err != nil {
		log.Fatal(err)
	}

	resPartTwo := SolvePartOne(entriesPartTwo)
	fmt.Println(resPartTwo)

}
