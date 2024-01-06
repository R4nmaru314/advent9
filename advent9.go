package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Coordinate struct {
	x int
	y int
}

const file = "input.txt"

func main() {
	file, _ := os.Open(file)
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
		calculateCoordinates(&coordinatesHead, coordinatesTails, &lastCoordinate, direction, count, &lastDirection)
	}
	lastHeadCoordinate := (coordinatesHead)[len(coordinatesHead)-1]
	calculateCoordinates(&coordinatesHead, coordinatesTails, &lastHeadCoordinate, lastDirection, 1, &lastDirection)
	coordinatesHead = coordinatesHead[:len(coordinatesHead)-1]
	log.Println(len(removeDuplicates(coordinatesTails[0])))
	log.Println(len(removeDuplicates(coordinatesTails[8])))
}

// calculateCoordinates updates the head and tail coordinates based on the direction and count of movement.
func calculateCoordinates(coordinatesHead *[]Coordinate, coordinatesTails [][]Coordinate, lastCoordinates *Coordinate, direction string, count int, lastDirection *string) {

	if direction == "R" {
		calculateCoordinatesRight(coordinatesHead, coordinatesTails, lastCoordinates, count)
	} else if direction == "L" {
		calculateCoordinatesLeft(coordinatesHead, coordinatesTails, lastCoordinates, count)
	} else if direction == "U" {
		calculateCoordinatesUp(coordinatesHead, coordinatesTails, lastCoordinates, count)
	} else if direction == "D" {
		calculateCoordinatesDown(coordinatesHead, coordinatesTails, lastCoordinates, count)
	}
	*lastDirection = direction
}

// calculateCoordinatesRight appends new coordinates to the right of the last coordinate.
func calculateCoordinatesRight(coordinatesHead *[]Coordinate, coordinatesTails [][]Coordinate, lastCoordinate *Coordinate, count int) {
	for i := 0; i < count; i++ {
		newCoordinate := Coordinate{lastCoordinate.x + 1, lastCoordinate.y}
		*coordinatesHead = append(*coordinatesHead, newCoordinate)
		calculateTails(coordinatesTails, lastCoordinate)
		*lastCoordinate = newCoordinate
	}
}

// calculateCoordinatesLeft appends new coordinates to the left of the last coordinate.
func calculateCoordinatesLeft(coordinatesHead *[]Coordinate, coordinatesTails [][]Coordinate, lastCoordinate *Coordinate, count int) {
	for i := 0; i < count; i++ {
		newCoordinate := Coordinate{lastCoordinate.x - 1, lastCoordinate.y}
		*coordinatesHead = append(*coordinatesHead, newCoordinate)
		calculateTails(coordinatesTails, lastCoordinate)
		*lastCoordinate = newCoordinate
	}
}

// calculateCoordinatesUp appends new coordinates above the last coordinate.
func calculateCoordinatesUp(coordinatesHead *[]Coordinate, coordinatesTails [][]Coordinate, lastCoordinate *Coordinate, count int) {
	for i := 0; i < count; i++ {
		newCoordinate := Coordinate{lastCoordinate.x, lastCoordinate.y + 1}
		*coordinatesHead = append(*coordinatesHead, newCoordinate)
		calculateTails(coordinatesTails, lastCoordinate)
		*lastCoordinate = newCoordinate
	}
}

// calculateCoordinatesDown appends new coordinates below the last coordinate.
func calculateCoordinatesDown(coordinatesHead *[]Coordinate, coordinatesTails [][]Coordinate, lastCoordinate *Coordinate, count int) {
	for i := 0; i < count; i++ {
		newCoordinate := Coordinate{lastCoordinate.x, lastCoordinate.y - 1}
		*coordinatesHead = append(*coordinatesHead, newCoordinate)
		calculateTails(coordinatesTails, lastCoordinate)
		*lastCoordinate = newCoordinate
	}
}

// calculateTails updates the tail coordinates for each of the 9 tails.
func calculateTails(coordinatesTails [][]Coordinate, lastCoordinate *Coordinate) {
	for i := 0; i < 9; i++ {
		calculateTail(&coordinatesTails[i], lastCoordinate)
		*lastCoordinate = (coordinatesTails[i])[len(coordinatesTails[i])-1]
	}
}

// calculateTail calculates the next point in a tail path.
func calculateTail(coordinatesTail *[]Coordinate, lastCoordinate *Coordinate) {
	lastTailCoordinate := (*coordinatesTail)[len(*coordinatesTail)-1]
	var deltaX, deltaY int

	if lastTailCoordinate.y < lastCoordinate.y {
		deltaY = 1
	} else {
		deltaY = -1
	}

	if lastTailCoordinate.x < lastCoordinate.x {
		deltaX = 1
	} else {
		deltaX = -1
	}

	if isTwoUnitsAwayY(*lastCoordinate, lastTailCoordinate) && isTwoUnitsAwayX(*lastCoordinate, lastTailCoordinate) {
		*coordinatesTail = append(*coordinatesTail, Coordinate{lastTailCoordinate.x + deltaX, lastTailCoordinate.y + deltaY})
	} else if isTwoUnitsAwayX(*lastCoordinate, lastTailCoordinate) {
		*coordinatesTail = append(*coordinatesTail, Coordinate{lastTailCoordinate.x + deltaX, lastCoordinate.y})
	} else if isTwoUnitsAwayY(*lastCoordinate, lastTailCoordinate) {
		*coordinatesTail = append(*coordinatesTail, Coordinate{lastCoordinate.x, lastTailCoordinate.y + deltaY})
	}
}

// isTwoUnitsAwayX checks if two coordinates are two units apart horizontally.
func isTwoUnitsAwayX(coordinate1, coordinate2 Coordinate) bool {
	dx := int(math.Abs(float64(coordinate1.x - coordinate2.x)))
	return dx == 2
}

// isTwoUnitsAwayY checks if two coordinates are two units apart vertically.
func isTwoUnitsAwayY(coordinate1, coordinate2 Coordinate) bool {
	dy := int(math.Abs(float64(coordinate1.y - coordinate2.y)))
	return dy == 2
}

// removeDuplicates removes duplicate coordinates from a slice.
func removeDuplicates(coordinates []Coordinate) []Coordinate {
	uniqueMap := make(map[Coordinate]bool)
	var uniqueCoordinates []Coordinate

	for _, coordinate := range coordinates {
		if !uniqueMap[coordinate] {
			uniqueMap[coordinate] = true
			uniqueCoordinates = append(uniqueCoordinates, coordinate)
		}
	}
	return uniqueCoordinates
}
