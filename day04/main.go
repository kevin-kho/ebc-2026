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
	Id        int
	TotalTime int
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
	mp := make(map[int]int) // key: id, value: sum of times
	for _, e := range entries {
		mp[e.Id] += e.Time
	}

	var participants []Participant
	for id, time := range mp {
		participants = append(participants, Participant{
			Id:        id,
			TotalTime: time,
		})
	}

	slices.SortFunc(participants, func(a, b Participant) int {
		return cmp.Compare(a.TotalTime, b.TotalTime)
	})

	var res int
	i := 1
	for _, p := range participants {
		res += p.TotalTime * i
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

}
