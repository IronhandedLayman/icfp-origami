package objs

import (
	"fmt"
	"strconv"
	"strings"
)

type Ratio struct {
	Num int64
	Den int64
}

var ZeroRatio = Ratio{0, 1}

func NewRatio(n int64, d int64) Ratio {
	if d == 0 {
		d = 1
	}
	return Ratio{n, d}.Reduced()
}

func (r Ratio) String() string {
	return fmt.Sprintf("%d/%d", r.Num, r.Den)
}

func (r Ratio) ToFloat() float64 {
	return float64(r.Num) / float64(r.Den)
}

func ParseRatio(instr string) (Ratio, error) {
	parts := strings.Split("/", instr)
	if len(parts) != 2 {
		return ZeroRatio, fmt.Errorf("Improper division placement")
	}

	npart, nerr := strconv.ParseInt(parts[0], 10, 64)
	if nerr != nil {
		return ZeroRatio, fmt.Errorf("Improper numerator")
	}

	dpart, derr := strconv.ParseInt(parts[1], 10, 64)
	if derr != nil {
		return ZeroRatio, fmt.Errorf("Improper denominator")
	}
	return Ratio{npart, dpart}.Reduced(), nil
}

func gcd(a int64, b int64) int64 {
	if a == 0 || b == 0 {
		return a + b
	}
	if a == 1 || b == 1 {
		return a * b
	}
	if b > a {
		return gcd(b, a)
	}
	return gcd(b, a%b)
}

func (r Ratio) Reduced() Ratio {
	g := gcd(r.Num, r.Den)
	if g == 0 {
		g = 1
	}
	if r.Num == 0 {
		return Ratio{Num: 0, Den: 1}
	}
	return Ratio{Num: r.Num / g, Den: r.Den / g}
}

func (r Ratio) Add(oth Ratio) Ratio {
	return Ratio{r.Num*oth.Den + r.Den*oth.Num, r.Den * oth.Den}.Reduced()
}

func (r Ratio) Multiply(oth Ratio) Ratio {
	return Ratio{r.Num * oth.Num, r.Den * oth.Den}.Reduced()
}

func (r Ratio) EqualTo(oth Ratio) bool {
	if r.Num == 0 && oth.Num == 0 {
		return true
	}
	return r.Num == oth.Num && r.Den == oth.Den
}

func (r Ratio) GreaterThan(oth Ratio) bool {
	return r.Num*oth.Den > r.Den*oth.Num
}
