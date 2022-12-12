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

    counter := 0
    groupCounter := 0
    elvesGroups := make([][]string, 100)
    var ss []string

    for sc.Scan() {
        counter++
        line := sc.Text()
        ss = append(ss, line)
        if counter >= 3 {
            elvesGroups[groupCounter] = ss
            ss = []string{}
            counter = 0
            groupCounter++
        }
    }

    total := 0

    for index, group := range elvesGroups {
        if (len(group) == 3) {
            group1 := group[0]
            group2 := []rune(group[1])
            group3 := []rune(group[2])
            for _, item := range group1 {
                if slices.Contains(group2, item) && slices.Contains(group3, item) {
                    sameItem := (item)
                    fmt.Println("Group: ", group)
                    fmt.Println("Indec: ", index)
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
    }
    fmt.Println("Total Point: " + strconv.Itoa(total))
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}
