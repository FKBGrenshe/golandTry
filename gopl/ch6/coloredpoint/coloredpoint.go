package main

import (
	"fmt"
	"image/color"
	"math"
	"sync"
	"time"
)

var (
	mu      sync.Mutex
	mapping = make(map[string]string)
)

func Lookup(key string) string {
	mu.Lock()
	v := mapping[key]
	mu.Unlock()
	return v
}

func goroutinePool() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Printf("%d: start looking up\n", i)
			Lookup("key")
			fmt.Printf("%d: end looking up\n", i)
		}(i)
	}
}

type Point struct {
	X, Y float64
}

func (q *Point) Distance(p Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

type ColoredPoint struct {
	Point
	Color color.RGBA
}

func (cp ColoredPoint) Distance(p Point) float64 {
	return cp.Point.Distance(p)
}

func main() {
	var cp ColoredPoint
	cp.X = 1
	fmt.Println(cp.Point.X)
	cp.Point.Y = 2
	fmt.Println(cp.Y)

	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}

	var p = ColoredPoint{Point{1, 1}, red}
	var q = ColoredPoint{Point{5, 4}, blue}
	fmt.Println(p.Distance(q.Point)) // "5"

	goroutinePool()
	time.Sleep(10 * time.Second)
}
