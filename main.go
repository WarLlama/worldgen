package main
import (
	"fmt"
	"github.com/aquilax/go-perlin"
	"github.com/logrusorgru/aurora"
	"math"
	"math/rand"
)

const (
	width = 10
	height = 10
	character = "\u2588"
)
func main() {
	p := perlin.NewPerlin(2, 2, 10, rand.Int63())
	maxRadius := math.Hypot(width, height)
	for x := 0.0; x < width; x+=0.1 {
		for y := 0.0; y < height; y+=0.1 {
			centerY := 0.5 * height
			centerX := 0.5 * width
			dx := x-centerX
			dy := y-centerY
			radius := math.Hypot(dx, dy)
			max := 1 - radius/(2 * maxRadius)
			min := 0.4 - radius/(2 *maxRadius)
			PrintBlock(clamp(p.Noise2D(x, y), max, min))
		}
		fmt.Print("\n")
    }
}

func clamp(v, max, min float64) float64 {
	if v > max {
		return max
	}
	if v < min {
		return min
	}
	return v
}

func PrintBlock(v float64) {
	if v > 0.8 {
		fmt.Print(aurora.Gray(22, character))
		return
	}
	if v > 0.6 {
		fmt.Print(aurora.Green(character))
		return
	}
	if v > 0.3 {
		fmt.Print(aurora.Yellow(character))
		return
	}
	if v > 0.2 {
		fmt.Print(aurora.BrightBlue(character))
		return
	}
		fmt.Print(aurora.Blue(character))
		return
}
