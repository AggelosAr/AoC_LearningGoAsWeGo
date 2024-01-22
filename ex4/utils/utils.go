package utils

import (
	"math"
	"strconv"
	"strings"
)

type Lottery struct {
	Winning, Current map[int]struct{}  
	Matches          int
}

func NewLottery(t string) *Lottery {

	parts := strings.Split(strings.Split(t, ":")[1], "|")

	winning := parsePart(parts[0])
	current := parsePart(parts[1])
	matches := getIntersection(winning, current)

	return &Lottery{
		Winning: winning,
		Current: current,
		Matches: matches,
	}
}


func (l Lottery) ComputeScore() int{
	// each match is worth 1, then it doubles for each match 
	return int(math.Pow(float64(2), float64(l.Matches - 1)))
}


func getIntersection(winning, current map[int]struct{}  ) int {
	matches := 0

	for k := range winning {
		_, exists := current[k]
		if exists {
			matches++
		}
	}

	return matches
}



func parsePart(t string) map[int]struct{} {

	f := make(map[int]struct{})
	for _, numberStr := range strings.Fields(t) {
		number, err := strconv.Atoi(numberStr)
		if err != nil {
			continue
		}
		f[number] = struct{}{}
	}
	return f

}
