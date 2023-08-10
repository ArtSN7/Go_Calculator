package main


import (
	"strconv"

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


	var first_number, second_number, symbol string
	var first_number_was, second_number_was, symbol_was bool

	first_number_was = false 
	second_number_was = false
	symbol_was = false

	first_number = ""
	second_number = ""
	symbol = ""


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
					input_text += "0"

				} else if !symbol_was{
					first_number += "0"
					input_text += "0"
				}}}
        
		if oneButton.Clicked(){
			if symbol_was{
				second_number += "1"
				input_text = second_number
				second_number_was = true
			}else{
				first_number += "1"
				input_text = first_number
				first_number_was = true
			}
		}
        
		if twoButton.Clicked(){
			if symbol_was{
				second_number += "2"
				input_text = second_number
				second_number_was = true
			}else{
				first_number += "2"
				input_text = first_number
				first_number_was = true
			}
		}

		if threeButton.Clicked(){
			if symbol_was{
				second_number += "3"
				input_text = second_number
				second_number_was = true
			}else{
				first_number += "3"
				input_text = first_number
				first_number_was = true
			}
		}

		if fourButton.Clicked(){
			if symbol_was{
				second_number += "4"
				input_text = second_number
				second_number_was = true
			}else{
				first_number += "4"
				input_text = first_number
				first_number_was = true
			}
		}

		if fiveButton.Clicked(){
			if symbol_was{
				second_number += "5"
				input_text = second_number
				second_number_was = true
			}else{
				first_number += "5"
				input_text = first_number
				first_number_was = true
			}
		}

		if sixButton.Clicked(){
			if symbol_was{
				second_number += "6"
				input_text = second_number
				second_number_was = true
			}else{
				first_number += "6"
				input_text = first_number
				first_number_was = true
			}
		}

		if sevenButton.Clicked(){
			if symbol_was{
				second_number += "7"
				input_text = second_number
				second_number_was = true
			}else{
				first_number += "7"
				input_text = first_number
				first_number_was = true
			}
		}

		if eightButton.Clicked(){
			if symbol_was{
				second_number += "8"
				input_text = second_number
				second_number_was = true
			}else{
				first_number += "8"
				input_text = first_number
				first_number_was = true
			}
		}

		if nineButton.Clicked(){
			if symbol_was{
				second_number += "9"
				input_text = second_number
				second_number_was = true
			}else{
				first_number += "9"
				input_text = first_number
				first_number_was = true
			}
		}

		//знаки арифм.

		if plusButton.Clicked(){
			if first_number_was{
				symbol = "+"
				symbol_was = true
				input_text = "+"
			}
		}

		if minusButton.Clicked(){
			if first_number_was{
				symbol = "-"
				symbol_was = true
				input_text = "-"
			}
		}

		if multButton.Clicked(){
			if first_number_was{
				symbol = "*"
				symbol_was = true
				input_text = "*"
			}
		}

		if devideButton.Clicked(){
			if first_number_was{
				symbol = "%"
				symbol_was = true
				input_text = "%"
			}
		}

		// =

		if equalButton.Clicked(){
			if symbol_was && first_number_was && second_number_was{
				input_text = calculation(first_number, second_number, symbol)
			}else{
				input_text = "ERROR"
			}
		}
        

		input_text = GoodOutput(input_text)

		case system.DestroyEvent:
			return e.Err
		}
	}
	return nil
}


func calculation(a string, b string, symbol string) string{
    answer := ""

	i, err := strconv.Atoi(a)

    if err != nil {
        // ... handle error
        panic(err)
    }

	j, err := strconv.Atoi(b)

    if err != nil {
        // ... handle error
        panic(err)
    }
    
	switch{
	case symbol == "+": answer = string(i + j)
	case symbol == "-": answer = string(i - j)
	case symbol == "*": answer = string(i * j)
	case symbol == "%": answer = string(i / j)
	}
    
	return answer
	}


func GoodOutput(a string) string{
	// 50 symbols
	len_of_a := len(a)

	if len_of_a < 50{
		len_of_a = 50 - len_of_a
	}else{len_of_a = 0}
    
	for i := 0; i < len_of_a; i++{
		a += " "
	}

	return a
}