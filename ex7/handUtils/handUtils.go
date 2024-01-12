package handutils

import (
	"fmt"
	"sort"
)


type Hand struct {
	Cards string
	Bid   int
}


func (h Hand) String() string {
	return fmt.Sprintf("{CARDS <%s> - BID <%d>}", h.Cards, h.Bid)
}


// true if card1 < card2
func compareCards(card1 string, card2 string, isJWild bool) bool {

	values := cardValues
	if isJWild {
		values = cardValuesJMod
	}

	for idx := 0; idx < len(card1); idx++ {
		
		card1Val := values[string(card1[idx])]
		card2Val := values[string(card2[idx])]

		if card1Val == card2Val {
			continue
		}
		if card1Val > card2Val {
			return false
		}
		if card1Val < card2Val {
			return true
		}
	}
	return false
}


func (h Hand) GetStrength(isJWild bool) int{
	occurences := [] int {}
	if isJWild {
		occurences = h.getCardsCountsWithJMOD()
	} else {
		occurences = h.getCardsCounts()
	}
	return getHandStrength(occurences)
}


func getHandStrength(occurences []int) int {
	if compareArrays(occurences, FiveOfAKind){
		return 6
	}else if  compareArrays(occurences, FourOfAKind){
		return 5
	}else if  compareArrays(occurences, FullHouse){
		return 4
	}else if  compareArrays(occurences, ThreeOfAKind){
		return 3
	}else if  compareArrays(occurences, TwoPair){
		return 2
	}else if  compareArrays(occurences, OnePair){
		return 1
	}else {
		return 0
	}	
}


func compareArrays(arr1 []int, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for idx := 0; idx < len(arr1); idx++ {
		if arr1[idx] != arr2[idx] {
			return false
		}
	}
	return true
}


func (h Hand) getCardsCounts() []int {

	counts := map[string] int {}

	for _, char := range h.Cards {
		counts[string(char)]++
	}

	var occurences []int

	for _, v := range counts {
		occurences = append(occurences, v)
	}

	sort.Ints(occurences)

	return occurences
}


func (h Hand) getCardsCountsWithJMOD() []int {

	counts := map[string] int {}
	jCounter := 0

	for _, char := range h.Cards {
		if string(char) == "J" {
			jCounter++
		} else {
			counts[string(char)]++
		}
	}

	if jCounter == 5 {
		return []int { 5 }
	}

	var occurences []int

	for _, v := range counts {
		occurences = append(occurences, v)
	}

	sort.Ints(occurences)

	for idx := len(occurences) - 1; idx > -1 ; idx-- {

		for occurences[idx] < 5 && jCounter > 0 {
			occurences[idx]++
			jCounter--
		}
	}

	return occurences
} 


var FiveOfAKind = []int { 5 }
var FourOfAKind = []int { 1, 4 }
var FullHouse = []int { 2, 3 }
var ThreeOfAKind = []int { 1, 1, 3 }
var TwoPair = []int { 1, 2, 2 }
var OnePair = []int { 1, 1, 1, 2 }
var HighCard = []int { 1, 1, 1, 1, 1 }


var cardValues = map[string]int{
	"A": 13,
	"K": 12,
	"Q": 11,
	"J": 10,
	"T": 9,
	"9": 8,
	"8": 7,
	"7": 6,
	"6": 5,
	"5": 4,
	"4": 3,
	"3": 2,
	"2": 1,
}

var cardValuesJMod = map[string]int{
	"A": 13,
	"K": 12,
	"Q": 11,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
	"J": 1,  
}


type ByCards struct {
	Hands   []Hand
	IsJWild bool
}


func (a ByCards) Len() int { 
	return len(a.Hands) 
}


func (a ByCards) Swap(i, j int) { 
	a.Hands[i], a.Hands[j] = a.Hands[j], a.Hands[i] 
}


func (a ByCards) Less(i, j int) bool { 
	return compareCards(a.Hands[i].Cards, a.Hands[j].Cards, a.IsJWild) 
}
