package main

import (
	"fmt"
	"math"

	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
)

type LiquidShape struct {
	p float64
}

func (s *LiquidShape) DrawFinder(ctx *standard.DrawContext) {
	w, h := ctx.Edge()
	x, y := ctx.UpperLeft()
	c := ctx.Color()

	ctx.SetColor(c)

	mask := ctx.Neighbours()
	switch {
	case mask == (NSelf | NBot | NLeft):
		//ctx.SetRGB(1, 0, 0)
		ctx.MoveTo(x, y)
		ctx.QuadraticTo(x+float64(w), y, x+float64(w), y+float64(h))
		ctx.LineTo(x, y+float64(h))
		ctx.ClosePath()
	case mask == (NSelf | NBot | NLeft | NBotLeft):
		//ctx.SetRGB(1, 0, 0)
		ctx.MoveTo(x, y)
		ctx.QuadraticTo(x+float64(w), y, x+float64(w), y+float64(h))
		ctx.LineTo(x, y+float64(h))
		ctx.ClosePath()

	case mask == (NSelf | NBot | NRight):
		//ctx.SetRGB(1, 0, 0)
		ctx.MoveTo(x, y+float64(h))
		ctx.QuadraticTo(x, y, x+float64(w), y)
		ctx.LineTo(x+float64(w), y+float64(h))
		ctx.ClosePath()
	case mask == (NSelf | NBot | NRight | NBotRight):
		//ctx.SetRGB(1, 0, 0)
		ctx.MoveTo(x, y+float64(h))
		ctx.QuadraticTo(x, y, x+float64(w), y)
		ctx.LineTo(x+float64(w), y+float64(h))
		ctx.ClosePath()

	case mask == (NSelf | NTop | NRight):
		//ctx.SetRGB(1, 0, 0)
		ctx.MoveTo(x, y)
		ctx.QuadraticTo(x, y+float64(h), x+float64(w), y+float64(h))
		ctx.LineTo(x+float64(w), y)
		ctx.ClosePath()
	case mask == (NSelf | NTop | NRight | NTopRight):
		//ctx.SetRGB(1, 0, 0)
		ctx.MoveTo(x, y)
		ctx.QuadraticTo(x, y+float64(h), x+float64(w), y+float64(h))
		ctx.LineTo(x+float64(w), y)
		ctx.ClosePath()

	case mask == (NSelf | NTop | NLeft):
		//ctx.SetRGB(1, 0, 0)
		ctx.MoveTo(x, y+float64(h))
		ctx.QuadraticTo(x+float64(w), y+float64(h), x+float64(w), y)
		ctx.LineTo(x, y)
		ctx.ClosePath()
	case mask == (NSelf | NTop | NLeft | NTopLeft):
		//ctx.SetRGB(1, 0, 0)
		ctx.MoveTo(x, y+float64(h))
		ctx.QuadraticTo(x+float64(w), y+float64(h), x+float64(w), y)
		ctx.LineTo(x, y)
		ctx.ClosePath()
	default:
		ctx.DrawRectangle(x, y, float64(w), float64(h))
	}

	ctx.Fill()
}

func newShape(p float64) standard.IShape {
	if p < 0 || p > 1 {
		p = 0
	}
	return &LiquidShape{
		p: p,
	}
}

