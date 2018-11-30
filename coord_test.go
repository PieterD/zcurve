package zcurve

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCoord_HighestLevel(t *testing.T) {
	for _, test := range []struct {
		c Coord
		l int
	}{
		{0, 0},
		{1, 0},
		{2, 0},
		{3, 0},
		{4, 1},
		{5, 1},
		{6, 1},
		{7, 1},
		{8, 1},
		{15, 1},
		{16, 2},
		{31, 2},
		{32, 2},
		{63, 2},
		{64, 3},
		{127, 3},
		{128, 3},
		{255, 3},
		{256, 4},
		{511, 4},
		{512, 4},
		{1023, 4},
		{1024, 5},

		{628, 4},
		{2000, 5},
		{2449, 5},
	} {
		t.Run(fmt.Sprintf("%d", test.c), func(t *testing.T) {
			want := test.l
			got := test.c.HighestLevel()
			if want != got {
				t.Fatalf("expected %d, got %d", want, got)
			}
		})
	}
}

func TestCoord_Quadrant(t *testing.T) {
	for _, test := range []struct {
		c Coord
		q []Quadrant
	}{
		{0, []Quadrant{0}},
		{1, []Quadrant{1}},
		{2, []Quadrant{2}},
		{3, []Quadrant{3}},
		{4, []Quadrant{0, 1}},
		{5, []Quadrant{1, 1}},
		{6, []Quadrant{2, 1}},
		{7, []Quadrant{3, 1}},
		{8, []Quadrant{0, 2}},
		{9, []Quadrant{1, 2}},

		{10, []Quadrant{2, 2}},
		{11, []Quadrant{3, 2}},
		{12, []Quadrant{0, 3}},
		{13, []Quadrant{1, 3}},
		{14, []Quadrant{2, 3}},
		{15, []Quadrant{3, 3}},
		{16, []Quadrant{0, 0, 1}},
		{17, []Quadrant{1, 0, 1}},
		{44, []Quadrant{0, 3, 2}},
		{63, []Quadrant{3, 3, 3}},
		{64, []Quadrant{0, 0, 0, 1}},
	} {
		t.Run(fmt.Sprintf("%d", test.c), func(t *testing.T) {
			want := test.q
			got := listQuadrants(test.c)
			if !reflect.DeepEqual(want, got) {
				t.Fatalf("expected %v, got %v", want, got)
			}
		})
	}
}

func listQuadrants(c Coord) []Quadrant {
	highest := c.HighestLevel()
	levels := make([]Quadrant, highest+1)
	for level := range levels {
		levels[level] = c.Quadrant(Level(level))
	}
	return levels
}

