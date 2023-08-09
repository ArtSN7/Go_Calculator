package main


import (
	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)


type C = layout.Context
type D = layout.Dimensions

func maker(w *app.Window) error {

	var ops op.Ops

	var startButton widget.Clickable

	th := material.NewTheme()

	for e := range w.Events() {

		switch e := e.(type) {

		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)
			layout.Flex{
				Axis: layout.Vertical,
				Spacing: layout.SpaceStart,

			}.Layout(gtx,
                    //1
				layout.Rigid(func(gtx C) D {
					margins := layout.Inset{
						Top:    unit.Dp(25),
						Bottom: unit.Dp(25),
						Right:  unit.Dp(35),
						Left:   unit.Dp(35),}
			
					return margins.Layout(gtx,
						func(gtx C) D {
						    btn := material.Button(th, &startButton, "1")
						    return btn.Layout(gtx)
						},
					)
				},),
			
			
					//2
				layout.Rigid(func(gtx C) D {
					margins := layout.Inset{
						Top:    unit.Dp(25),
						Bottom: unit.Dp(25),
						Right:  unit.Dp(35),
						Left:   unit.Dp(35),}
			
					return margins.Layout(gtx,
						func(gtx C) D {
						    btn := material.Button(th, &startButton, "Start")
						    return btn.Layout(gtx)
						},
					)
				},),
			
			
			
					//3
				layout.Rigid(func(gtx C) D {
					margins := layout.Inset{
						Top:    unit.Dp(25),
						Bottom: unit.Dp(25),
						Right:  unit.Dp(35),
						Left:   unit.Dp(35),}
			
					return margins.Layout(gtx,
						func(gtx C) D {
						    btn := material.Button(th, &startButton, "Start")
						    return btn.Layout(gtx)
						},
					)
				},),
			
			
					//4
				layout.Rigid(func(gtx C) D {
					margins := layout.Inset{
						Top:    unit.Dp(25),
						Bottom: unit.Dp(25),
						Right:  unit.Dp(35),
						Left:   unit.Dp(35),}
			
					return margins.Layout(gtx,
						func(gtx C) D {
						    btn := material.Button(th, &startButton, "Start")
						    return btn.Layout(gtx)
						},
					)
				},),
			
			
					//5
				layout.Rigid(func(gtx C) D {
					margins := layout.Inset{
						Top:    unit.Dp(25),
						Bottom: unit.Dp(25),
						Right:  unit.Dp(35),
						Left:   unit.Dp(35),}
			
					return margins.Layout(gtx,
						func(gtx C) D {
						    btn := material.Button(th, &startButton, "Start")
						    return btn.Layout(gtx)
						},
					)
				},),
			
			
					//6
				layout.Rigid(func(gtx C) D {
					margins := layout.Inset{
						Top:    unit.Dp(25),
						Bottom: unit.Dp(25),
						Right:  unit.Dp(35),
						Left:   unit.Dp(35),}
			
					return margins.Layout(gtx,
						func(gtx C) D {
						    btn := material.Button(th, &startButton, "Start")
						    return btn.Layout(gtx)
						},
					)
				},),
			
			
					//7
				layout.Rigid(func(gtx C) D {
					margins := layout.Inset{
						Top:    unit.Dp(25),
						Bottom: unit.Dp(25),
						Right:  unit.Dp(35),
						Left:   unit.Dp(35),}
			
					return margins.Layout(gtx,
						func(gtx C) D {
						    btn := material.Button(th, &startButton, "Start")
						    return btn.Layout(gtx)
						},
					)
				},),
			
			
					//8
				layout.Rigid(func(gtx C) D {
					margins := layout.Inset{
						Top:    unit.Dp(25),
						Bottom: unit.Dp(25),
						Right:  unit.Dp(35),
						Left:   unit.Dp(35),}
			
					return margins.Layout(gtx,
						func(gtx C) D {
						    btn := material.Button(th, &startButton, "Start")
						    return btn.Layout(gtx)
						},
					)
				},),
			
			
					//9
				layout.Rigid(func(gtx C) D {
					margins := layout.Inset{
						Top:    unit.Dp(25),
						Bottom: unit.Dp(25),
						Right:  unit.Dp(35),
						Left:   unit.Dp(35),}
			
					return margins.Layout(gtx,
						func(gtx C) D {
						    btn := material.Button(th, &startButton, "Start")
						    return btn.Layout(gtx)
						},
					)
				},),
			
			
					//0
				layout.Rigid(func(gtx C) D {
					margins := layout.Inset{
						Top:    unit.Dp(25),
						Bottom: unit.Dp(25),
						Right:  unit.Dp(35),
						Left:   unit.Dp(35),}
			
					return margins.Layout(gtx,
						func(gtx C) D {
						    btn := material.Button(th, &startButton, "Start")
						    return btn.Layout(gtx)
						},
					)
				},),
			
			
					//clear
				layout.Rigid(func(gtx C) D {
					margins := layout.Inset{
						Top:    unit.Dp(25),
						Bottom: unit.Dp(25),
						Right:  unit.Dp(35),
						Left:   unit.Dp(35),}
			
					return margins.Layout(gtx,
						func(gtx C) D {
						    btn := material.Button(th, &startButton, "Start")
						    return btn.Layout(gtx)
						},
					)
				},),
			
			
					//=
				layout.Rigid(func(gtx C) D {
					margins := layout.Inset{
						Top:    unit.Dp(25),
						Bottom: unit.Dp(25),
						Right:  unit.Dp(35),
						Left:   unit.Dp(35),}
			
					return margins.Layout(gtx,
						func(gtx C) D {
						    btn := material.Button(th, &startButton, "Start")
						    return btn.Layout(gtx)
						},
					)
				},),
			
				//+
				layout.Rigid(func(gtx C) D {
					margins := layout.Inset{
						Top:    unit.Dp(25),
						Bottom: unit.Dp(25),
						Right:  unit.Dp(35),
						Left:   unit.Dp(35),}
			
					return margins.Layout(gtx,
						func(gtx C) D {
						    btn := material.Button(th, &startButton, "Start")
						    return btn.Layout(gtx)
						},
					)
				},),
			
				//-
				layout.Rigid(func(gtx C) D {
					margins := layout.Inset{
						Top:    unit.Dp(25),
						Bottom: unit.Dp(25),
						Right:  unit.Dp(35),
						Left:   unit.Dp(35),}
			
					return margins.Layout(gtx,
						func(gtx C) D {
						    btn := material.Button(th, &startButton, "Start")
						    return btn.Layout(gtx)
						},
					)
				},),
			

				//*
				layout.Rigid(func(gtx C) D {
					margins := layout.Inset{
						Top:    unit.Dp(25),
						Bottom: unit.Dp(25),
						Right:  unit.Dp(35),
						Left:   unit.Dp(35),}
			
					return margins.Layout(gtx,
						func(gtx C) D {
						    btn := material.Button(th, &startButton, "Start")
						    return btn.Layout(gtx)
						},
					)
				},),
			
				// %
				layout.Rigid(func(gtx C) D {
					margins := layout.Inset{
						Top:    unit.Dp(25),
						Bottom: unit.Dp(25),
						Right:  unit.Dp(35),
						Left:   unit.Dp(35),}
			
					return margins.Layout(gtx,
						func(gtx C) D {
						    btn := material.Button(th, &startButton, "Start")
						    return btn.Layout(gtx)
						},
					)
				},),
			



		// the end of layout.Rigid()
			)
			e.Frame(gtx.Ops)


		case system.DestroyEvent:
			return e.Err
		}
	}
	return nil
}