func (s *LiquidShape) Draw(ctx *standard.DrawContext) {
	w, h := ctx.Edge()
	x, y := ctx.UpperLeft()
	c := ctx.Color()

	l := float64(w) * s.p
	_ = l

	// choose a proper radius values
	radius := w / 2
	r2 := h / 2
	if r2 <= radius {
		radius = r2
	}

	r := float64(radius)
	_ = r

	cx, cy := x+float64(w)/2.0, y+float64(h)/2.0 // get center point

	ctx.SetColor(c)

	AngTopRight := func(ctx *standard.DrawContext) {
		//ctx.Push()
		ctx.MoveTo(cx, cy+r)
		ctx.LineTo(cx-r, cy)
		ctx.LineTo(cx-r, y)
		ctx.LineTo(cx+r, y-l) // todo: % вставить
		ctx.QuadraticTo(cx+r, cy-r, x+float64(w)+l, cy-r)
		ctx.LineTo(x+float64(w), cy+r)
		ctx.ClosePath()
		//ctx.Pop()
	}
	_ = AngTopRight

	AngTopLeft := func(ctx *standard.DrawContext) {
		//ctx.Push()
		ctx.MoveTo(cx, cy+r)
		ctx.LineTo(cx+r, cy)
		ctx.LineTo(cx+r, y-l)
		ctx.LineTo(cx-r, y-l)
		ctx.QuadraticTo(cx-r, cy-r, x-l, cy-r)
		ctx.LineTo(x-l, cy+r)
		ctx.ClosePath()
		//ctx.Pop()
	}
	_ = AngTopLeft

	AngBotLeft := func(ctx *standard.DrawContext) {
		//ctx.Push()
		ctx.MoveTo(cx, cy-r)
		ctx.LineTo(cx+r, cy)
		ctx.LineTo(cx+r, y+float64(h)+l)
		ctx.LineTo(cx-r, y+float64(h)+l)
		ctx.QuadraticTo(cx-r, cy+r, x-l, cy+r)
		ctx.LineTo(x-l, cy-r)
		ctx.ClosePath()
		//ctx.Pop()
	}
	_ = AngBotLeft

	AngBotRight := func(ctx *standard.DrawContext) {
		ctx.MoveTo(cx, cy-r)
		ctx.LineTo(cx-r, cy)
		ctx.LineTo(cx-r, y+float64(h)+l)
		ctx.LineTo(cx+r, y+float64(h)+l)
		ctx.QuadraticTo(cx+r, cy+r, x+float64(w)+l, cy+r)
		ctx.LineTo(x+float64(w), cy-r)
		ctx.ClosePath()
		ctx.Fill()
	}
	_ = AngBotRight

	mask := ctx.Neighbours()
	switch {
	// top right
	//case (mask&(NTop|NRight|NSelf) == (NTop | NRight | NSelf)) && (mask&NTopRight == 0):
	//	//ctx.SetRGB(0, 1, 0)
	//	ctx.DrawCircle(cx, cy, 2)
	//	AngTopRight(ctx)
	//	ctx.Fill()
	// top left
	//case (mask&(NTop|NLeft|NSelf) == (NTop | NLeft | NSelf)) && (mask&NTopLeft == 0):
	//	//ctx.SetRGB(1, 0, 0)
	//	ctx.DrawCircle(cx, cy, 2)
	//	AngTopLeft(ctx)
	//	ctx.Fill()
	// bot left
	//case (mask&(NBot|NLeft|NSelf) == (NBot | NLeft | NSelf)) && (mask&NBotLeft == 0):
	//	//ctx.SetRGB(0, 1, 0)
	//	ctx.DrawCircle(cx, cy, 1)
	//	AngBotLeft(ctx)
	//	ctx.Fill()
	// bot right
	//case (mask&(NBot|NRight|NSelf) == (NBot | NRight | NSelf)) && (mask&NBotRight == 0):
	//case has(mask, NBot|NRight|NSelf) && mask&NBotRight == 0:
	//	ctx.SetRGB(1, 0, 0)
	//	ctx.DrawCircle(cx, cy, 4)
	//	//AngBotRight(ctx)
	//	ctx.Fill()
	case mask == NRight|NSelf:
		//ctx.SetRGB(0, 1, 0)
		ctx.DrawRectangle(cx, cy-float64(radius), float64(w)/2.0, 2*float64(radius))
		ctx.Fill()
	case mask == NTop|NSelf:
		//ctx.SetRGB(0, 1, 0)
		ctx.DrawRectangle(cx-float64(radius), y, 2*float64(radius), float64(h)/2.0)
		ctx.Fill()
	case mask == NLeft|NSelf:
		//ctx.SetRGB(0, 1, 0)
		ctx.DrawRectangle(x, cy-float64(radius), float64(w)/2.0, 2*float64(radius))
		ctx.Fill()
	case mask == NBot|NSelf:
		//ctx.SetRGB(0, 1, 0)
		ctx.DrawRectangle(cx-float64(radius), y+float64(h)/2, 2*float64(radius), float64(h)/2.0)
		ctx.Fill()
	case mask == NLeft|NSelf|NRight:
		//ctx.SetRGB(0, 0, 1)
		ctx.DrawRectangle(x-float64(w)/2, cy-r, 2*float64(w), 2*r)
		ctx.Fill()
	default:
		fmt.Printf("Unsupported mask %v\n", mask)
	}

	if mask&(NLeft|NSelf|NRight) == (NLeft | NSelf | NRight) {
		ctx.DrawRectangle(x-float64(w)/2, cy-r, 2*float64(w), 2*r)
		ctx.Fill()
	}
	if mask&(NTop|NSelf|NBot) == (NTop | NSelf | NBot) {
		ctx.DrawRectangle(cx-r, y-float64(h)/2, 2*r, 2*float64(h))
		ctx.Fill()
	}
	if has(mask, NLeft|NSelf) {
		//ctx.SetRGB(1, 0, 1)
		ctx.DrawRectangle(x, cy-float64(radius), float64(w)/2.0, 2*float64(radius))
		ctx.Fill()
	}
	if has(mask, NSelf|NRight) {
		//ctx.SetRGB(1, 0, 1)
		ctx.DrawRectangle(cx, cy-float64(radius), float64(w)/2.0, 2*float64(radius))
		ctx.Fill()
	}
	if has(mask, NSelf|NTop) {
		//ctx.SetRGB(1, 0, 1)
		ctx.DrawRectangle(cx-float64(radius), y, 2*float64(radius), float64(h)/2.0)
		ctx.Fill()
	}
	if has(mask, NSelf|NBot) {
		//ctx.SetRGB(1, 0, 1)
		ctx.DrawRectangle(cx-float64(radius), y+float64(h)/2, 2*float64(radius), float64(h)/2.0)
		ctx.Fill()
	}

	if has(mask, NBot|NRight|NSelf) && mask&NBotRight == 0 {
		//ctx.SetRGB(1, 0, 0)
		//ctx.DrawCircle(cx, cy, 4)
		AngBotRight(ctx)
		ctx.Fill()
	}

	if (mask&(NBot|NLeft|NSelf) == (NBot | NLeft | NSelf)) && (mask&NBotLeft == 0) {
		//ctx.SetRGB(0, 1, 0)
		ctx.DrawCircle(cx, cy, 1)
		AngBotLeft(ctx)
		ctx.Fill()
	}

	if (mask&(NTop|NLeft|NSelf) == (NTop | NLeft | NSelf)) && (mask&NTopLeft == 0) {
		//ctx.SetRGB(1, 0, 0)
		ctx.DrawCircle(cx, cy, 2)
		AngTopLeft(ctx)
		ctx.Fill()
	}

	if (mask&(NTop|NRight|NSelf) == (NTop | NRight | NSelf)) && (mask&NTopRight == 0) {
		//ctx.SetRGB(0, 1, 0)
		ctx.DrawCircle(cx, cy, 2)
		AngTopRight(ctx)
		ctx.Fill()
	}

	ctx.DrawCircle(cx, cy, float64(radius))
	ctx.Fill()

	//ctx.SetRGB(1, 1, 1)
	//ctx.SetRGB(0, 0, 0)
	//ctx.SetFontFace()
	//ctx.DrawString(strconv.FormatInt(int64(mask), 10), x, cy)

}

