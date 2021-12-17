package main

import (
	"fmt"
	"testing"
)

func Test_PelletsInsert(t *testing.T) {

	pellets := NewPellets(5, 5)
	pellets.insert(Position{1, 1}, 2)
	pellets.insert(Position{5, 5}, 10)
	fmt.Printf("%+v", pellets)
}
