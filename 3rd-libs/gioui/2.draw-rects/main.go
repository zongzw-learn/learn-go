package main

import (
	"image"
	"image/color"
	"log"
	"math/rand"
	"time"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

type shape struct {
	r clip.Rect
	c color.NRGBA
}

var (
	rs     [50]shape
	index  int = 0
	size       = image.Point{200, 300}
	window     = new(app.Window)
)

func main() {
	go func() {
		var ops op.Ops
		window.Option(app.Title("Rand Rectangles"))
		for {
			switch ev := window.Event().(type) {
			case app.ConfigEvent:
			case app.DestroyEvent:
				log.Fatal(ev.Err)
			case app.FrameEvent:
				gtx := app.NewContext(&ops, ev)
				// log.Printf("max %v", gtx.Constraints.Max)
				// log.Printf("min %v", gtx.Constraints.Min)
				size = gtx.Constraints.Max
				// drawRedRect(gtx)
				drawRects(gtx)
				ev.Frame(gtx.Ops)
			}
		}
	}()

	go func() {
		for range time.Tick(time.Millisecond * 200) {
			addRects(size)
			window.Invalidate()
		}
	}()
	app.Main()
}

func addRects(sz image.Point) {
	rs[index] = shape{r: randRect(sz), c: randNRGBA()}
	index += 1
	index = index % len(rs)
}

func drawRects(gtx layout.Context) {
	for i := len(rs) - 1; i >= 0; i-- {
		id := (i + index) % len(rs)
		a := rs[id].r.Push(gtx.Ops)
		paint.ColorOp{Color: rs[id].c}.Add(gtx.Ops)
		paint.PaintOp{}.Add(gtx.Ops)
		a.Pop()
	}
}

func randNRGBA() color.NRGBA {
	r := uint8(rand.Uint32() % 256)
	g := uint8(rand.Uint32() % 256)
	b := uint8(rand.Uint32() % 256)
	a := uint8(rand.Uint32() % 256)
	return color.NRGBA{R: r, G: g, B: b, A: a}
}

func randRect(sz image.Point) clip.Rect {
	x := rand.Int() % sz.X
	y := rand.Int() % sz.Y
	e := rand.Int() % 300
	return clip.Rect{Min: image.Pt(x, y), Max: image.Pt(x+e, y+e)}
}
