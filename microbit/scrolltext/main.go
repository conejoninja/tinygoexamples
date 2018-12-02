package main

import (
	"time"

	"github.com/conejoninja/tinygoexamples/microbit/scrolltext/scroll"
	"github.com/aykevl/tinygo/src/machine"
)

type Matrix [3][9]uint8

var matrixSettings = [5][5][2]uint8{
	{{0, 0}, {1, 3}, {0, 1}, {1, 4}, {0, 2}},
	{{2, 3}, {2, 4}, {2, 5}, {2, 6}, {2, 7}},
	{{1, 1}, {0, 8}, {1, 2}, {2, 8}, {1, 0}},
	{{0, 7}, {0, 6}, {0, 5}, {0, 4}, {0, 3}},
	{{2, 2}, {1, 6}, {2, 0}, {1, 5}, {2, 1}},
}

func main() {

	println("Starting micro:bit")

	d := scroll.NewDisplay()
	d.SetText("HELLO GOPHERS! ")


	buttonA := machine.GPIO{machine.BUTTONA}
	buttonA.Configure(machine.GPIOConfig{Mode: machine.GPIO_INPUT})
	statusA := false

	buttonB := machine.GPIO{machine.BUTTONB}
	buttonB.Configure(machine.GPIOConfig{Mode: machine.GPIO_INPUT})
	statusB := false


	var rotation uint8
	var direction uint8

	d.SetRotation(rotation)
	d.SetDirection(direction)

	t := time.Now()
	for {

		d.ShowFrame()

		delta := time.Now().Sub(t)
		if delta > time.Millisecond*150 {
			if !buttonA.Get() {
				if !statusA {
					rotation = (rotation+1)%4
					d.SetRotation(rotation)
				}
				statusA = true
			} else {
				statusA = false
			}

			if !buttonB.Get() {
				if !statusB {
					direction = (direction+1)%4
					d.SetDirection(direction)
				}
				statusB = true
			} else {
				statusB = false
			}

			t = time.Now()

			d.NextFrame()

		}
	}

}

