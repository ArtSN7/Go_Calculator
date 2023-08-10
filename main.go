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
			app.Size(unit.Dp(1000), unit.Dp(125)),
            app.MaxSize(unit.Dp(1000), unit.Dp(125)),
            app.MinSize(unit.Dp(1000), unit.Dp(125)),
		)

		if err := maker(w); err != nil {
			log.Fatal(err)
		}

		os.Exit(0)
	}()

	app.Main()
}