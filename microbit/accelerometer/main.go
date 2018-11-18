package main

import (
	"time"

	"github.com/conejoninja/tinygoexamples/microbit/accelerometer/MMA8653"
	"github.com/aykevl/tinygo/src/machine"
)



func main() {

	machine.I2C0.Configure(machine.I2CConfig{})

	println("Starting accelerometer")
	machine.I2C0.WriteRegister(MMA8653.I2C_ADDR, MMA8653.CTRL_REG2, []byte{MMA8653.CTRL_REG2_RESET})
	time.Sleep(time.Millisecond*10)
	standby()
	machine.I2C0.WriteRegister(MMA8653.I2C_ADDR, MMA8653.XYZ_DATA_CFG, []byte{MMA8653.MODE_2G})
	active()

	machine.InitLEDMatrix()
	for {

		x, y, _ := getMMA8653Data()
		col := 2-x/20
		if col > 4 {
			col = 4
		}
		if col < 0 {
			col = 0
		}
		row := 2-y/20
		if row > 4 {
			row = 4
		}
		if row < 0 {
			row = 0
		}

		machine.SetLEDMatrix(uint8(col), uint8(row))
		time.Sleep(time.Millisecond * 100)
		getMMA8653Data()
	}

}

func getMMA8653Data() (x,y,z int){
	buf := make([]byte, 4)
	machine.I2C0.ReadRegister(MMA8653.I2C_ADDR, 0x00, buf)
	x = int(buf[1])
	y = int(buf[2])
	z = int(buf[3])
	return
}

func standby() {
	buf := make([]byte, 1)
	machine.I2C0.ReadRegister(MMA8653.I2C_ADDR, MMA8653.CTRL_REG1, buf)
	machine.I2C0.WriteRegister(MMA8653.I2C_ADDR, MMA8653.CTRL_REG1, []byte{buf[0] &^ MMA8653.CTRL_REG1_VALUE_ACTIVE})
}

func active() {
	buf := make([]byte, 1)
	machine.I2C0.ReadRegister(MMA8653.I2C_ADDR, 0x00, buf)
	machine.I2C0.WriteRegister(MMA8653.I2C_ADDR, MMA8653.CTRL_REG2, []byte{0x09})
	value := buf[0] | MMA8653.CTRL_REG1_VALUE_ACTIVE | MMA8653.CTRL_REG1_VALUE_F_READ | MMA8653.ODR_6_25
	machine.I2C0.WriteRegister(MMA8653.I2C_ADDR, MMA8653.CTRL_REG1, []byte{value})
}


