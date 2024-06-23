package main

import (
	"log"
	"time"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget/material"
)

func main() {
	window := new(app.Window)
	var op op.Ops
	go func() {
		theme := material.NewTheme()
		window.Option(app.Title("Animation"))
		for {
			switch ev := window.Event().(type) {
			case app.ConfigEvent:
			case app.DestroyEvent:
				log.Fatal(ev.Err)
			case app.FrameEvent:
				t := time.Now().Format("2006-01-02 15:04:05")

				gtx := app.NewContext(&op, ev)

				tt := material.Body1(theme, t)
				layout.Center.Layout(gtx, tt.Layout)
				ev.Frame(gtx.Ops)
			}
		}
	}()
	go func() {
		for range time.Tick(time.Second) {
			window.Invalidate()
		}
	}()
	app.Main()
}
