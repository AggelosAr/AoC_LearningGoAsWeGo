package utils

import (
	"fmt"
	"strconv"
	"strings"
)


type Workflow struct {
	Rules []Rule
}

func GetNewWorkflow(text string) (string, Workflow) {
	parts := strings.Split(text, "{")
	name := parts[0]
	rules := strings.Split(strings.TrimRight(parts[1], "}"), ",")

	newWorkflow := Workflow{}
	for _, rule := range rules{
		newWorkflow.Rules = append(newWorkflow.Rules, getNewRule(rule))
	}

	return name, newWorkflow
}

// returns the next Workflow name or A: accepted / R: rejected 
func (w Workflow) Evaluate(part Part) string {

	result := ""
	for _, rule := range w.Rules {
		if rule.Category == "" || rule.Evaluate(part.Data[rule.Category]) {
			result = rule.Result
			break
		}
	}
	return result
}



type Rule struct {
	
	Category string
	Operation string
	Limit int

	Result string
}


func (r Rule) String() string {
	return fmt.Sprintf("{ %s %s %d }", r.Category, r.Operation, r.Limit)
}

func getNewRule(text string) Rule{
	
	newRule := Rule{}

	hasOp := false
	for idx := 0; idx < len(text); idx++ {
		if string(text[idx]) == "<" || string(text[idx]) == ">" {
			hasOp = true

			newRule.Operation = string(text[idx])
			break
		}
	}

	if hasOp {

		parts := strings.Split(text, ":")

		comparison := strings.Split(parts[0], newRule.Operation)

		newRule.Category = comparison[0]
		newRule.Limit, _ = strconv.Atoi(comparison[1])

		newRule.Result = parts[1]

	} else {
		
		newRule.Result = text
	}

	return newRule
}


func (r Rule) Evaluate(v int) bool {
	
	isValid := false
	if r.Operation != "" {
		if r.Operation == ">" {
			isValid = v > r.Limit
		} else {
			isValid = v < r.Limit
		}	
	} else {
		isValid = true
	}
	return isValid
}



type Part struct {
	Data map[string]int
	Worth int
}

func GetNewPart(text string) Part {

	parts := strings.Split(text[1 : len(text)-1], ",")
	
	newPart := Part{Data: map[string]int {}}
	for _, part := range parts {
		category, rating := getPartData(part)
		newPart.Data[category] = rating
		newPart.Worth += rating
	}

	return newPart
}

func getPartData(text string) (string, int) {
	parts := strings.Split(text, "=")
	rating, _ := strconv.Atoi(parts[1])
	return parts[0], rating
}


//start with ranges [1, 40000]

// evcaluate untill A
// stop when all ramghes are full
