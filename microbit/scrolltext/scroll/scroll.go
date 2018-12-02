package scroll

import (
	"time"

	"github.com/aykevl/tinygo/src/device/nrf"
	"github.com/aykevl/tinygo/src/machine"
	"github.com/conejoninja/tinygoexamples/microbit/scrolltext/fonts"
)

type Matrix [3][9]uint8

type Display struct {
	direction  uint8
	rotation   uint8
	msg        []byte
	charBuffer [5][12]uint8
	buffer     Matrix
	bigStep    int
	step       int
}

var matrixSettings = [4][5][5][2]uint8{
	{ // 0
		{{0, 0}, {1, 3}, {0, 1}, {1, 4}, {0, 2}},
		{{2, 3}, {2, 4}, {2, 5}, {2, 6}, {2, 7}},
		{{1, 1}, {0, 8}, {1, 2}, {2, 8}, {1, 0}},
		{{0, 7}, {0, 6}, {0, 5}, {0, 4}, {0, 3}},
		{{2, 2}, {1, 6}, {2, 0}, {1, 5}, {2, 1}},
	},
	{ // 90 CCW
		{{0, 2}, {2, 7}, {1, 0}, {0, 3}, {2, 1}},
		{{1, 4}, {2, 6}, {2, 8}, {0, 4}, {1, 5}},
		{{0, 1}, {2, 5}, {1, 2}, {0, 5}, {2, 0}},
		{{1, 3}, {2, 4}, {0, 8}, {0, 6}, {1, 6}},
		{{0, 0}, {2, 3}, {1, 1}, {0, 7}, {2, 2}},
	},
	{ // 180
		{{2, 1}, {1, 5}, {2, 0}, {1, 6}, {2, 2}},
		{{0, 3}, {0, 4}, {0, 5}, {0, 6}, {0, 7}},
		{{1, 0}, {2, 8}, {1, 2}, {0, 8}, {1, 1}},
		{{2, 7}, {2, 6}, {2, 5}, {2, 4}, {2, 3}},
		{{0, 2}, {1, 4}, {0, 1}, {1, 3}, {0, 0}},
	},
	{ // 270
		{{2, 2}, {0, 7}, {1, 1}, {2, 3}, {0, 0}},
		{{1, 6}, {0, 6}, {0, 8}, {2, 4}, {1, 3}},
		{{2, 0}, {0, 5}, {1, 2}, {2, 5}, {0, 1}},
		{{1, 5}, {0, 4}, {2, 8}, {2, 6}, {1, 4}},
		{{2, 1}, {0, 3}, {1, 0}, {2, 7}, {0, 2}},
	},
}

func NewDisplay() *Display {
	var d Display
	machine.InitLEDMatrix()
	return &d
}

func (d *Display) SetRotation(rotation uint8) {
	if rotation < 4 {
		d.rotation = rotation
	}
	d.Reset()
}

func (d *Display) SetDirection(direction uint8) {
	if direction < 4 {
		d.direction = direction
	}
	d.Reset()
}

func (d *Display) Rotation() uint8 {
	return d.direction
}

func (d *Display) ShowFrame() {
	for row := 0; row < 3; row++ {
		machine.ClearLEDMatrix()
		set := 0
		set |= 1 << uint8(machine.LED_ROW_1+row)
		nrf.GPIO.OUTSET = nrf.RegValue(set)

		set = 0
		for col := 0; col < 9; col++ {

			if d.buffer[row][col] > 0 {
				set |= 1 << uint8(machine.LED_COL_1+col)
			}

		}
		nrf.GPIO.OUTCLR = nrf.RegValue(set)
		time.Sleep(time.Millisecond * 2)
	}
}

func (d *Display) NextFrame() {
	d.step++
	if d.step > 5 {
		d.step = 0
		d.bigStep++
		if d.bigStep >= len(d.msg) {
			d.bigStep = 0
		}
	}
	d.doCharBuffer()
	d.doMatrixBuffer()
}

func (d *Display) SetText(msg string) {
	d.msg = []byte(msg)
	d.Reset()
}
func (d *Display) Reset() {
	d.bigStep = 0
	d.step = 0
	d.doCharBuffer()
	d.doMatrixBuffer()
}

func (d *Display) doCharBuffer() {
	rotation := (d.direction+d.rotation)%4
	for i := 0; i < 2; i++ {
		c := d.msg[(i+d.bigStep)%len(d.msg)]
		m := fonts.Microbit5x5Font[fonts.CharOffset(c)]
		for x := 0; x < 5; x++ {
			for y := 0; y < 5; y++ {
				if rotation == 0 {
					d.charBuffer[x][i*6+y] = m[x][y]
				} else if rotation == 1 {
					d.charBuffer[x][i*6+y] = m[y][4-x]
				} else if rotation == 2 {
					d.charBuffer[x][i*6+y] = m[4-x][4-y]
				} else {
					d.charBuffer[x][i*6+y] = m[4-y][x]
				}
			}
		}
	}
}

func (d *Display) doMatrixBuffer() {
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			d.buffer[matrixSettings[d.direction][x][y][0]][matrixSettings[d.direction][x][y][1]] = d.charBuffer[x][d.step+y]
		}
	}
}
