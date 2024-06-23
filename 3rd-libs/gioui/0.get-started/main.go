package main

import (
	"log"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget/material"
)

func main() {
	go func() {
		window := new(app.Window)
		var ops op.Ops
		theme := material.NewTheme()
		window.Option(app.Title("Get Started"))
		for {
			switch ev := window.Event().(type) {
			case app.ConfigEvent:
			case app.DestroyEvent:
				log.Fatal(ev.Err)
			case app.FrameEvent:
				gtx := app.NewContext(&ops, ev)
				tt := material.Body1(theme, "Andrew Zong")
				layout.Center.Layout(gtx, tt.Layout)
				ev.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}
