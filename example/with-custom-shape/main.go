package main

import (
	"image/color"

	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
)

type smallerCircle struct {
	smallerPercent float64
}

func (sc *smallerCircle) DrawFinder(ctx *standard.DrawContext) {
	//backup := sc.smallerPercent
	//sc.smallerPercent = 1.0
	//sc.Draw(ctx)
	//sc.smallerPercent = backup

	w, h := ctx.Edge()
	x, y := ctx.UpperLeft()
	color := ctx.Color()
	//ctx.DrawRegularPolygon()

	// choose a proper radius values
	//radius := w / 2
	//r2 := h / 2
	//if r2 <= radius {
	//	radius = r2
	//}

	// 80 percent smaller
	//radius = int(float64(radius) * sc.smallerPercent)

	//cx, cy := x+float64(w)/2.0, y+float64(h)/2.0 // get center point
	//ctx.DrawCircle(cx, cy, float64(radius))
	ctx.DrawRoundedRectangle(x, y, float64(w), float64(h), 1)
	ctx.SetColor(color)
	ctx.Fill()
}

func newShape(radiusPercent float64) standard.IShape {
	return &smallerCircle{smallerPercent: radiusPercent}
}

func (sc *smallerCircle) Draw(ctx *standard.DrawContext) {
	w, h := ctx.Edge()
	x, y := ctx.UpperLeft()
	color := ctx.Color()

	// choose a proper radius values
	radius := w / 2
	r2 := h / 2
	if r2 <= radius {
		radius = r2
	}

	// 80 percent smaller
	radius = int(float64(radius) * sc.smallerPercent)

	cx, cy := x+float64(w)/2.0, y+float64(h)/2.0 // get center point
	ctx.DrawCircle(cx, cy, float64(radius))
	ctx.SetColor(color)
	ctx.Fill()

}

func main() {
	shape := newShape(0.8)
	qrc, err := qrcode.NewWith("https://www.wildberries.ru/lk/mywallet/purchases", qrcode.WithErrorCorrectionLevel(qrcode.ErrorCorrectionLow))
	// qrc, err := qrcode.New("with-custom-shape", qrcode.WithCircleShape())
	if err != nil {
		panic(err)
	}
	//g := standard.NewGradient(270, []standard.ColorStop{
	//	{
	//		T:     0,
	//		Color: color.RGBA{R: 0, G: 151, B: 57, A: 255},
	//	},
	//	{
	//		T:     0.5,
	//		Color: color.RGBA{R: 254, G: 221, B: 0, A: 255},
	//	},
	//	{
	//		T:     1,
	//		Color: color.RGBA{R: 1, G: 33, B: 105, A: 255},
	//	},
	//	//{
	//	//	T:     0,
	//	//	Color: color.RGBA{R: 203, G: 17, B: 171, A: 255},
	//	//},
	//	//{
	//	//	T:     0.5,
	//	//	Color: color.RGBA{R: 153, G: 0, B: 153, A: 255},
	//	//},
	//	//{
	//	//	T:     1,
	//	//	Color: color.RGBA{R: 72, G: 17, B: 115, A: 255},
	//	//},
	//}...)
	w, err := standard.New("./smaller.png",
		//standard.WithFgGradient(g),
		standard.WithBgColor(color.RGBA{R: 255, G: 255, B: 255, A: 255}),
		standard.WithCustomShape(shape),
		//standard.WithLogoImageFilePNG("./with-custom-shape/wblogo1_200x200.png"),
		//standard.WithLogoSizeMultiplier(1),
	)
	if err != nil {
		panic(err)
	}

	err = qrc.Save(w)
	if err != nil {
		panic(err)
	}
}
