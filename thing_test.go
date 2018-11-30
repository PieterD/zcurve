package zcurve

import (
	"fmt"
	"testing"
)

func TestQuick(t *testing.T) {
	b := Box{2279, 49978}
	_, _, b1, b2 := b.Split()
	fmt.Printf("b  = %#v\n", b)
	fmt.Printf("b1 = %#v\n", b1)
	fmt.Printf("b2 = %#v\n", b2)
}
