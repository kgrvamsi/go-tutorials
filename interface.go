package main

import (
	"fmt"
)

type Shaper interface {
	Area() int
}

type Rectangle struct {
	length, width int
}

func (r Rectangle) Area() int {

	return r.length * r.width
}

func main() {

	s := Rectangle{length: 4, width: 5}

	fmt.Println(s)

	fmt.Println("this is with just the function ", s.Area())

	r := Shaper(s)

	fmt.Println("This is with interface ", r.Area())

}
