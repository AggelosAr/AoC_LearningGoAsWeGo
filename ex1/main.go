package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {

	data := getData()
	total := 0

	for _, line := range data {
		result := decode(line)
		total += result
	}

	fmt.Printf("Total : %d", total)
}

	

type occurence struct {
	number string
	start, end int
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

	for key := range numberStrings {
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
				
		}
		if last.number != ""{
			text = replaceAtIdx(text, last.start, last.end, numberStrings[last.number])
		}
	}	

	return text

}

func replaceAtIdx(t string, start int, end int, r string) string {
	return string(t[:start]) + r + string(t[end:])
}


func decode(text string) int {

	text = replaceFirstLastWords(text)

	var (
		size = len(text) - 1
		value = ""
		idx = 0
	)
	

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

	num, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}

	return num
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
		data = append(data, scanner.Text())
	}

	return data
}
