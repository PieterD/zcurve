package zcurve

import (
	"fmt"
	"math/bits"
)

type Quadrant byte

func (q Quadrant) Diff(q2 Quadrant) (x, y bool) {
	c := q ^ q2
	if c&1 == 1 {
		x = true
	}
	if c&2 == 2 {
		y = true
	}
	return x, y
}

type Level byte

type Coord uint64

func (c Coord) String() string {
	return c.GoString()
}

func (c Coord) GoString() string {
	return fmt.Sprintf("%d(%d,%d)", c, c.X(), c.Y())
}

func NewCoord(x uint32, y uint32) Coord {
	return Coord(interleave(x, y))
}

func (c Coord) X() uint32 {
	return unsparse(uint64(c))
}

func (c Coord) Y() uint32 {
	return unsparse(uint64(c >> 1))
}

func (c Coord) separate() (x, y Coord) {
	x = c & 0x5555555555555555
	y = c & 0x5555555555555555 << 1
	return x, y
}

func (c Coord) Quadrant(level Level) Quadrant {
	return Quadrant(uint64(c) >> (level * 2) & 0x3)
}

func (c Coord) HighestLevel() int {
	return (64 - bits.LeadingZeros64(uint64(c)) - 1) / 2
}

func max(level Level) uint32 {
	mask := ^(^uint32(0) << level)
	return 0xFFFFFFFF & mask
}

type Subtree struct {
	Level    Level
	Parent   Coord
	Quadrant Quadrant
	Subtree  Coord
}

func (c Coord) Subtree(level Level) Subtree {
	parentMask := ^Coord(0) << (level * 2)
	childMask := ^parentMask
	parentMask <<= 2

	return Subtree{
		Level:    level,
		Parent:   c & parentMask,
		Quadrant: c.Quadrant(level),
		Subtree:  c & childMask,
	}
}

func (s Subtree) Combine() Coord {
	return s.Parent | s.Subtree | (Coord(s.Quadrant) << (s.Level * 2))
}
