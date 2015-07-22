package main

import (
	"fmt"
)

type Vertex struct {
	X, Y int
}

func main() {

	p := Vertex{1, 2}

	v := p

	fmt.Println(v)

	k := p

	fmt.Println(k)

}
