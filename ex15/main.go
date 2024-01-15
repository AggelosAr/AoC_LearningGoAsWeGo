package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"aggelos.com/go/aoc/ex15/utils"
)


func main() {

	data := getData()

	start := time.Now()
	part1result := solvePart1(data)
	elapsed := time.Since(start)
	fmt.Printf("Time took %s\n", elapsed)
	fmt.Println("PART1 - RESULT - ", part1result)

	
	start = time.Now()
	part2result := solvePart2(data)
	elapsed = time.Since(start)

	fmt.Printf("Time took %s\n", elapsed)
	fmt.Println("PART2 - RESULT - ", part2result)

	if part1result == 505427 && part2result == 243747 {
		fmt.Println(" G O O D ")
	} else {
		fmt.Println(" B A D ")
	}

}


func getFilePath() string {
	rootPath, _ := os.Getwd()
	return filepath.Join(rootPath, "data", "input.txt")
}


func getData() []string {

	file, err := os.Open(getFilePath())
	if err != nil {
		fmt.Println("Error opening file : ", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	data := []string {}

	for scanner.Scan() {
		line := scanner.Text()
		data = strings.Split(line, ",")
	}

	return data
}
	


func solvePart1(data []string) int {
	result := 0
	for _, el := range data {
		result += hashFucntion(el)
	}
	return result
}


func solvePart2(data []string) int {

	hashMap := utils.GetNewHashMap()

	for _, el := range data {

		currentEntry := parseDataPart2(el)
		
		if currentEntry.Operation == "REMOVE" {
			hashMap.Remove(currentEntry)
		} else if currentEntry.Operation == "ADD" {
			hashMap.Add(currentEntry)
		}
		
	}

	result := 0
	for _, slot := range hashMap.Fields {
		if slot.List.Len() > 0 {

			slotIdx := 1
			for e := slot.List.Front(); e != nil; e = e.Next() {
				entry := e.Value.(utils.Entry)
				result += (entry.BoxHash + 1) * slotIdx * entry.FocalLength
				slotIdx++
			}
		}
	}
	return result
}


func parseDataPart2(text string) utils.Entry {
	newEntry := utils.Entry {}
	if strings.Contains(text, "-") {
		newEntry = customSplit(text, "-")
		newEntry.Operation = "REMOVE"
	} else if strings.Contains(text, "=") {
		newEntry = customSplit(text, "=")
		newEntry.Operation = "ADD"
	}
	return newEntry
}


func customSplit(text, sep string) utils.Entry {
	newEntry := utils.Entry {}
	parts := strings.Split(text, sep)
	if len(parts) == 2 {
		number, err := strconv.Atoi(parts[1])
		if err == nil {
			newEntry.FocalLength = number
		}
		
	}
	newEntry.Label = parts[0]
	newEntry.BoxHash = hashFucntion(parts[0])
	return newEntry
}


func hashFucntion(text string) int {
	value := 0
	for _, char := range text {
		value += int(char)
		value *= 17
		value %= 256
	}
	return value
}

