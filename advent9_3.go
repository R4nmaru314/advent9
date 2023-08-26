package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var knots = make([][2]int, 10)

var dd = map[string][2]int{
	"R": {1, 0},
	"U": {0, 1},
	"L": {-1, 0},
	"D": {0, -1},
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	tailVisited := make(map[[2]int]struct{})
	tailVisited[[2]int{knots[9][0], knots[9][1]}] = struct{}{}

	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, " ")
		op := values[0]
		amount, _ := strconv.Atoi(values[1])
		dx, dy := dd[op][0], dd[op][1]

		for i := 0; i < amount; i++ {
			move(dx, dy)
			tailVisited[[2]int{knots[9][0], knots[9][1]}] = struct{}{}
		}
	}
	fmt.Println(len(tailVisited))
}

func touching(x1, y1, x2, y2 int) bool {
	return int(math.Abs(float64(x1-x2))) <= 1 && int(math.Abs(float64(y1-y2))) <= 1
}

func move(dx, dy int) {
	knots[0][0] += dx
	knots[0][1] += dy

	for i := 1; i < 10; i++ {
		hx, hy := knots[i-1][0], knots[i-1][1]
		tx, ty := knots[i][0], knots[i][1]

		if !touching(hx, hy, tx, ty) {
			signX := 0
			signY := 0
			if hx == tx {
				signX = 0
			} else {
				signX = (hx - tx) / int(math.Abs(float64(hx-tx)))
			}
			if hy == ty {
				signY = 0
			} else {
				signY = (hy - ty) / int(math.Abs(float64(hy-ty)))
			}

			tx += signX
			ty += signY
		}
		knots[i][0] = tx
		knots[i][1] = ty
	}
}
