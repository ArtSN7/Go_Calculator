package main


import (
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/unit"

)


// start
func main() {
	go func() {

		w := app.NewWindow(
			app.Title("CALCULATOR"),
			app.Size(unit.Dp(800), unit.Dp(120)),
		)

		if err := maker(w); err != nil {
			log.Fatal(err)
		}

		os.Exit(0)
	}()

	app.Main()
}