func has(mask, bits uint16) bool {
	return mask&bits == bits
}

func DrawCorner(ctx *standard.DrawContext, cx, cy, r, l, w float64, angle float64) {
	// Координаты для 4 возможных вариантов углов:
	points := []struct{ x, y float64 }{
		{cx, cy + r},                  // верхняя центральная точка
		{cx - r, cy},                  // левая точка
		{cx - r, cy + float64(l)},     // нижняя левая точка
		{cx + r, cy + float64(l)},     // правая нижняя точка
		{cx + r, cy - r},              // контрольная точка для QuadraticTo
		{cx + float64(w) + l, cy - r}, // конечная точка для QuadraticTo
		{cx + float64(w), cy + r},     // окончательная точка для линии
	}

	// Поворот угла в зависимости от заданного угла
	x0, y0 := rotatePoint(points[0].x, points[0].y, cx, cy, angle)
	ctx.MoveTo(x0, y0)

	// Проход по линиям
	for i := 1; i < 4; i++ {
		xi, yi := rotatePoint(points[i].x, points[i].y, cx, cy, angle)
		ctx.LineTo(xi, yi)
	}

	// QuadraticTo (для изогнутой линии)
	ctrlX, ctrlY := rotatePoint(points[4].x, points[4].y, cx, cy, angle)
	endX, endY := rotatePoint(points[5].x, points[5].y, cx, cy, angle)
	ctx.QuadraticTo(ctrlX, ctrlY, endX, endY)

	// Завершающая линия
	xLast, yLast := rotatePoint(points[6].x, points[6].y, cx, cy, angle)
	ctx.LineTo(xLast, yLast)
	ctx.ClosePath()
}

func rotatePoint(x, y, cx, cy, angle float64) (float64, float64) {
	dx := x - cx
	dy := y - cy
	cosA := math.Cos(angle)
	sinA := math.Sin(angle)
	return cx + dx*cosA - dy*sinA, cy + dx*sinA + dy*cosA
}

const (
	NTopLeft  uint16 = 1 << iota // top-left
	NTop                         // top
	NTopRight                    // top-right
	NLeft                        // left
	NSelf                        // center (self)
	NRight                       // right
	NBotLeft                     // bottom-left
	NBot                         // bottom
	NBotRight                    // bottom-right
)

func main() {
	shape := newShape(0.5)
	qrc, err := qrcode.New("https://t.me/Potassium4O")
	// qrc, err := qrcode.New("with-custom-shape", qrcode.WithCircleShape())
	if err != nil {
		panic(err)
	}

	w, err := standard.New("./smaller.png",
		standard.WithCustomShape(shape),
	)
	if err != nil {
		panic(err)
	}

	err = qrc.Save(w)
	if err != nil {
		panic(err)
	}
}
