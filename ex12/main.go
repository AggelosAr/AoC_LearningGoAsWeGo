package main

import (
	"bufio"
	
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"aggelos.com/go/aoc/ex12/recordUtils"
)

func main() {

	records := getData()
	
	start := time.Now()
	part1result := solvePart1(records)
	elapsed := time.Since(start)
	fmt.Printf("Time took %s\n", elapsed)
	fmt.Println("PART1 - RESULT - ", part1result)


	start = time.Now()
	unfoldAll(records)
	part2result := solvePart2(records)
	elapsed = time.Since(start)
	
	fmt.Printf("Time took %s\n", elapsed)
	fmt.Println("PART2 - RESULT - ", part2result)

	if part1result == 7221 && part2result == 7139671893722 {
		fmt.Println(" G O O D ")
	} else {
		fmt.Println(" B A D ")
	}

}


func getFilePath() string {
	rootPath, _ := os.Getwd()
	return filepath.Join(rootPath, "data", "input.txt")

}


func getData() []recordUtils.Record {

	file, err := os.Open(getFilePath())
	if err != nil {
		fmt.Println("Error opening file : ", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	data := []recordUtils.Record {}
	for scanner.Scan() {
		line := scanner.Text()
		currentRecord := getRecord(line)
		data = append(data,currentRecord )	
	}

	return data
}

func getRecord(line string) recordUtils.Record {
	
	parts := strings.Fields(line)
	data := strings.Split(parts[0], "")
	groups := []int {}

	for _, el := range strings.Split(parts[1], ",") {
		number, _ := strconv.Atoi(el)
		groups = append(groups, number)
	}
	return recordUtils.Record{Data: data, Groups: groups}
}

func solvePart1(records []recordUtils.Record) int {
	totalWays := 0
	for _, record := range records {
		memo := map[[3]int]int {}
		totalWays += record.SolveRecord(0, 0, 0, memo)
		//fmt.Printf(" - %v - %d arrangements\n", record.Data, current_ways)
	}
	return totalWays
}


func unfoldAll(records []recordUtils.Record) {
	for idx := range records {
        records[idx].Unfold()
    }
	
}
func solvePart2(records []recordUtils.Record) int {
	totalWays := 0
	for _, record := range records {
		memo := map[[3]int]int {}
		totalWays += record.SolveRecord(0, 0, 0, memo)
	}
	return totalWays
}

