package utils

import (
	"errors"
	"fmt"
)

type Grid struct {
	Data [][]string
	Rows, Cols int
}


type Point struct {
	X, Y int
}


func (g Grid) Print() {
	for _, row := range g.Data {
		fmt.Println(row)
	}
	fmt.Println("__________________________________")
}


// Suppose there must be a reflection !?
// isRelaxed is used for Part2
func (g Grid) FindMirrorLine(isRelaxed bool) (int, int, string) {

	mid1, mid2, err := g.findReflection(true, isRelaxed)
	if err == nil {
		return mid1, mid2, "Vertical"
	}
	mid1, mid2, _ = g.findReflection(false, isRelaxed)

	return mid1, mid2, "Horizontal"
}


func (g Grid) findReflection(isVertical, isRelaxed bool) (int, int, error) {

	end := 0
	if isVertical {
		end = g.Cols
	} else {
		end = g.Rows
	}

	// Try all possibilities
	for start := 0; start < end - 1; start++ {

		idxA := start
		idxB := start + 1
		
		_, _, err := g.expand(isVertical, isRelaxed, idxA, idxB, end)
		if err == nil {
			return idxA, idxB, nil
		}
	}

	err := errors.New(fmt.Sprintf("No Valid Reflection found - isVertical: %v", isVertical))
	return 0, 0, err
}


// Retruns the left limit and right limit that forms a valid expansion OR err
func (g Grid) expand(isVertical, isRelaxed bool, idxA, idxB, end int) (int, int, error){

	areSame := true
	currentDiff := 0
	forgivingTries := 0


	for idxA > -1 && idxB < end && areSame{

		arrA := []string {}
		arrB := []string {}
		if isVertical {
			arrA = g.getCol(idxA)
			arrB = g.getCol(idxB)
		} else {
			arrA = g.getRow(idxA)
			arrB = g.getRow(idxB)
		}

		if isRelaxed {
			currentDiff = arrMustDifferByOne(arrA, arrB)
			if currentDiff == 1 {
				forgivingTries++
			} else if currentDiff == 2 {
				areSame = false
			}
			if forgivingTries == 2 {
				areSame = false
			} 
		} else {
			areSame = compareArrs(arrA, arrB)
		}
		
		idxA--
		idxB++
	}

	if isRelaxed {
		if areSame && (idxA == -1 || idxB == end) && forgivingTries == 1 {
			return idxA + 1, idxB - 1, nil
		}
	} else {
		if areSame && (idxA == -1 || idxB == end) {
			return idxA + 1, idxB - 1, nil
		}
	}
	
	return 0, 0, errors.New("Cant Expand To Border.")
}
		

func (g Grid) getRow(idx int) []string {
	return g.Data[idx]
}


func (g Grid) getCol(idx int) []string {
	col := []string {}
	for _, row := range g.Data {
		col = append(col, row[idx])
	}
	return col
}


func compareArrs(r1 []string, r2 []string) bool {
	for idx := 0; idx < len(r1); idx++ {
		if r1[idx] != r2[idx] {
			return false
		}
	}
	return true
}


func arrMustDifferByOne(r1 []string, r2 []string) int {
	diff := 0
	for idx := 0; idx < len(r1); idx++ {
		if r1[idx] != r2[idx] {
			diff++
		}
		if diff == 2 {
			return diff
		}
	}
	return diff
}