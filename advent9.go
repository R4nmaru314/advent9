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
	coordinatesTail := []Coordinate{{x: 0, y: 0}}
	lastCoordinate := Coordinate{x: 0, y: 0}
	var lastDirection string

	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, " ")
		direction := values[0]
		count, _ := strconv.Atoi(values[1])
		calculateCoordinates(&coordinatesHead, &coordinatesTail, &lastCoordinate, direction, count, &lastDirection)
	}
	lastHeadCoordinate := (coordinatesHead)[len(coordinatesHead)-1]
	calculateCoordinates(&coordinatesHead, &coordinatesTail, &lastHeadCoordinate, lastDirection, 1, &lastDirection)
	coordinatesHead = coordinatesHead[:len(coordinatesHead)-1]
	log.Println(len(removeDuplicates(coordinatesTail)))
	log.Println(part2())
	log.Println(part3())
}

func calculateCoordinates(coordinatesHead, coordinatesTail *[]Coordinate, lastCoordinates *Coordinate, direction string, count int, lastDirection *string) {

	if direction == "R" {
		calculateCoordinatesRight(coordinatesHead, coordinatesTail, lastCoordinates, count)
	} else if direction == "L" {
		calculateCoordinatesLeft(coordinatesHead, coordinatesTail, lastCoordinates, count)
	} else if direction == "U" {
		calculateCoordinatesUp(coordinatesHead, coordinatesTail, lastCoordinates, count)
	} else if direction == "D" {
		calculateCoordinatesDown(coordinatesHead, coordinatesTail, lastCoordinates, count)
	}
	*lastDirection = direction
}

func calculateCoordinatesRight(coordinatesHead, coordinatesTail *[]Coordinate, lastCoordinate *Coordinate, count int) {
	for i := 0; i < count; i++ {
		newCoordinate := Coordinate{lastCoordinate.x + 1, lastCoordinate.y}
		*coordinatesHead = append(*coordinatesHead, newCoordinate)
		calculateTail(coordinatesTail, lastCoordinate)
		*lastCoordinate = newCoordinate
	}
}

func calculateCoordinatesLeft(coordinatesHead, coordinatesTail *[]Coordinate, lastCoordinate *Coordinate, count int) {
	for i := 0; i < count; i++ {
		newCoordinate := Coordinate{lastCoordinate.x - 1, lastCoordinate.y}
		*coordinatesHead = append(*coordinatesHead, newCoordinate)
		calculateTail(coordinatesTail, lastCoordinate)
		*lastCoordinate = newCoordinate
	}
}

func calculateCoordinatesUp(coordinatesHead, coordinatesTail *[]Coordinate, lastCoordinate *Coordinate, count int) {
	for i := 0; i < count; i++ {
		newCoordinate := Coordinate{lastCoordinate.x, lastCoordinate.y + 1}
		*coordinatesHead = append(*coordinatesHead, newCoordinate)
		calculateTail(coordinatesTail, lastCoordinate)
		*lastCoordinate = newCoordinate
	}
}

func calculateCoordinatesDown(coordinatesHead, coordinatesTail *[]Coordinate, lastCoordinate *Coordinate, count int) {
	for i := 0; i < count; i++ {
		newCoordinate := Coordinate{lastCoordinate.x, lastCoordinate.y - 1}
		*coordinatesHead = append(*coordinatesHead, newCoordinate)
		calculateTail(coordinatesTail, lastCoordinate)
		*lastCoordinate = newCoordinate
	}
}

func calculateTail(coordinatesTail *[]Coordinate, lastCoordinate *Coordinate) {
	lastTailCoordinate := (*coordinatesTail)[len(*coordinatesTail)-1]

	if isTwoUnitsAwayX(*lastCoordinate, lastTailCoordinate) {
		if lastTailCoordinate.x < lastCoordinate.x {
			*coordinatesTail = append(*coordinatesTail, Coordinate{lastTailCoordinate.x + 1, lastCoordinate.y})
		} else {
			*coordinatesTail = append(*coordinatesTail, Coordinate{lastTailCoordinate.x - 1, lastCoordinate.y})
		}
	} else if isTwoUnitsAwayY(*lastCoordinate, lastTailCoordinate) {
		if lastTailCoordinate.y < lastCoordinate.y {
			*coordinatesTail = append(*coordinatesTail, Coordinate{lastCoordinate.x, lastTailCoordinate.y + 1})
		} else {
			*coordinatesTail = append(*coordinatesTail, Coordinate{lastCoordinate.x, lastTailCoordinate.y - 1})
		}
	}
}

func isTwoUnitsAwayX(coordinate1, coordinate2 Coordinate) bool {
	dx := int(math.Abs(float64(coordinate1.x - coordinate2.x)))
	return dx == 2
}

func isTwoUnitsAwayY(coordinate1, coordinate2 Coordinate) bool {
	dy := int(math.Abs(float64(coordinate1.y - coordinate2.y)))
	return dy == 2
}

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
