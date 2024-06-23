package main

import (
	"image"
	"image/color"
	"log"
	"os"
	"time"

	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/widget/material"
)

func main() {
	go func() {
		window := new(app.Window)
		err := run(window)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func run(window *app.Window) error {
	theme := material.NewTheme()
	var ops op.Ops
	window.Option(app.Title("sample demo"))
	for {
		switch e := window.Event().(type) {
		case app.ConfigEvent:
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			// This graphics context is used for managing the rendering state.
			t := time.Now().String()
			gtx := app.NewContext(&ops, e)

			// Define an large label with an appropriate text:
			title := material.H2(theme, t)

			// Change the color of the label.
			maroon := color.NRGBA{R: 127, G: 0, B: 0, A: 255}
			title.Color = maroon

			// Change the position of the label.
			title.Alignment = text.Middle
			// Draw the label to the graphics context.
			title.Layout(gtx)

			layout.Center.Layout(gtx, title.Layout)

			content := material.Body1(theme, "zongzw")
			content.Alignment = text.Middle
			content.Layout(gtx)

			redButtonBackground(gtx.Ops)
			// move(gtx.Ops)
			// Pass the drawing operations to the GPU.
			// strokeTriangle(gtx.Ops)
			e.Frame(gtx.Ops)
		}
	}
}

func drawRedRect(ops *op.Ops) {
	// defer clip.Rect{Max: image.Pt(100, 100)}.Push(ops).Pop()
	// paint.ColorOp{Color: color.NRGBA{R: 0x80, A: 0xFF}}.Add(ops)
	// paint.PaintOp{}.Add(ops)
	stack := clip.Rect{Min: image.Pt(300, 300), Max: image.Pt(400, 400)}.Push(ops)
	defer stack.Pop()
	paint.ColorOp{Color: color.NRGBA{R: 0x80, A: 0xFF}}.Add(ops)
	paint.PaintOp{}.Add(ops)
}

func redButtonBackground(ops *op.Ops) {
	const r = 10 // roundness
	bounds := image.Rect(0, 0, 400, 400)
	stack := clip.RRect{Rect: bounds, SE: r, SW: r, NW: r, NE: r}.Push(ops)
	paint.ColorOp{Color: color.NRGBA{R: 0x50, A: 0xf0}}.Add(ops)
	paint.PaintOp{}.Add(ops)
	drawRedRect(ops)
	stack.Pop()
}

func move(ops *op.Ops) {
	drawRedRect(ops)
	op.Offset(image.Pt(100, 0)).Push(ops).Pop()
	// <-time.After(time.Second)
	// op.Offset(image.Pt(200, 0)).Push(ops).Pop()
	// <-time.After(time.Second)
	// op.Offset(image.Pt(300, 0)).Push(ops).Pop()
	// <-time.After(time.Second)
}

func strokeTriangle(ops *op.Ops) {
	var path clip.Path
	path.Begin(ops)
	path.MoveTo(f32.Pt(30, 30))
	path.LineTo(f32.Pt(70, 30))
	path.LineTo(f32.Pt(50, 70))
	path.Close()

	paint.FillShape(ops, color.NRGBA{R: 0x80, A: 0xFF},
		clip.Stroke{
			Path:  path.End(),
			Width: 4,
		}.Op())
}
