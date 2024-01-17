package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"aggelos.com/go/aoc/ex20/utils"
)


func main() {

	modules := getData()
	
	

	start := time.Now()
	part1result := solvePart1(modules)
	elapsed := time.Since(start)
	fmt.Printf("Time took %s\n", elapsed)
	fmt.Println("PART1 - RESULT - ", part1result)
	

	/*
	start = time.Now()
	part2result := solvePart2(signals)
	elapsed = time.Since(start)
	fmt.Printf("Time took %s\n", elapsed)
	fmt.Println("PART2 - RESULT - ", part2result)
	
	isDefaultTestCase := true

	if isDefaultTestCase {
		if part1result == 0 && part2result == 0 {
			fmt.Println(" G O O D ")
		} else {
			fmt.Println(" B A D ")
		}
	} else {
		if part1result == 0 && part2result == 0 {
			fmt.Println(" G O O D ")
		} else {
			fmt.Println(" B A D ")
		}
	}
	*/
}


func getFilePath() string {
	rootPath, _ := os.Getwd()
	return filepath.Join(rootPath, "data", "input.txt")
}


func getData() utils.Modules {

	file, err := os.Open(getFilePath())
	if err != nil {
		fmt.Println("Error opening file : ", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	sources := []string {}
	destinations := [][]string {}

	for scanner.Scan() {
		line := scanner.Text()

		source, currDestinations := formatLine(line)
		sources = append(sources, source)
		destinations = append(destinations, currDestinations)
	}

	return utils.GetNewModulesMap(sources, destinations)
}


func formatLine(line string) (string, []string) {
	parts := strings.Split(line, "->")
	source := strip(parts[0])
	destinationsMustStrip := strings.Split(strip(parts[1]), ",")
	destinations := []string {}
	for _, dest := range destinationsMustStrip {
		destinations = append(destinations, strip(dest))
	}
	return source, destinations
}


func strip(text string) string {
	return strings.TrimRight(strings.TrimLeft(text, " "), " ")
}




func solvePart1(modules utils.Modules) int {

	name := "broadcaster"
	signal := utils.NewSignal("low")

	btmPresses := 1000
	// We dont count the initial push on the handleBtnPress
	lowPulses := btmPresses
	highPulses := 0

	for idx := 0; idx < btmPresses; idx++ {
		currLowPulses, currHighPulses := handleBtnPress(signal, name, modules)
		lowPulses += currLowPulses 
		highPulses += currHighPulses
	}

	fmt.Println("LOW -> ", lowPulses)
	fmt.Println("HIGH -> ", highPulses)
	
	return lowPulses * highPulses
}


func solvePart2(modules utils.Modules) int {
	result := 0	
	return result
}


func handleBtnPress(signal utils.Signal, moduleName string, modules utils.Modules) (int, int) {
	
	
	lowPulses := 0
	highPulses := 0

	// stopping conditions
	module, err := modules.Get(moduleName)
	if err != nil {
		return lowPulses, highPulses
	}
	

	nextSignals := []utils.Signal {}
	nextModuleNames := []string {}
	for _, nextModuleName := range module.Destinations {

		nextModule, err := modules.Get(nextModuleName)
		updateCounts(&lowPulses, &highPulses, signal)
		//fmt.Printf(" <%s> -%s-> <%s>\n",moduleName, signal.Status, nextModuleName)
		if err != nil {
			continue
		}
		
		nextSignal := utils.Signal{}
		
		if nextModule.GetType() == "Conjunction" {
			nextSignal, err = nextModule.Send(signal, moduleName)
		} else {
			nextSignal, err = nextModule.Send(signal)
		}
		
		if err != nil {
			continue
		}
		
		nextSignals = append(nextSignals, nextSignal)
		nextModuleNames = append(nextModuleNames, nextModuleName)

	}

	for idx := 0; idx < len(nextSignals); idx++ {

		sig := nextSignals[idx]
		nam := nextModuleNames[idx]
		currLowPulses, currHighPulses := handleBtnPress(sig, nam, modules)
		lowPulses += currLowPulses
		highPulses += currHighPulses
	
	}
	return lowPulses, highPulses
}


func updateCounts(lowPulses, hightPulses *int, pulse utils.Signal) {

	if pulse.Status == utils.Low {
		*lowPulses++
	} else {
		*hightPulses++
	}
}



//DELETE THIS
// returns the steps, all the borders found
// returns the steps, all the borders found
func runBFS(modules utils.Modules) (int, int) {

	lowPulses := 0
	highPulses := 0

	name := "broadcaster"
	signal := utils.NewSignal("low")

	sendData := utils.SendType{Signal: signal, Name:name}

	q := utils.GetNewQ()
	q.Add(sendData)

	for q.HasNext() {

		nextQ :=  utils.GetNewQ()

		for q.HasNext() {

			data, err := q.PopLeft()
			if err != nil {
				break
			}

			moduleName := data.Name
			signal := data.Signal

			fmt.Println("MODULE -> ", moduleName)

			module, err := modules.Get(moduleName)
			if err != nil {
				continue
			}

			updateCounts(&lowPulses, &highPulses, signal)
			newSignal, err := module.Send(signal)
			if err != nil {
				continue
			}

			

			for _, nextModuleName := range module.Destinations {
				
				updateCounts(&lowPulses, &highPulses, newSignal)
				sendData.Name = nextModuleName
				sendData.Signal = newSignal
				nextQ.Add(sendData)
				fmt.Printf("  ADDING -> %s <%s>\n", nextModuleName, newSignal.Status)
			}
		}
		fmt.Println("_______________")
		q.List = nextQ.List
	}
	return lowPulses, highPulses
}

