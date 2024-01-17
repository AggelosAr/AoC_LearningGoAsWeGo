package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"aggelos.com/go/aoc/ex19/utils"
)


func main() {

	workflows, parts := getData()
	
	start := time.Now()
	part1result := solvePart1(workflows, parts)
	elapsed := time.Since(start)
	fmt.Printf("Time took %s\n", elapsed)
	fmt.Println("PART1 - RESULT - ", part1result)
	
	start = time.Now()
	part2result := solvePart2(workflows)
	elapsed = time.Since(start)
	fmt.Printf("Time took %s\n", elapsed)
	fmt.Println("PART2 - RESULT - ", part2result)
	

	
	isDefaultTestCase := false

	if isDefaultTestCase {
		if part1result == 19114 && part2result == 167409079868000 {
			fmt.Println(" G O O D ")
		} else {
			fmt.Println(" B A D ")
		}
	} else {
		if part1result == 409898 && part2result == 113057405770956 {
			fmt.Println(" G O O D ")
		} else {
			fmt.Println(" B A D ")
		}
	}

}


func getFilePath() string {
	rootPath, _ := os.Getwd()
	return filepath.Join(rootPath, "data", "input.txt")
}


func getData() (map[string]utils.Workflow, []utils.Part) {

	file, err := os.Open(getFilePath())
	if err != nil {
		fmt.Println("Error opening file : ", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	workflows := map[string]utils.Workflow{}
	parts := []utils.Part {}

	isWorkflow := true
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			isWorkflow = !isWorkflow
			continue
		}

		if isWorkflow {
			name, workflow := utils.GetNewWorkflow(line)
			workflows[name] = workflow
		} else {
			parts = append(parts, utils.GetNewPart(line))
		}

	}

	return workflows, parts
}




func solvePart1(workflows map[string]utils.Workflow, parts []utils.Part) int {
	result := 0
	for _, part := range parts {
		//fmt.Println("PART -> ", part)
		result += evaluatePart(workflows, part)
		//fmt.Println("___________________________")
	}
	return result
}


func evaluatePart(workflows map[string]utils.Workflow, part utils.Part) int {

	workflowName := "in"

	for {
		//fmt.Println("  WORKFLOW -> ", workflowName)
		currentWorkflow := workflows[workflowName]
		result := currentWorkflow.Evaluate(part)

		if result == "A" {
			return part.Worth
		} else if result == "R" {
			return 0
		} else {
			workflowName = result
		}
	}
	
}

func solvePart2(workflows map[string]utils.Workflow) int {
	return findLimits("in", workflows, getNewRanges())
}



//Recursively tries to find all A: Accepts in all workflows
func findLimits(workflowName string, workflows map[string]utils.Workflow, ranges map[string][2]int) int {
	
	combinations := 0
	workflow := workflows[workflowName]
	
	// Create a deep copy of the current ranges
	currentRanges := copyRanges(ranges)
	
	for _, r := range workflow.Rules {
		//fmt.Println("  -RULE : ", r)
		
		minLimit := currentRanges[r.Category][0]
		maxLimit := currentRanges[r.Category][1]
		
		rangesPassed := [2]int{}
		rangesFailed := [2]int{}

		if r.Operation != "" {
			if r.Operation == ">" {
				rangesPassed = [2]int{r.Limit + 1, maxLimit}
				rangesFailed = [2]int{minLimit, r.Limit}
			} else {
				rangesPassed = [2]int{minLimit, r.Limit - 1}
				rangesFailed = [2]int{r.Limit, maxLimit}
			}	
		}

		// ranges to send to the next Workflow or Accept
		currentRanges[r.Category] = rangesPassed
		//printLimits(currentRanges)

		if r.Result == "A" {
			combinations += findCombinations(currentRanges)
		} else if r.Result == "R" {
			//pass
		} else {
			combinations += findLimits(r.Result, workflows, currentRanges)
		}
		// ranges to send to the next Rule or Reject 
		// Modify current since we are not looking back
		currentRanges[r.Category] = rangesFailed

		//fmt.Println("___________________________")
	}
	return combinations
}


func copyRanges(ranges map[string][2]int) map[string][2]int {
	newRanges := getNewRanges()

	for k, v := range ranges {
		newRanges[k] = v
	}

	return newRanges
}

 
func getNewRanges() map[string][2]int {
	return map[string][2]int {
		"x": {1, 4000},
		"m": {1, 4000},
		"a": {1, 4000},
		"s": {1, 4000},
	}
}


func findCombinations(ranges map[string][2]int) int {
	res := 1
	for _, v := range ranges {
		res *= (v[1] - v[0] + 1)
	}
	return res
}


func printLimits(ranges map[string][2]int) {
	fmt.Println()
	for k, v := range ranges {
		fmt.Printf("%s -> %v\n", k, v)
	}
	fmt.Println()
}

