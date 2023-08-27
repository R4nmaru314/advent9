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
		calculateCoordinates(&coordinatesHead, &coordinatesTail, coordinatesTails, &lastCoordinate, direction, count, &lastDirection)
	}
	lastHeadCoordinate := (coordinatesHead)[len(coordinatesHead)-1]
	calculateCoordinates(&coordinatesHead, &coordinatesTail, coordinatesTails, &lastHeadCoordinate, lastDirection, 1, &lastDirection)
	coordinatesHead = coordinatesHead[:len(coordinatesHead)-1]
	log.Println(len(removeDuplicates(coordinatesTail)))
	log.Println(len(removeDuplicates(coordinatesTails[8])))
}

func calculateCoordinates(coordinatesHead, coordinatesTail *[]Coordinate, coordinatesTails [][]Coordinate, lastCoordinates *Coordinate, direction string, count int, lastDirection *string) {

	if direction == "R" {
		calculateCoordinatesRight(coordinatesHead, coordinatesTail, coordinatesTails, lastCoordinates, count)
	} else if direction == "L" {
		calculateCoordinatesLeft(coordinatesHead, coordinatesTail, coordinatesTails, lastCoordinates, count)
	} else if direction == "U" {
		calculateCoordinatesUp(coordinatesHead, coordinatesTail, coordinatesTails, lastCoordinates, count)
	} else if direction == "D" {
		calculateCoordinatesDown(coordinatesHead, coordinatesTail, coordinatesTails, lastCoordinates, count)
	}
	*lastDirection = direction
}

func calculateCoordinatesRight(coordinatesHead, coordinatesTail *[]Coordinate, coordinatesTails [][]Coordinate, lastCoordinate *Coordinate, count int) {
	for i := 0; i < count; i++ {
		newCoordinate := Coordinate{lastCoordinate.x + 1, lastCoordinate.y}
		*coordinatesHead = append(*coordinatesHead, newCoordinate)
		calculateTail(coordinatesTail, lastCoordinate)
		calculateTails(coordinatesTails, lastCoordinate)
		*lastCoordinate = newCoordinate
	}
}

func calculateCoordinatesLeft(coordinatesHead, coordinatesTail *[]Coordinate, coordinatesTails [][]Coordinate, lastCoordinate *Coordinate, count int) {
	for i := 0; i < count; i++ {
		newCoordinate := Coordinate{lastCoordinate.x - 1, lastCoordinate.y}
		*coordinatesHead = append(*coordinatesHead, newCoordinate)
		calculateTail(coordinatesTail, lastCoordinate)
		calculateTails(coordinatesTails, lastCoordinate)
		*lastCoordinate = newCoordinate
	}
}

func calculateCoordinatesUp(coordinatesHead, coordinatesTail *[]Coordinate, coordinatesTails [][]Coordinate, lastCoordinate *Coordinate, count int) {
	for i := 0; i < count; i++ {
		newCoordinate := Coordinate{lastCoordinate.x, lastCoordinate.y + 1}
		*coordinatesHead = append(*coordinatesHead, newCoordinate)
		calculateTail(coordinatesTail, lastCoordinate)
		calculateTails(coordinatesTails, lastCoordinate)
		*lastCoordinate = newCoordinate
	}
}

func calculateCoordinatesDown(coordinatesHead, coordinatesTail *[]Coordinate, coordinatesTails [][]Coordinate, lastCoordinate *Coordinate, count int) {
	for i := 0; i < count; i++ {
		newCoordinate := Coordinate{lastCoordinate.x, lastCoordinate.y - 1}
		*coordinatesHead = append(*coordinatesHead, newCoordinate)
		calculateTail(coordinatesTail, lastCoordinate)
		calculateTails(coordinatesTails, lastCoordinate)
		*lastCoordinate = newCoordinate
	}
}

func calculateTails(coordinatesTails [][]Coordinate, lastCoordinate *Coordinate) {
	calculateTail(&coordinatesTails[0], lastCoordinate)
	*lastCoordinate = (coordinatesTails[0])[len(coordinatesTails[0])-1]
	for i := 1; i < 9; i++ {
		calculateTail(&coordinatesTails[i], lastCoordinate)
		*lastCoordinate = (coordinatesTails[i])[len(coordinatesTails[i])-1]
	}
}

func calculateTail(coordinatesTail *[]Coordinate, lastCoordinate *Coordinate) {
	lastTailCoordinate := (*coordinatesTail)[len(*coordinatesTail)-1]

	if isTwoUnitsAwayY(*lastCoordinate, lastTailCoordinate) && isTwoUnitsAwayX(*lastCoordinate, lastTailCoordinate) {
		if lastTailCoordinate.y < lastCoordinate.y {
			if lastTailCoordinate.x < lastCoordinate.x {
				*coordinatesTail = append(*coordinatesTail, Coordinate{lastTailCoordinate.x + 1, lastTailCoordinate.y + 1})

			} else {
				*coordinatesTail = append(*coordinatesTail, Coordinate{lastTailCoordinate.x - 1, lastTailCoordinate.y + 1})
			}
		} else {
			if lastTailCoordinate.x < lastCoordinate.x {
				*coordinatesTail = append(*coordinatesTail, Coordinate{lastTailCoordinate.x + 1, lastTailCoordinate.y - 1})
			} else {
				*coordinatesTail = append(*coordinatesTail, Coordinate{lastTailCoordinate.x - 1, lastTailCoordinate.y - 1})
			}
		}
	} else if isTwoUnitsAwayX(*lastCoordinate, lastTailCoordinate) {
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
