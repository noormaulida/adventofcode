package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

    mapset "github.com/deckarep/golang-set"
)

var fileName = "day04/input.txt"

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

	total := 0

	for sc.Scan() {
		line := sc.Text()
		elvesTasks := strings.Split(line, ",")

        firstSections := strings.Split(elvesTasks[0], "-")
        firstSectionMin, _:= strconv.Atoi(firstSections[0])
        firstSectionMax, _ := strconv.Atoi(firstSections[1])
        firstSectionSequence := sliceToSet(makeRange(firstSectionMin, firstSectionMax))
        
        secondSections := strings.Split(elvesTasks[1], "-")
        secondSectionMin, _:= strconv.Atoi(secondSections[0])
        secondSectionMax, _ := strconv.Atoi(secondSections[1])
        secondSectionSequence := sliceToSet(makeRange(secondSectionMin, secondSectionMax))
        
        if (firstSectionSequence.IsSubset(secondSectionSequence) || secondSectionSequence.IsSubset(firstSectionSequence)) {
			total++
        }
	}

	fmt.Println("Total Point: " + strconv.Itoa(total))
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func makeRange(min, max int) []int {
    a := make([]int, max-min+1)
    for i := range a {
        a[i] = min + i
    }
    return a
}

func sliceToSet(mySlice []int) mapset.Set {
    mySet := mapset.NewSet()
    for _, ele := range mySlice {
        mySet.Add(ele)
    }   
    return mySet
}