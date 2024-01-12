package nodeUtils

type Node struct {
	Left  string
	Right string
}

type NodeMap struct {
	Instructions string
	Paths        map[string]Node
}

func (n NodeMap) WalkMap(startingNode string, relaxedEnding bool) int {
	steps := 0
	incructionIdx := 0

	currentNode := startingNode
	direction := ""
	nodeDestinations := Node{}

	for {
		if incructionIdx == len(n.Instructions) {
			incructionIdx = 0
		}

		direction = string(n.Instructions[incructionIdx])
		nodeDestinations, _ = n.Paths[currentNode]

		nextNode := walkCurrentNode(direction, nodeDestinations)

		if relaxedEnding {
			if string(currentNode[2]) == "Z" {
				break
			}
		} else {
			if currentNode == "ZZZ" {
				break
			}
		}

		incructionIdx++
		steps++
		currentNode = nextNode
	}
	return steps
}

func walkCurrentNode(direction string, nodeDestinations Node) string {
	nextNode := ""
	if direction == "L" {
		nextNode = nodeDestinations.Left
	} else {
		nextNode = nodeDestinations.Right
	}
	return nextNode
}

func (n NodeMap) WalkMapPart2(relaxedEnding bool) int {

	steps := []int{}
	for _, node := range n.getStartNodes() {

		steps = append(steps, n.WalkMap(node, relaxedEnding))

	}
	// Lowest Common Multiple  of all the steps
	return LCM(steps)
}

func (n NodeMap) getStartNodes() []string {
	startingNodes := []string{}

	for node := range n.Paths {
		if string(node[2]) == "A" {
			startingNodes = append(startingNodes, node)
		}
	}
	return startingNodes
}

func checkNodesAreFinished(nodes []string) bool {
	for _, node := range nodes {
		if string(node[2]) != "Z" {
			return false
		}
	}
	return true
}


// SOURCE -> https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(numbers []int) int {
	result := numbers[0]

	for i := 1; i < len(numbers); i++ {
		result = (result * numbers[i]) / GCD(result, numbers[i])
	}
	return result
}