package recordUtils

import "fmt"

type Record struct {
	Data   []string
	Groups []int
}



func (r Record) Print() {
	fmt.Printf("DATA : %v | GROUPS : %v\n", r.Data, r.Groups)
}


func (r *Record) Unfold() {

	newData := []string{}
	newGroups := []int{}

	for i := 0; i < 4; i++ {
		newData = append(newData, r.Data...)
		newData = append(newData, "?")

		newGroups = append(newGroups, r.Groups...)
	}

	newData = append(newData, r.Data...)
	newGroups = append(newGroups, r.Groups...)

	r.Data = newData
	r.Groups = newGroups
}

func (r Record) SolveRecord(dataIdx int, groupIdx int, groupCount int, memo map[[3]int]int) int {
	if val, ok := memo[[3]int{dataIdx, groupIdx, groupCount}]; ok {
		return val
	}

	if dataIdx == len(r.Data) {
		if groupIdx == len(r.Groups) && groupCount == 0 {
			return 1
		}

		if groupIdx == len(r.Groups)-1 && r.Groups[groupIdx] == groupCount {
			return 1
		}

		return 0
	}

	ways := 0
	el := r.Data[dataIdx]

	if el == "." {

		if groupCount == 0 {
			ways += r.SolveRecord(dataIdx+1, groupIdx, groupCount, memo)
		} else if groupIdx < len(r.Groups) && r.Groups[groupIdx] == groupCount {
			ways += r.SolveRecord(dataIdx+1, groupIdx+1, 0, memo)
		}

	} else if el == "#" {

		ways += r.SolveRecord(dataIdx+1, groupIdx, groupCount+1, memo)

	} else if el == "?" {

		if groupCount == 0 {
			ways += r.SolveRecord(dataIdx+1, groupIdx, groupCount, memo)
		} else if groupIdx < len(r.Groups) && r.Groups[groupIdx] == groupCount {
			ways += r.SolveRecord(dataIdx+1, groupIdx+1, 0, memo)
		}

		ways += r.SolveRecord(dataIdx+1, groupIdx, groupCount+1, memo)

	}

	memo[[3]int{dataIdx, groupIdx, groupCount}] = ways
	return ways
}
