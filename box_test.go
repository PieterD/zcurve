package zcurve

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestBox_ComprehensiveSplit(t *testing.T) {
	for i := 0; i < 1000; i++ {
		x := rand.Intn(1000000000)
		y := rand.Intn(1000000000)
		w := rand.Intn(100)
		h := rand.Intn(100)
		b := Box{NewCoord(uint32(x), uint32(y)), NewCoord(uint32(x+w), uint32(y+h))}
		t.Run(fmt.Sprintf("%d-%d", b.TopLeft, b.BottomRight), func(t *testing.T) {
			t.Logf("box has area %d", b.Area())
			t.Logf("box has range %d", b.Range())
			boxes := b.ComprehensiveSplit(3, 10)
			if len(boxes) > 10 {
				t.Fatalf("invalid number of boxes %d, expected 10 or fewer for %#v", len(boxes), b)
			}
			t.Logf("split into %d boxes\n", len(boxes))
			boxesArea := area(boxes)
			if b.Area() != boxesArea {
				t.Fatalf("wrong area %d, expected %d for %#v", boxesArea, b.Area(), b)
			}
			oCoords := make(map[Coord]struct{})
			for y := b.TopLeft.Y(); y <= b.BottomRight.Y(); y++ {
				for x := b.TopLeft.X(); x <= b.BottomRight.X(); x++ {
					oCoords[NewCoord(x, y)] = struct{}{}
				}
			}
			splitCoords := make(map[Coord]struct{})
			for _, box := range boxes {
				for i := box.TopLeft; i <= box.BottomRight; i++ {
					splitCoords[i] = struct{}{}
				}
			}
			for coord := range oCoords {
				_, ok := splitCoords[coord]
				if !ok {
					t.Fatalf("missing coordinate %#v for %#v", coord, b)
				}
			}
			t.Logf("split range = %d", len(splitCoords))
		})
	}
}

func area(boxes []Box) int {
	sum := 0
	for _, box := range boxes {
		sum += box.Area()
	}
	return sum
}
