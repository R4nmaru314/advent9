package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func part2() int {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	coordinatesHead := []Coordinate{{x: 0, y: 0}}
	coordinatesTails := make([][]Coordinate, 9)
	for i := 0; i < 9; i++ {
		coordinatesTails[i] = []Coordinate{{x: 0, y: 0}}
	}
	lastCoordinate := Coordinate{x: 0, y: 0}
	var lastDirection string

	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, " ")
		direction := values[0]
		count, _ := strconv.Atoi(values[1])
		calculateCoordinatesPart2(&coordinatesHead, coordinatesTails, &lastCoordinate, direction, count, &lastDirection)
	}
	lastHeadCoordinate := (coordinatesHead)[len(coordinatesHead)-1]
	calculateCoordinatesPart2(&coordinatesHead, coordinatesTails, &lastHeadCoordinate, lastDirection, 1, &lastDirection)
	coordinatesHead = coordinatesHead[:len(coordinatesHead)-1]
	for i := 0; i < 9; i++ {
		log.Println(i, coordinatesTails[i])
	}
	return len(removeDuplicates(coordinatesTails[8]))
}

func calculateCoordinatesPart2(coordinatesHead *[]Coordinate, coordinatesTails [][]Coordinate, lastCoordinates *Coordinate, direction string, count int, lastDirection *string) {

	if direction == "R" {
		calculateCoordinatesRightPart2(coordinatesHead, coordinatesTails, lastCoordinates, count)
	} else if direction == "L" {
		calculateCoordinatesLeftPart2(coordinatesHead, coordinatesTails, lastCoordinates, count)
	} else if direction == "U" {
		calculateCoordinatesUpPart2(coordinatesHead, coordinatesTails, lastCoordinates, count)
	} else if direction == "D" {
		calculateCoordinatesDownPart2(coordinatesHead, coordinatesTails, lastCoordinates, count)
	}
	*lastDirection = direction
}

func calculateCoordinatesRightPart2(coordinatesHead *[]Coordinate, coordinatesTails [][]Coordinate, lastCoordinate *Coordinate, count int) {
	for i := 0; i < count; i++ {
		newCoordinate := Coordinate{lastCoordinate.x + 1, lastCoordinate.y}
		*coordinatesHead = append(*coordinatesHead, newCoordinate)
		calculateTailsPart2(coordinatesTails, lastCoordinate)
		*lastCoordinate = newCoordinate
	}
}

func calculateCoordinatesLeftPart2(coordinatesHead *[]Coordinate, coordinatesTails [][]Coordinate, lastCoordinate *Coordinate, count int) {
	for i := 0; i < count; i++ {
		newCoordinate := Coordinate{lastCoordinate.x - 1, lastCoordinate.y}
		*coordinatesHead = append(*coordinatesHead, newCoordinate)
		calculateTailsPart2(coordinatesTails, lastCoordinate)
		*lastCoordinate = newCoordinate
	}
}

func calculateCoordinatesUpPart2(coordinatesHead *[]Coordinate, coordinatesTails [][]Coordinate, lastCoordinate *Coordinate, count int) {
	for i := 0; i < count; i++ {
		newCoordinate := Coordinate{lastCoordinate.x, lastCoordinate.y + 1}
		*coordinatesHead = append(*coordinatesHead, newCoordinate)
		calculateTailsPart2(coordinatesTails, lastCoordinate)
		*lastCoordinate = newCoordinate
	}
}

func calculateCoordinatesDownPart2(coordinatesHead *[]Coordinate, coordinatesTails [][]Coordinate, lastCoordinate *Coordinate, count int) {
	for i := 0; i < count; i++ {
		newCoordinate := Coordinate{lastCoordinate.x, lastCoordinate.y - 1}
		*coordinatesHead = append(*coordinatesHead, newCoordinate)
		calculateTailsPart2(coordinatesTails, lastCoordinate)
		*lastCoordinate = newCoordinate
	}
}

func calculateTailsPart2(coordinatesTails [][]Coordinate, lastCoordinate *Coordinate) {
	calculateTail(&coordinatesTails[0], lastCoordinate)
	*lastCoordinate = (coordinatesTails[0])[len(coordinatesTails[0])-1]
	for i := 1; i < 9; i++ {
		calculateTail(&coordinatesTails[i], lastCoordinate)
		*lastCoordinate = (coordinatesTails[i])[len(coordinatesTails[i])-1]
	}
}
