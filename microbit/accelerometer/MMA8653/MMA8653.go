package MMA8653

const (
	I2C_ADDR = 0x1D

	CTRL_REG1     = 0x2A
	CTRL_REG1_VALUE_ACTIVE = 0x01
	CTRL_REG1_VALUE_F_READ = 0x02

	CTRL_REG2       = 0x2B
	CTRL_REG2_RESET = 0x40

	CTRL_REG3          = 0x2C
	CTRL_REG3_VALUE_OD = 0x01

	CTRL_REG4                = 0x2D
	CTRL_REG4_VALUE_INT_ASLP = 0x80
	CTRL_REG4_VALUE_INT_ENLP = 0x10
	CTRL_REG4_VALUE_INT_FFMT = 0x04
	CTRL_REG4_VALUE_INT_DRDY = 0x01

	CTRL_REG5 = 0x2E // 1: routed to INT1

	PL_STATUS = 0x10
	PL_CFG    = 0x11
	PL_EN     = 0x40

	XYZ_DATA_CFG = 0x0E
	MODE_2G      = 0x00 //Set Sensitivity to 2g
	MODE_4G      = 0x01 //Set Sensitivity to 4g
	MODE_8G      = 0x02 //Set Sensitivity to 8g

	FF_MT_CFG     = 0x15
	FF_MT_CFG_ELE = 0x80
	FF_MT_CFG_OAE = 0x40
	FF_MT_CFG_XYZ = 0x38

	FF_MT_SRC    = 0x16
	FF_MT_SRC_EA = 0x80

	FF_MT_THS = 0x17

	FF_MT_COUNT = 0x18

	PULSE_CFG     = 0x21
	PULSE_CFG_ELE = 0x80

	PULSE_SRC    = 0x22
	PULSE_SRC_EA = 0x80

	// Sample rate
	ODR_800  = 0x00
	ODR_400  = 0x08
	ODR_200  = 0x10
	ODR_100  = 0x18 // default ratio 100 samples per second
	ODR_50   = 0x20
	ODR_12_5 = 0x28
	ODR_6_25 = 0x30
	ODR_1_56 = 0x38
)
