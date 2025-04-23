package standard

import (
	"image/color"
	"math"
	"sort"
)

// ColorStop represents a single color stop in a gradient.
// T defines the position along the gradient line, ranging from 0.0 to 1.0.
// Color defines the color at this stop.
type ColorStop struct {
	T     float64 // from 0.0 to 1.0.
	Color color.RGBA
}

// LinearGradient defines a linear gradient with angle and color stops.
// The gradient progresses in the direction of the given angle (in degrees).
// Angle is interpreted as: 0 - right, 90 - up, 180 - left, 270 - down.
type LinearGradient struct {
	Stops []ColorStop // Ordered list of color stops along the gradient
	Angle float64     // Gradient angle in degrees
}

// NewGradient creates a new LinearGradient with the specified angle and color stops.
// The stops are sorted in ascending order of T.
func NewGradient(angle float64, stops ...ColorStop) *LinearGradient {
	sort.Slice(stops, func(i, j int) bool {
		return stops[i].T < stops[j].T
	})
	return &LinearGradient{Stops: stops, Angle: angle}

}

// At computes the interpolated color at a given pixel position (x, y)
// within a bounding rectangle defined by width (w) and height (h).
// The interpolation is based on the projection of the point onto
// the gradient direction vector
func (g *LinearGradient) At(x, y float64, w, h int) color.RGBA {
	rad := g.Angle * math.Pi / 180.0
	dx := math.Cos(rad)
	dy := -math.Sin(rad) // invert Y-axis as screen coordinates grow downward

	// Project all four corners of the bounding rectangle
	corners := [4][2]float64{
		{0, 0},
		{float64(w), 0},
		{0, float64(h)},
		{float64(w), float64(h)},
	}
	minT, maxT := math.Inf(1), math.Inf(-1)
	for _, p := range corners {
		t := p[0]*dx + p[1]*dy
		if t < minT {
			minT = t
		}
		if t > maxT {
			maxT = t
		}
	}

	// Project the target point and normalize it
	dot := x*dx + y*dy
	t := (dot - minT) / (maxT - minT)
	t = math.Max(0, math.Min(1, t))

	if len(g.Stops) == 0 {
		return color_BLACK
	}
	if t <= g.Stops[0].T {
		return g.Stops[0].Color
	}
	if t >= g.Stops[len(g.Stops)-1].T {
		return g.Stops[len(g.Stops)-1].Color
	}

	// Interpolate color between two enclosing stops
	for i := 0; i < len(g.Stops)-1; i++ {
		a, b := g.Stops[i], g.Stops[i+1]
		if t >= a.T && t <= b.T {
			nt := (t - a.T) / (b.T - a.T)
			return interpolateColor(a.Color, b.Color, nt)
		}
	}
	return color_BLACK
}

// interpolateColor linearly interpolates between two colors based on parameter t (0 to 1).
func interpolateColor(start, end color.Color, t float64) color.RGBA {
	sr, sg, sb, sa := start.RGBA()
	er, eg, eb, ea := end.RGBA()

	r := uint8(float64(sr>>8)*(1-t) + float64(er>>8)*t)
	g := uint8(float64(sg>>8)*(1-t) + float64(eg>>8)*t)
	b := uint8(float64(sb>>8)*(1-t) + float64(eb>>8)*t)
	a := uint8(float64(sa>>8)*(1-t) + float64(ea>>8)*t)

	return color.RGBA{R: r, G: g, B: b, A: a}
}
