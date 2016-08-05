package objs

import (
	"fmt"
	"strconv"
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

type Points []Point

func (ps Points) String() string {
	ans := fmt.Sprintf("%d", len(ps))
	for _, p := range ps {
		ans += fmt.Sprintf(" %s", p)
	}
	return ans
}

func ParsePoints(instr string) (Points, error) {
	parts := strings.Split(" ", instr)

	np, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil || int(np) != (len(parts)-1) {
		return nil, fmt.Errorf("Incorrect counting index")
	}
	ans := make([]Point, np)
	var perr error
	for i := 1; i <= int(np); i++ {
		ans[i-1], perr = ParsePoint(parts[i])
		if perr != nil {
			return nil, fmt.Errorf("Error parsing index #%d", i-1)
		}
	}
	return ans, nil
}
