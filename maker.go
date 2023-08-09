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

	var oneButton, twoButton, threeButton, fourButton, fiveButton, sixButton, sevenButton, eightButton, nineButton, zeroButton, clearButton, plusButton, minusButton, devideButton, multButton, equalButton widget.Clickable
	var inputButton widget.Clickable


    var input_text string
    input_text = "There will be ur input ( press = to get a result )"


	th := material.NewTheme()

	for e := range w.Events() {

		switch e := e.(type) {

		case system.FrameEvent:
			// buttons 
			gtx := layout.NewContext(&ops, e)
			layout.Flex{
				Axis: layout.Horizontal,
				Spacing: layout.Spacing(2),

			}.Layout(gtx,
                
				//Answer
				layout.Rigid(func(gtx C) D {
					margins := layout.Inset{
						Top:    unit.Dp(80),
						Bottom: unit.Dp(5),
						Right:  unit.Dp(0),
						Left:   unit.Dp(3),}
			
					return margins.Layout(gtx,
						func(gtx C) D {
						    btn := material.Button(th, &inputButton, input_text)
						    return btn.Layout(gtx)
						},
					)
				},),

                    //1
				layout.Rigid(func(gtx C) D {
					margins := layout.Inset{
						Top:    unit.Dp(80),
						Bottom: unit.Dp(3),
						Right:  unit.Dp(0),
						Left:   unit.Dp(3),}
			
					return margins.Layout(gtx,
						func(gtx C) D {
						    btn := material.Button(th, &oneButton, "1")
						    return btn.Layout(gtx)
						},
					)
				},),
			
			
					//2
				layout.Rigid(func(gtx C) D {
					margins := layout.Inset{
						Top:    unit.Dp(80),
						Bottom: unit.Dp(3),
						Right:  unit.Dp(0),
						Left:   unit.Dp(3),}
			
					return margins.Layout(gtx,
						func(gtx C) D {
						    btn := material.Button(th, &twoButton, "2")
						    return btn.Layout(gtx)
						},
					)
				},),
			
			
			
					//3
				layout.Rigid(func(gtx C) D {
					margins := layout.Inset{
						Top:    unit.Dp(80),
						Bottom: unit.Dp(3),
						Right:  unit.Dp(0),
						Left:   unit.Dp(3),}
			
					return margins.Layout(gtx,
						func(gtx C) D {
						    btn := material.Button(th, &threeButton, "3")
						    return btn.Layout(gtx)
						},
					)
				},),
			
			
					//4
				layout.Rigid(func(gtx C) D {
					margins := layout.Inset{
						Top:    unit.Dp(80),
						Bottom: unit.Dp(3),
						Right:  unit.Dp(0),
						Left:   unit.Dp(3),}
			
					return margins.Layout(gtx,
						func(gtx C) D {
						    btn := material.Button(th, &fourButton, "4")
						    return btn.Layout(gtx)
						},
					)
				},),
			
			
					//5
				layout.Rigid(func(gtx C) D {
					margins := layout.Inset{
						Top:    unit.Dp(80),
						Bottom: unit.Dp(3),
						Right:  unit.Dp(0),
						Left:   unit.Dp(3),}
			
					return margins.Layout(gtx,
						func(gtx C) D {
						    btn := material.Button(th, &fiveButton, "5")
						    return btn.Layout(gtx)
						},
					)
				},),
			
			
					//6
				layout.Rigid(func(gtx C) D {
					margins := layout.Inset{
						Top:    unit.Dp(80),
						Bottom: unit.Dp(3),
						Right:  unit.Dp(0),
						Left:   unit.Dp(3),}
			
					return margins.Layout(gtx,
						func(gtx C) D {
						    btn := material.Button(th, &sixButton, "6")
						    return btn.Layout(gtx)
						},
					)
				},),
			
			
					//7
				layout.Rigid(func(gtx C) D {
					margins := layout.Inset{
						Top:    unit.Dp(80),
						Bottom: unit.Dp(3),
						Right:  unit.Dp(0),
						Left:   unit.Dp(3),}
			
					return margins.Layout(gtx,
						func(gtx C) D {
						    btn := material.Button(th, &sevenButton, "7")
						    return btn.Layout(gtx)
						},
					)
				},),
			
			
					//8
				layout.Rigid(func(gtx C) D {
					margins := layout.Inset{
						Top:    unit.Dp(80),
						Bottom: unit.Dp(3),
						Right:  unit.Dp(0),
						Left:   unit.Dp(3),}
			
					return margins.Layout(gtx,
						func(gtx C) D {
						    btn := material.Button(th, &eightButton, "8")
						    return btn.Layout(gtx)
						},
					)
				},),
			
			
					//9
				layout.Rigid(func(gtx C) D {
					margins := layout.Inset{
						Top:    unit.Dp(80),
						Bottom: unit.Dp(3),
						Right:  unit.Dp(0),
						Left:   unit.Dp(3),}
			
					return margins.Layout(gtx,
						func(gtx C) D {
						    btn := material.Button(th, &nineButton, "9")
						    return btn.Layout(gtx)
						},
					)
				},),
			
			
					//0
				layout.Rigid(func(gtx C) D {
					margins := layout.Inset{
						Top:    unit.Dp(80),
						Bottom: unit.Dp(3),
						Right:  unit.Dp(0),
						Left:   unit.Dp(3),}
			
					return margins.Layout(gtx,
						func(gtx C) D {
						    btn := material.Button(th, &zeroButton, "0")
						    return btn.Layout(gtx)
						},
					)
				},),
			
			
					//clear
				layout.Rigid(func(gtx C) D {
					margins := layout.Inset{
						Top:    unit.Dp(80),
						Bottom: unit.Dp(3),
						Right:  unit.Dp(0),
						Left:   unit.Dp(3),}
			
					return margins.Layout(gtx,
						func(gtx C) D {
						    btn := material.Button(th, &equalButton, "=")
						    return btn.Layout(gtx)
						},
					)
				},),
			
			
					//=
				layout.Rigid(func(gtx C) D {
					margins := layout.Inset{
						Top:    unit.Dp(80),
						Bottom: unit.Dp(3),
						Right:  unit.Dp(0),
						Left:   unit.Dp(3),}
			
					return margins.Layout(gtx,
						func(gtx C) D {
						    btn := material.Button(th, &clearButton, "CLEAR")
						    return btn.Layout(gtx)
						},
					)
				},),
			
				//+
				layout.Rigid(func(gtx C) D {
					margins := layout.Inset{
						Top:    unit.Dp(80),
						Bottom: unit.Dp(3),
						Right:  unit.Dp(0),
						Left:   unit.Dp(3),}
			
					return margins.Layout(gtx,
						func(gtx C) D {
						    btn := material.Button(th, &plusButton, "+")
						    return btn.Layout(gtx)
						},
					)
				},),
			
				//-
				layout.Rigid(func(gtx C) D {
					margins := layout.Inset{
						Top:    unit.Dp(80),
						Bottom: unit.Dp(3),
						Right:  unit.Dp(0),
						Left:   unit.Dp(3),}
			
					return margins.Layout(gtx,
						func(gtx C) D {
						    btn := material.Button(th, &minusButton, "-")
						    return btn.Layout(gtx)
						},
					)
				},),
			

				//*
				layout.Rigid(func(gtx C) D {
					margins := layout.Inset{
						Top:    unit.Dp(80),
						Bottom: unit.Dp(3),
						Right:  unit.Dp(0),
						Left:   unit.Dp(3),}
			
					return margins.Layout(gtx,
						func(gtx C) D {
						    btn := material.Button(th, &multButton, "*")
						    return btn.Layout(gtx)
						},
					)
				},),
			
				// %
				layout.Rigid(func(gtx C) D {
					margins := layout.Inset{
						Top:    unit.Dp(80),
						Bottom: unit.Dp(3),
						Right:  unit.Dp(0),
						Left:   unit.Dp(3),}
			
					return margins.Layout(gtx,
						func(gtx C) D {
						    btn := material.Button(th, &devideButton, "%")
						    return btn.Layout(gtx)
						},
					)
				},),
			
        


		// the end of layout.Rigid()
			)
			e.Frame(gtx.Ops)

        // отвечает за проверку нажатий на кнопку

		var first_number, second_number, symbol string
		var first_number_was, second_number_was, symbol_was bool

		first_number_was = false 
		second_number_was = false
		symbol_was = false

		first_number = ""
		second_number = ""
		symbol = ""

		//кнопка очищения
        if clearButton.Clicked(){
			first_number_was = false 
			second_number_was = false
			symbol_was = false
	
			first_number = ""
			second_number = ""
			symbol = ""
			input_text = "There will be ur input ( press = to get a result )"
		}

        // около 50 знаков надо для верного размера кнопки

        // кнопки чисел
        if zeroButton.Clicked(){
			if first_number_was{
				if second_number_was{
					second_number += "0"

				} else if symbol_was == false{
					first_number += "0"
				}}}

        









		case system.DestroyEvent:
			return e.Err
		}
	}
	return nil
}