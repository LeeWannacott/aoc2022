package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type whatYouShouldPlay struct {
	rock    int
	paper   int
	scissor int
}
type opponentPlayed struct {
	A string
	B string
	C string
}

func main() {
	fmt.Println("Hello, World!")
	// read txt file line by line
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example

	scoreYourHand := whatYouShouldPlay{
		rock:    1,
		paper:   2,
		scissor: 3,
	}

	var score = 0
	for scanner.Scan() {
		// x rock, y, paper, z scissor
		var opponentsHand = strings.Split(scanner.Text(), " ")[0]
		var yourHand = strings.Split(scanner.Text(), " ")[1]
		if yourHand == "X" {
			score += scoreYourHand.rock
		} else if yourHand == "Y" {
			score += scoreYourHand.paper
		} else if yourHand == "Z" {
			score += scoreYourHand.scissor
		}

		// x rock, y, paper, z scissor
		// a rock, b, paper, c scissor
		// draw
		if yourHand == "X" && opponentsHand == "A" || yourHand == "Y" && opponentsHand == "B" || yourHand == "Z" && opponentsHand == "C" {
			score += 3
		}
		// win
		if yourHand == "X" && opponentsHand == "C" || yourHand == "Y" && opponentsHand == "A" || yourHand == "Z" && opponentsHand == "B" {
			score += 6
		}

		fmt.Println(score)

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

	}
}
