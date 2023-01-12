package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"github.com/robotn/gohook"
	"time"
)

func main() {
	fmt.Println("Auto Click")
	s := hook.Start()
	defer hook.End()
	autoClickActive := false
	const rawCode uint16 = 117

	go func() {
		for {
			if autoClickActive {
				robotgo.Click("left", false)
			}
			time.Sleep(time.Millisecond * 10)
		}
	}()

	fmt.Println("Press F6 to enable/disable the Auto Click")
	for {

		select {
		case i := <-s:
			if i.Rawcode == rawCode {
				if i.Kind > hook.KeyDown && i.Kind < hook.KeyUp {

					autoClickActive = !autoClickActive
					if autoClickActive {
						fmt.Println("Enabled")
					} else {
						fmt.Println("Disabled")
					}
				}
			}
		}
	}

}
