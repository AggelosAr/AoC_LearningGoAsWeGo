package main

import (
	"bufio"
	"fmt"
	"math"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"os"
)


func main() {

	rawData := getData()
	seeds, maps := getFormattedData(rawData)

	// 0 s
	start := time.Now()
	part1Solution := solverPart1(seeds, maps)
	elapsed := time.Since(start)
	fmt.Printf("PART1 took %s\n", elapsed)
	fmt.Println("The result is PART1 : ", part1Solution)
	fmt.Println("__________________________")

	/*
	// 8 s
	start = time.Now()
	part1SolutionReverse := solveInReverse(seeds, maps)
	elapsed = time.Since(start)
	fmt.Printf("PART1 REVERSE took %s\n", elapsed)
	fmt.Println("The result is PART1 : ", part1SolutionReverse)
	fmt.Println("__________________________")

	// 4.20 mins
	start = time.Now()
	part2Solution := solverPart2(seeds, maps)
	elapsed = time.Since(start)
	fmt.Printf("Brute Force took %s\n", elapsed)
	fmt.Println("The result is PART2 : ", part2Solution)
	fmt.Println("__________________________")

	// 1.17 mins
	start = time.Now()
	part2Solution = solverPart2Routines(seeds, maps)
	elapsed = time.Since(start)
	fmt.Printf("Go routines took %s\n", elapsed)
	fmt.Println("The result is PART2 : ", part2Solution)
	fmt.Println("__________________________")
	*/

	//155.6069ms
	start = time.Now()
	part2Solution := solveInReversePart2(seeds, maps)
	elapsed = time.Since(start)
	fmt.Printf("Reverse Part2 took %s\n", elapsed)
	fmt.Println("The result is PART2 : ", part2Solution)
	fmt.Println("__________________________")

	if 88151870 == part1Solution &&  2008785 == part2Solution {
		fmt.Println(" N I C E ")
	} else {
		fmt.Println(" B A D ")
	}
}


func getDataPath() string {
	rootPath, _ := os.Getwd()
	return filepath.Join(rootPath, "data", "input.txt")
}


func getData() []string {

	file, err := os.Open(getDataPath())
	if err != nil {
		fmt.Println("Error opening file : ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	
	var data []string
	part := ""

	for scanner.Scan() {
		
		line := scanner.Text()

		if line == "" {
			data = append(data, part)
			part = ""
		}else {
			part += line + " - "
		}
	}
	data = append(data, part)
	return data
}



func extractNumbers(text string) []int { 
	parts := strings.Split(text, " ")

	numbers := []int {}

	for _, part := range parts {

		number, err := strconv.Atoi(part)
		if err != nil {
			continue
		}
		numbers = append(numbers, number)
	}
	return numbers
}

func extractTriplets(text string) [][]int {

	// here is all the numbers
	rightPart := strings.Split(text, ":")[1]

	// we formated it with "-" among all pairs of triplets 
	// in case of seeds there is a single line so w/e e.g. the first line
	triplets := strings.Split(rightPart, " - ")

	extractedTriplets := [][]int {}

	for _, triplet := range triplets {
			
		numbers := extractNumbers(triplet)

		// formatting issues since we appened " - "
		if len(numbers) > 0 {
			extractedTriplets = append(extractedTriplets, numbers)
		}
	}
	return extractedTriplets	
}


func getFormattedData(rawData []string) ([] int, [][][]int) {
	seeds := rawData[0]
	maps := [][][]int {}

	for i := 1; i < len(rawData); i++{
		maps = append(maps, extractTriplets(rawData[i]))
	}

	return extractNumbers(seeds), maps
}


func solverPart1(seeds []int, maps[][][]int) int {
	
	minimumLocation := math.MaxInt64

	for _, seed := range seeds {
		minimumLocation = min(minimumLocation, iterateMaps(seed, maps))
	}
	return minimumLocation
}


func iterateMaps(seed int, maps [][][]int) int {

	searching := seed
	
	for _, mapper := range maps {

		for _, ranger := range mapper {
			source := ranger[1]
			destination := ranger[0]
			step := ranger[2]
			
			if source <= searching && searching < source + step{

				offset := searching - source
				searching = destination + offset
				break
			}
		}
	}

	return searching
}


func solverPart2(seeds []int, maps[][][]int) int {
	
	minimumLocation := math.MaxInt64

	for i := 0 ; i < len(seeds)-1; i++ {
		
		seed := seeds[i]
		seedRange := seeds[i + 1]
		i++
		for j := 0; j < seedRange; j++ {
			minimumLocation = min(minimumLocation, iterateMaps(seed + j, maps))
		}
	}

	return minimumLocation
}


func solverPart2Routines(seeds []int, maps [][][]int) int {
	minimumLocation := math.MaxInt64

	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < len(seeds)-1; i++ {
		seed := seeds[i]
		seedRange := seeds[i+1]
		i++

		wg.Add(1)
		go func(seed, seedRange int) {
			defer wg.Done()

			localMinimum := math.MaxInt64
			for j := 0; j < seedRange; j++ {
				
				value := iterateMaps(seed+j, maps)
				localMinimum = min(localMinimum, value)
			}

			mu.Lock()
			minimumLocation = min(minimumLocation, localMinimum)
			mu.Unlock()
		}(seed, seedRange)
	}

	wg.Wait()
	return minimumLocation
}

// Instead we should go in reverse!!
// we should start from smallest last map values and try to find a seed that 
// exists by going in reverse

func solveInReverse(seeds []int, maps [][][]int) int {

	// make seeds a set
	seedsSet := map[int]struct{}{}

	for _, seed := range seeds {
		seedsSet[seed] = struct{}{}
	}
	location := 0
	for true {
		
		seedFound := iterateMapsInRevrse(location, maps)
		_, exists := seedsSet[seedFound]
		//fmt.Printf("Location <%d> - Seed <%d>\n", location, seedFound)
		if exists {
			return location
		}
		
		location++
	} 

	fmt.Println("Error -- this should never happen -- assuming valid data")
	return 0 // should never happen
}

func iterateMapsInRevrse(location int, maps [][][]int) int {

	searching := location
	for i := len(maps) - 1; i > -1; i-- {
		
		for _, ranger := range  maps[i] {
			source := ranger[0]
			destination := ranger[1]
			step := ranger[2]

			if source <= searching && searching < source + step{

				offset := searching - source
				searching = destination + offset
				break
			}
		}
	}
	return searching
}


func solveInReversePart2(seeds []int, maps [][][]int) int {

	// create a set of ranges for the seeds
	seedRanges := [][2] int {}

	for i := 0 ; i < len(seeds)-1; i++ {
		
		start := seeds[i]
		end := seeds[i + 1] + start

		currentRange := [2]int {start, end}
		seedRanges = append(seedRanges, currentRange)
		i++
	}


	location := 0
	for true {
		
		seedFound := iterateMapsInRevrse(location, maps)
		
		// check if the seed found is in the range set
		for _, currentRange := range seedRanges {
			if currentRange[0] <= seedFound && seedFound <= currentRange[1] {
				return location
			}
		}
		location++
	} 

	fmt.Println("Error -- this should never happen -- assuming valid data")
	return 0 // should never happen
}



// implement sorted map ?!
// then do binary search for finding in which range inside the mapper is
