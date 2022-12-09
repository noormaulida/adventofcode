package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

var fileName = "day01/input.txt"

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func main() {
	fmt.Println("Opening input file ")
	var file, err = os.OpenFile(fileName, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	if err := sc.Err(); err != nil {
		log.Fatalf("Scan file error: %v", err)
		return
	}

	reindeerTotalCalories := make([]int, 1000)
	count := 0
	total := 0

	for sc.Scan() {
		line := sc.Text()

		if count == 0 && line != "" {
			total, _ = strconv.Atoi(line)
		} else if line != "" {
			total, _ = strconv.Atoi(line)
		} else {
			count++
			total = 0 // reset total
		}
		reindeerTotalCalories[count] += total
	}

	sort.Sort(sort.Reverse(sort.IntSlice(reindeerTotalCalories))) // sort descending

	totalTop3 := 0
	for index, totalCalories := range reindeerTotalCalories {
		if totalCalories > 0 && index < 3 {
			totalTop3 += totalCalories
			fmt.Println(totalTop3)
		}
	}

	totalTop3Value := fmt.Sprintf("Top 3 Total Calories: %v", totalTop3)
	fmt.Println(totalTop3Value)

}
