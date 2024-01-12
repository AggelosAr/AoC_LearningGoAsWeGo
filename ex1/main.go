package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	

	rootPath, _  := os.Getwd()

	rootPath += "\\data\\input.txt"


	file, err := os.Open(rootPath)
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	var total int
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("Current line is : %s\n", line)
		result := decode(line)
		fmt.Printf("Result is : %d\n\n", result)
		
		total += result


	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Total : %d", total)

	
	
}

	


type occurence struct {
	number string
	start int
	end int
}


func (el occurence) String() string {
	return fmt.Sprintf("START <%d> END <%d> VAL <%s>\n", el.start, el.end, el.number)
}



func replaceFirstLastWords( text string ) string {
	numberStrings := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	
	first := occurence{start: 10000000, number: ""}
	last  := occurence{start: -1, number: ""}
	for key, _ := range numberStrings {
		firstIdx := strings.Index(text, key)
		if firstIdx != -1 && firstIdx < first.start{
			first.start = firstIdx
			first.end = firstIdx + len(key)
			first.number = key

		}
		lastIdx := strings.LastIndex(text, key)
		if lastIdx != -1 && lastIdx > last.start{
			last.start = lastIdx
			last.end = lastIdx + len(key)
			last.number = key
		}
	}

	//fmt.Printf("FIRST -> %s\nLast -> %s\n", first, last)
	if first.start == last.start {
		text = replaceAtIdx(text, first.start, first.end, numberStrings[first.number])
		
	} else {
		if first.number != "" {
			text = replaceAtIdx(text, first.start, first.end, numberStrings[first.number])
			
			if last.number != "" && first.start + len(first.number) > last.start {
				last.start += 1
				last.end -= 1
				
			}
			
			last.start -= len(first.number) - 1
			last.end -= len(first.number) - 1
			
			//fmt.Printf("FIRST -> %sLast -> %s\n", first, last)
			
		}
		if last.number != ""{
			text = replaceAtIdx(text, last.start, last.end, numberStrings[last.number])
		}
	}	

	//fmt.Printf("--------- New line is : %s\n", text)
	return text

}

func replaceAtIdx(mainText string, start int, end int, replacement string) string {
	part1 := string(mainText[:start])
	part3 := string(mainText[end:])
	return part1 + replacement + part3
}

func decode(text string) int {

	
	
	text = replaceFirstLastWords(text)
	size := len(text) - 1
	value := ""



	var idx int

	
	idx = 0
	for idx < size + 1 {
		_, err := strconv.Atoi(string(text[idx]))
		
		if err == nil {
			value += string(text[idx])
			break
		}
		idx++
	}

	idx = size
	for idx >= 0 {
		_, err := strconv.Atoi(string(text[idx]))
		if err == nil {
			value += string(text[idx])
			break
		}
		idx--
	}

	number, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}

	return number
}



