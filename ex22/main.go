package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"aggelos.com/go/aoc/ex22/utils"
)


func main() {

	bricks := getData()
	
	start := time.Now()
	part1result := solvePart1(bricks)
	elapsed := time.Since(start)
	fmt.Printf("Time took %s\n", elapsed)
	fmt.Println("PART1 - RESULT - ", part1result)
	
	start = time.Now()
	part2result := solvePart2(bricks)
	elapsed = time.Since(start)
	fmt.Printf("Time took %s\n", elapsed)
	fmt.Println("PART2 - RESULT - ", part2result)
	

	
	if part1result == 0 && part2result == 0 {
		fmt.Println(" G O O D ")
	} else {
		fmt.Println(" B A D ")
	}

}


func getFilePath() string {
	rootPath, _ := os.Getwd()
	return filepath.Join(rootPath, "data", "input.txt")
}


func getData() []utils.Brick {

	file, err := os.Open(getFilePath())
	if err != nil {
		fmt.Println("Error opening file : ", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	
	bricks := []utils.Brick {}

	brickId := 1
	
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "~")

		start := parts[0]
		end := parts[1]

		x1, y1, z1 := getNums(start)
    	x2, y2, z2 := getNums(end)
		currentBrick := utils.GetNewBrick(brickId, x1, y1, z1, x2, y2, z2)
		bricks = append(bricks, currentBrick)
		brickId++

		
	}

	return bricks
}


func getNums(text string) (int, int, int) {
	nums := []int {}

	for _, numStr := range strings.Split(text, ",") {
		num, _ := strconv.Atoi(numStr)
		nums = append(nums, num)
	}
	return nums[0], nums[1], nums[2]
}


func collapseBricks(bricks []utils.Brick) int {


	maxX, maxY, maxZ := findMaxXYZ(bricks)
	

	xView := getXYAxisView(maxX, maxZ, "X", bricks)
	yView := getXYAxisView(maxY, maxZ, "Y", bricks)
	printView(xView)
	printView(yView)


	
	
	return 0
}


func solvePart2(bricks []utils.Brick) int {
	return 0
}



func findMaxXYZ(bricks []utils.Brick) (int, int, int) {
	maxX, maxY, maxZ := -1, -1, -1

	for _, brick := range bricks {
		maxX = max(brick.Start.X, maxX)
		maxY = max(brick.Start.Y, maxY)
		maxZ = max(brick.Start.Z, maxZ)

		maxX = max(brick.End.X, maxX)
		maxY = max(brick.End.Y, maxY)
		maxZ = max(brick.End.Z, maxZ)
	}

	return maxX, maxY, maxZ
}



func getXYAxisView(maxXY, maxZ int, axis string, bricks []utils.Brick) [][]int {

	view := createArr(maxZ, maxXY)

	for _, brick := range bricks {
		
		for z := brick.Start.Z; z < brick.End.Z + 1; z++ {
			start := 0
			end := 0
			if axis == "X" {
				start = brick.Start.X
				end = brick.End.X + 1
			} else {
				start = brick.Start.Y
				end = brick.End.Y + 1
			}
			for xy := start; xy < end; xy++ {
				view[z][xy] = brick.Id
			}
		}
	}
	
	return view
}

func createArr(rows, cols int) [][]int {
	rows++
	cols++
	arr := make([][]int, rows)
	for i := range arr {
		arr[i] = make([]int, cols)
	}
	return arr
}

func printView(view [][]int) {
	for idx := len(view) - 1; idx > -1 ; idx-- {
		fmt.Println(idx, view[idx])
	}
	fmt.Println("_______________________")
}



