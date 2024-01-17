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
	
	start = time.Now()
	part2result := solvePart2()
	elapsed = time.Since(start)
	fmt.Printf("Time took %s\n", elapsed)
	fmt.Println("PART2 - RESULT - ", part2result)
	

	if part1result == 703315117 && part2result == 0 {
		fmt.Println(" G O O D ")
	} else {
		fmt.Println(" B A D ")
	}
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

	return lowPulses * highPulses
}

//2nd example out of order ?!?!?maybe
// recurse in reverse
func solvePart2() int {
	
	name := "broadcaster"
	signal := utils.NewSignal("low")
	
	modules := getData()
	btmPresses := 0
	
	for {
		//modules := getData()
		status := handleBtnPress2(signal, name, &modules)
		btmPresses += 1
		if status {
			break
		}
	}

	fmt.Println("PRESSES -> ", btmPresses)
	return btmPresses
}





func updateCounts(lowPulses, hightPulses *int, pulse utils.Signal) {

	if pulse.Status == utils.Low {
		*lowPulses++
	} else {
		*hightPulses++
	}
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
		//fmt.Printf(" <%s> -%s-> <%s>\n",moduleName, signal.Status, nextModuleName)
		updateCounts(&lowPulses, &highPulses, signal)
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


func handleBtnPress2(signal utils.Signal, moduleName string, modules *utils.Modules) bool {
	
	// stopping conditions
	module, err := modules.Get(moduleName)
	if err != nil {
		return false
	}
	
	nextSignals := []utils.Signal {}
	nextModuleNames := []string {}
	for _, nextModuleName := range module.Destinations {

		nextModule, err := modules.Get(nextModuleName)
		
		if nextModuleName == "rx" && signal.Status == utils.Low {
			fmt.Println("FOUND")
			return true
		}
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

	status := false
	for idx := 0; idx < len(nextSignals); idx++ {

		sig := nextSignals[idx]
		nam := nextModuleNames[idx]
		status = handleBtnPress2(sig, nam, modules)
		if status {
			return status
		}
	
	}
	return false
}