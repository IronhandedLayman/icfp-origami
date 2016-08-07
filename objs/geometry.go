package objs

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

type Point struct {
	X *big.Rat
	Y *big.Rat
}

var (
	Origin        = Point{big.NewRat(0, 1), big.NewRat(0, 1)}
	NoLine        = Line{Origin, Origin}
	NoProblem     = Problem{nil, nil}
	InitialSquare = Polygon{
		Point{big.NewRat(0, 1), big.NewRat(0, 1)},
		Point{big.NewRat(1, 1), big.NewRat(0, 1)},
		Point{big.NewRat(1, 1), big.NewRat(1, 1)},
		Point{big.NewRat(0, 1), big.NewRat(1, 1)},
	}
	InitialSkeleton = Skeleton{
		Line{Point{big.NewRat(0, 1), big.NewRat(0, 1)}, Point{big.NewRat(1, 1), big.NewRat(0, 1)}},
		Line{Point{big.NewRat(1, 1), big.NewRat(0, 1)}, Point{big.NewRat(1, 1), big.NewRat(1, 1)}},
		Line{Point{big.NewRat(1, 1), big.NewRat(1, 1)}, Point{big.NewRat(0, 1), big.NewRat(1, 1)}},
		Line{Point{big.NewRat(0, 1), big.NewRat(1, 1)}, Point{big.NewRat(0, 1), big.NewRat(0, 1)}},
	}
)

func (p Point) String() string {
	return fmt.Sprintf("%s,%s", p.X, p.Y)
}

func ParsePoint(instr string) (Point, error) {
	parts := strings.Split(instr, ",")
	if len(parts) != 2 {
		return Origin, fmt.Errorf("Malformatted point")
	}
	xcoord, xerr := (&big.Rat{}).SetString(parts[0])
	if !xerr {
		return Origin, fmt.Errorf("Illegal x-coordinate: %v", parts[0])
	}

	ycoord, yerr := (&big.Rat{}).SetString(parts[1])
	if !yerr {
		return Origin, fmt.Errorf("Illegal y-coordinate: %v", parts[1])
	}

	return Point{xcoord, ycoord}, nil
}

func (p Point) Add(oth Point) Point {
	nx := big.NewRat(0, 1)
	ny := big.NewRat(0, 1)
	return Point{
		nx.Add(p.X, oth.X),
		ny.Add(p.Y, oth.Y),
	}
}

func (p Point) Sub(oth Point) Point {
	nx := big.NewRat(0, 1)
	ny := big.NewRat(0, 1)
	return Point{
		nx.Sub(p.X, oth.X),
		ny.Sub(p.Y, oth.Y),
	}
}

func (p Point) Mult(scr *big.Rat) Point {
	nx := big.NewRat(0, 1)
	ny := big.NewRat(0, 1)
	return Point{
		nx.Mul(p.X, scr),
		ny.Mul(p.Y, scr),
	}
}

type Polygon []Point
type Line [2]Point
type Skeleton []Line

func (ps Polygon) String() string {
	ans := fmt.Sprintf("%d", len(ps))
	for _, p := range ps {
		ans += fmt.Sprintf(" %s", p)
	}
	return ans
}

func ParsePolygon(instr []string) (Polygon, error) {
	np, err := strconv.ParseInt(instr[0], 10, 32)
	if err != nil || int(np) > (len(instr)-1) {
		return nil, fmt.Errorf("Incorrect counting index")
	}
	ans := make([]Point, np)
	var perr error
	for i := 1; i <= int(np); i++ {
		ans[i-1], perr = ParsePoint(instr[i])
		if perr != nil {
			return nil, fmt.Errorf("Error parsing index #%d: %v", i-1, perr)
		}
	}
	return ans, nil
}

func ParseLine(instr string) (Line, error) {
	ps := strings.Split(instr, " ")
	if len(ps) != 2 {
		return NoLine, fmt.Errorf("Misplaced spacing")
	}

	p1, perr := ParsePoint(ps[0])
	if perr != nil {
		return NoLine, fmt.Errorf("Point 1 malformed: %v", perr)
	}
	p2, perr := ParsePoint(ps[1])
	if perr != nil {
		return NoLine, fmt.Errorf("Point 2 malformed: %v", perr)
	}
	return [2]Point{p1, p2}, nil
}

func ParseSkeleton(instr []string) (Skeleton, error) {
	np, err := strconv.ParseInt(instr[0], 10, 32)
	if err != nil || int(np) > (len(instr)-1) {
		return nil, fmt.Errorf("Incorrect counting index")
	}
	ans := make([]Line, np)
	var perr error
	for i := 1; i <= int(np); i++ {
		ans[i-1], perr = ParseLine(instr[i])
		if perr != nil {
			return nil, fmt.Errorf("Error parsing index #%d", i-1)
		}
	}
	return ans, nil
}

type Problem struct {
	Silhouette   []Polygon
	ProbSkeleton Skeleton
}

func ParseProblem(instr string) (Problem, error) {
	lines := strings.Split(instr, "\n")
	np, err := strconv.ParseInt(lines[0], 10, 32)
	if err != nil {
		return NoProblem, fmt.Errorf("Couldn't parse polygon count: %v", err)
	}
	lineno := 1
	ans := Problem{}
	ans.Silhouette = make([]Polygon, np)
	for i := 0; i < (int)(np); i++ {
		ans.Silhouette[i], err = ParsePolygon(lines[lineno:])

		if err != nil {
			return NoProblem, fmt.Errorf("Couldn't parse polygon #%d: %v", i, err)
		}
		lineno += len(ans.Silhouette[i]) + 1
	}
	ans.ProbSkeleton, err = ParseSkeleton(lines[lineno:])
	if err != nil {
		return NoProblem, fmt.Errorf("Couldn't parse skeleton: %v", err)
	}
	return ans, nil
}

type Solution struct {
	InitialPoints Polygon
	Facets        []Polygon
	FinalPoints   Polygon
}

func BlankSlate() (Problem, Solution) {
	return Problem{[]Polygon{InitialSquare}, InitialSkeleton}, Solution{InitialSquare, []Polygon{InitialSquare}, InitialSquare}
}
