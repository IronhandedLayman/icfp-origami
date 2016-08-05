package objs

import (
	"fmt"
	"strings"
)

type Point struct {
	X Ratio
	Y Ratio
}

var Origin = Point{ZeroRatio, ZeroRatio}

func (p Point) String() string {
	return fmt.Sprintf("%s,%s", p.X, p.Y)
}

func ParsePoint(instr string) (Point, error) {
	parts := strings.Split(",", instr)
	if len(parts) != 2 {
		return Origin, fmt.Errorf("Malformatted point")
	}
	xcoord, xerr := ParseRatio(parts[0])
	if xerr != nil {
		return Origin, fmt.Errorf("Illegal x-coordinate")
	}

	ycoord, yerr := ParseRatio(parts[1])
	if yerr != nil {
		return Origin, fmt.Errorf("Illegal y-coordinate")
	}

	return Point{xcoord, ycoord}, nil
}
