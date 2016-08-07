package foldproc

import (
	"math"
	"math/big"

	"github.com/IronhandedLayman/icfp-origami/objs"
)

func Centroid(po *objs.Polygon) objs.Point {
	ans := objs.Origin
	for _, pt := range *po {
		ans = ans.Add(pt)
	}
	return ans.Mult(big.NewRat(1, (int64)(len(*po))))
}

func OAngle(p1 objs.Point, p2 objs.Point) float64 {
	opt := p1.Sub(p2)
	ox, _ := opt.X.Float64()
	oy, _ := opt.Y.Float64()

	return math.Atan2(oy, ox)
}

func CycleSort(p *objs.Polygon) {
	cen := Centroid(p)
	// Instead of writing all that framework for sorting, I'm going to be slow here.
	// TODO: rewrite into Go's sorting framework
	n := len(*p)
	angs := make([]float64, n)
	for i := 0; i < n; i++ {
		angs[i] = OAngle(cen, (*p)[i])
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if angs[i] < angs[j] {
				(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
				angs[i], angs[j] = angs[j], angs[i]
			}
		}
	}
}

func TriangleArea(o, p1, p2 objs.Point) float64 {
	d1 := p1.Sub(o)
	d2 := p2.Sub(o)
	d1x, _ := d1.X.Float64()
	d1y, _ := d1.Y.Float64()
	d2x, _ := d2.X.Float64()
	d2y, _ := d2.Y.Float64()
	return math.Abs(.5 * (d1x*d2y - d1y*d2x))
}

func SimplePolyArea(p *objs.Polygon) float64 {
	CycleSort(p)
	cen := Centroid(p)
	var ans float64 = 0
	n := len(*p)
	for i := 0; i <= n; i++ {
		ans += TriangleArea(cen, (*p)[i], (*p)[(i+1)%n])
	}
	return ans
}
