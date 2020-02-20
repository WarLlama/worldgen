package main

import (
	"fmt"
	"github.com/aquilax/go-perlin"
	"github.com/logrusorgru/aurora"
	"math"
	"math/rand"
)

type Grid [][]float64
const (
	width = 100
	height = 100
	centerX = 85
	centerY = 70
	character = "\u2588\u2588"
)
func main() {
    grid := Normalize(GenerateNoise())
    grid = Normalize(ApplyGradient(centerX, centerY, grid))
    grid = Normalize(ApplyGradient(20, 30, grid))
    Print(grid)

}

func GenerateNoise() *Grid {
	multiplier := 0.1
	p := perlin.NewPerlin(2, 2, 10, rand.Int63())
	var grid Grid = make([][]float64, 0)

	for x := 0; x < width; x++ {
		row := make([]float64, 0)
		fx := float64(x)
		for y := 0; y < height; y++ {
			fy := float64(y)
			z := p.Noise2D(fx * multiplier, fy * multiplier)
			row = append(row, z)
		}
		grid = append(grid, row)
    }
    return &grid
}

func Normalize(grid *Grid) *Grid {
	max := 0.0
	min := 0.0
    for _, row := range *grid {
		for _, cell := range row {
			min = math.Min(min, cell)
			max = math.Max(max, cell)
		}
	}
	rangeDist := max - min
	var newGrid Grid = make([][]float64, 0)
    for _, row := range *grid {
		newRow := make([]float64, 0)
    	for _, cell := range row {
    		newRow = append(newRow, (cell+math.Abs(min))/rangeDist)
		}
		newGrid = append(newGrid, newRow)
	}
	return &newGrid
}

func ApplyGradient(centerX, centerY float64, grid *Grid) *Grid {
	var newGrid Grid = make([][]float64, 0)
	lenX := float64(len(*grid)-1)
    for x, row := range *grid {
		fx := float64(x)
		lenY := float64(len(row)-1)
		newRow := make([]float64, 0)
    	for y, cell := range row {
			fy := float64(y)
			xgradient := 0.0
			if fx < centerX {
				xgradient = 1 - fx/centerX
			} else {
				xgradient = (fx-centerX)/(lenX-centerX)
			}
			ygradient := 0.0
			if fy < centerY {
				ygradient = 1 - fy/centerX
			} else {
				ygradient = (fy-centerY)/(lenY-centerY)
			}
			gradient := 1 - math.Max(xgradient, ygradient)
    		newRow = append(newRow, cell * gradient)
		}
		newGrid = append(newGrid, newRow)
	}
	return &newGrid
}

func Print(grid *Grid) {
	for _, row := range *grid {
    	for _, cell := range row {
			PrintBlock(cell)
		}
		fmt.Print("\n")
	}
}

func PrintGray(v float64) {
	fmt.Print(aurora.Gray(uint8(v * 24), character))
}

func PrintBlock(v float64) {
	if v > 0.7 {
		fmt.Print(aurora.Gray(22, character))
		return
	}
	if v > 0.3 {
		fmt.Print(aurora.Green(character))
		return
	}
	if v > 0.2 {
		fmt.Print(aurora.Yellow(character))
		return
	}
	if v > 0.1 {
		fmt.Print(aurora.BrightBlue(character))
		return
	}
	fmt.Print(aurora.Blue(character))
	return
}
