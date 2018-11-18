package main

import (
	"time"

	"github.com/conejoninja/tinygoexamples/microbit/scrolltext/fonts"
	"github.com/aykevl/tinygo/src/device/nrf"
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


	// TODO: proper mapping 
	//                   H  E  L   L   O       G  O   P   H  E  R   S   !
	msg := []int{63, 63, 7, 4, 11, 11, 14, 63, 6, 14, 15, 7, 4, 17, 18, 62, 63}
	msgPattern := msg2pattern(msg)

	scrollIndex := 0
	t := time.Now()
	m := scroll2matrix(msgPattern, scrollIndex)
	machine.InitLEDMatrix()
	for {

		for row := 0; row < 3; row++ {
			machine.ClearLEDMatrix()
			set := 0
			set |= 1 << uint8(machine.LED_ROW_1+row)
			nrf.GPIO.OUTSET = nrf.RegValue(set)

			set = 0
			for col := 0; col < 9; col++ {

				if m[row][col] > 0 {
					set |= 1 << uint8(machine.LED_COL_1+col)
				}

			}
			nrf.GPIO.OUTCLR = nrf.RegValue(set)
			time.Sleep(time.Millisecond * 2)
		}

		delta := time.Now().Sub(t)
		if delta > time.Millisecond*150 {
			println("HELLO GOPHERS! ")
			t = time.Now()
			scrollIndex++
			if scrollIndex >= (len(msg)-1)*6 {
				scrollIndex = 0
			}
			m = scroll2matrix(msgPattern, scrollIndex)
		}
	}

}


func scroll2matrix(pattern [5][]uint8, index int) (m Matrix) {
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			if pattern[x][index+y] > 0 {
				m[matrixSettings[x][y][0]][matrixSettings[x][y][1]] = pattern[x][index+y]
			}
		}
	}
	return m
}

func msg2pattern(msg []int) (d [5][]uint8) {
	for i := 0; i < 5; i++ {
		d[i] = make([]uint8, len(msg)*6)
	}
	for i, c := range msg {
		for x := 0; x < 5; x++ {
			for y := 0; y < 5; y++ {
				d[x][i*6+y] = fonts.Microbit5x5Font[c][x][y]
			}
		}
	}
	return d
}
