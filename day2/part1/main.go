package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var fileName = "day2/input.txt"

type RockScissorPaper struct {
	Name           string
	Point          int
	WinAgainst     string
	OpponentLetter string
	OurLetter      string
}

var data = []RockScissorPaper{
	{
		Name:           "rock",
		Point:          1,
		WinAgainst:     "scissor",
		OpponentLetter: "A",
		OurLetter:      "X",
	},
	{
		Name:           "paper",
		Point:          2,
		WinAgainst:     "rock",
		OpponentLetter: "B",
		OurLetter:      "Y",
	},
	{
		Name:           "scissor",
		Point:          3,
		WinAgainst:     "paper",
		OpponentLetter: "C",
		OurLetter:      "Z",
	},
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

	total := 0
	lostPoint := 0
	drawPoint := 3
	winPoint := 6

	for sc.Scan() {
		line := sc.Text()
		inputLine := strings.Split(line, " ")
		opponentLetter, ourLetter := inputLine[0], inputLine[1]
		opponent := searchByField("opponentLetter", opponentLetter)
		our := searchByField("ourLetter", ourLetter)

		if our.Name == opponent.Name { // draw
			total += (drawPoint + our.Point)
		} else if our.WinAgainst == opponent.Name { // win
			total += (winPoint + our.Point)
		} else {
			total += (lostPoint + our.Point) // lost
		}
	}

	fmt.Println("Total Point: " + strconv.Itoa(total))
}

func searchByField(field string, param string) RockScissorPaper {
	for _, rockScissorPapper := range data {
		var value string
		switch field {
		case "opponentLetter":
			value = rockScissorPapper.OpponentLetter
		case "ourLetter":
			value = rockScissorPapper.OurLetter
		}

		if value == param {
			return rockScissorPapper
		}
	}
	return RockScissorPaper{}
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}
