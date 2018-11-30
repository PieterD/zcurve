package zcurve

import "fmt"

type Box struct {
	TopLeft     Coord
	BottomRight Coord
}

func (b Box) GoString() string {
	return fmt.Sprintf("%#v-%#v (area %d)", b.TopLeft, b.BottomRight, b.Area())
}

func (b Box) Height() int {
	return int(b.BottomRight.Y()-b.TopLeft.Y()) + 1
}

func (b Box) Width() int {
	return int(b.BottomRight.X()-b.TopLeft.X()) + 1
}

func (b Box) Area() int {
	return b.Width() * b.Height()
}

func (b Box) Range() int {
	return int(b.BottomRight - b.TopLeft + 1)
}

func (b Box) Split() (bool, Level, Box, Box) {
	highestLevel := b.BottomRight.HighestLevel()
	for level := Level(highestLevel); level > 0; level-- {
		quadrant := b.BottomRight.Quadrant(level)
		otherQuadrant := b.TopLeft.Quadrant(level)
		if quadrant == otherQuadrant {
			continue
		}

		xDiff, yDiff := quadrant.Diff(otherQuadrant)
		tl := b.TopLeft.Subtree(level)
		br := b.BottomRight.Subtree(level)
		var litmax, bigmin Coord
		if yDiff {
			litmaxSubtree := tl
			litmaxSubtree.Subtree = NewCoord(br.Subtree.X(), max(level))
			litmaxSubtree.Quadrant = 0x0 | (br.Quadrant & 0x1)
			litmax = litmaxSubtree.Combine()
			bigminSubtree := br
			bigminSubtree.Subtree = NewCoord(tl.Subtree.X(), 0)
			bigminSubtree.Quadrant = 0x2 | (tl.Quadrant & 0x1)
			bigmin = bigminSubtree.Combine()
		}
		if xDiff {
			litmaxSubtree := tl
			litmaxSubtree.Subtree = NewCoord(max(level), br.Subtree.Y())
			litmaxSubtree.Quadrant = 0x0 | (br.Quadrant & 0x2)
			litmax2 := litmaxSubtree.Combine()
			bigminSubtree := br
			bigminSubtree.Subtree = NewCoord(0, tl.Subtree.Y())
			bigminSubtree.Quadrant = 0x1 | (tl.Quadrant & 0x2)
			bigmin2 := bigminSubtree.Combine()
			if !yDiff || bigmin2-litmax2 > bigmin-litmax {
				litmax = litmax2
				bigmin = bigmin2
			}
		}
		if litmax+1 == bigmin {
			return false, 0, Box{}, Box{}
		}
		return true, level, Box{b.TopLeft, litmax}, Box{bigmin, b.BottomRight}
	}
	return false, 0, Box{}, Box{}
}

func (b Box) ComprehensiveSplit(lowLevel Level, maxBoxes int) []Box {
	if lowLevel < 0 {
		lowLevel = 0
	}
	if maxBoxes == 0 {
		maxBoxes = 1
	}
	var boxes = []Box{b}
	boxNum := 1

	done := false

	for !done {
		var expanded = make([]Box, 0, len(boxes)*2)
		for _, b := range boxes {
			if maxBoxes > 0 && boxNum >= maxBoxes {
				expanded = append(expanded, b)
				done = true
				continue
			}
			ok, level, b1, b2 := b.Split()
			if !ok || level <= lowLevel {
				expanded = append(expanded, b)
				continue
			}
			expanded = append(expanded, b1, b2)
			boxNum++
		}
		if len(boxes) == len(expanded) {
			done = true
		}
		boxes = expanded
	}
	return boxes
}
