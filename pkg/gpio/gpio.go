package gpio

import (
	"fmt"
	"os/exec"
)

// PinMode 引脚模式
type PinMode string

// PinLevel 电平水平
type PinLevel int8

const (
	IN   PinMode  = "in"
	OUT  PinMode  = "out"
	UP   PinLevel = 1
	DOWN PinLevel = 0
)

type Pin struct {
	Number int8
	Mode   PinMode
}

const (
	direction string = "echo %s > /sys/class/gpio/gpio%d/direction"
	export    string = "echo %d > /sys/class/gpio/export"
	value     string = "echo %d > /sys/class/gpio/gpio%d/value"
	unexport  string = "echo %d > /sys/class/gpio/unexport"
)

func (p *Pin) SetLevel(level PinLevel) error {
	cmd := fmt.Sprintf(value, level, p.Number)
	if err := exec.Command(cmd).Start(); err != nil {
		return err
	}
	return nil
}

func (p *Pin) Close() error {
	cmd := fmt.Sprintf(unexport, p.Number)
	if err := exec.Command(cmd).Start(); err != nil {
		return err
	}
	return nil
}
