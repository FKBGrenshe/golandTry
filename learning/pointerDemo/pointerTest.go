package main

import "fmt"

func main() {
	var a int = 20
	var p = &a

	fmt.Println(a)
	fmt.Println(p)
	fmt.Println(*p)

	fmt.Println("====================")

	r := Point{1, 2}
	fmt.Println("r: " + fmt.Sprintf("%v", r))
	fmt.Println("&r: " + fmt.Sprintf("%p", &r))

	rAddress := &r
	fmt.Println("rAddress: " + fmt.Sprintf("%p", rAddress))

	r.move(1, 1)
	fmt.Println("r after move: " + fmt.Sprintf("%v", r))
	(&r).move(-1, -1)
	fmt.Println("r after move: " + fmt.Sprintf("%v", r))
}

type Point struct {
	x int
	y int
}

func (p *Point) move(dx, dy int) {
	p.x += dx
	p.y += dy
}

// type pAddress *Point

// func (p pAddress) move(dx, dy int) {
// 	p.x += dx
// 	p.y += dy
// }
