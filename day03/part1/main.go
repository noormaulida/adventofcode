package main

import (
	"bufio"
	"fmt"
	"golang.org/x/exp/slices"
	"log"
	"os"
	"strconv"
)

var fileName = "day03/input.txt"

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
		divided := len(line) / 2
		var ss []string
		for i := 1; i < len(line); i++ {
			if i%divided == 0 {
				ss = append(ss, line[:i])
				line = line[i:]
				i = 1
			}
		}

		ss = append(ss, line)
		fmt.Println(ss)

		firstCompartment, secondCompartment := (ss[0]), []rune(ss[1])
		for _, item := range firstCompartment {
			if slices.Contains(secondCompartment, item) {
				sameItem := (item)
				fmt.Println("Same Item: ", string(sameItem))
				if sameItem < 'a' {
					total += int(sameItem - 'A' + 27)
				} else {
					total += int(sameItem - 'a' + 1)
				}
				break
			}
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
