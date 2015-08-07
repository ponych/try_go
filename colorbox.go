package main

import "fmt"

const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type Color byte

type Box struct {
	width, height, depth float64
	color Color
}

type BoxList []Box // a slice of boxes

func (b Box)Volume() float64 {
	return b.width * b.height * b.depth
}

func (b *Box)SetColor(c Color) {
	b.color = c
}

func (c Color)String() string {
	strings := []string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return strings[c]
}

func (bl BoxList)BiggeestColor()(k Color) {
	v := 0
	k = Color(WHITE)

	for _, b := range bl {
		if bv := b.Volume(); bv > v {
			v = bv
			k = b.color
		}
	}

	return
}

func main() {
	boxes := BoxList {
		Box{ 4, ,4 ,4 , RED },
		Box{ 10 ,10, 1, YELLOW},
		Box{ 1  ,1 , 20, BLACK} ,
		Box{ 10, 10, BLUE} ,
		Box{ 10, 30, 1, WHITE} ,
		Box{ 20, 20 ,20, YELLOW},
	}

	fmt.Printf("We have %d boxes in our set\n", len(boxes))
	fmt.Println("The volume of the first box is", boxes[0].Volume())
	fmt.Println("the color of last box", boxes[len(boxes)-1].color.String())
	fmt.Println("The Biggest one is", boxes.BiggeestColor().String())
}
