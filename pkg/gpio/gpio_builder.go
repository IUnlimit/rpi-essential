package gpio

import (
	"fmt"
	"os/exec"
	"sort"
)

type PinBuilder struct {
	flow []*Flow
}

type Flow struct {
	index int8
	opt   func(p *Pin) error
}

func Builder() *PinBuilder {
	return &PinBuilder{}
}

// Init the pin
func (b *PinBuilder) Init(pin int8) *PinBuilder {
	b.flow = append(b.flow, &Flow{
		index: 0,
		opt: func(p *Pin) error {
			cmd := fmt.Sprintf(export, pin)
			if err := exec.Command(cmd).Start(); err != nil {
				return err
			}
			p.Number = pin
			return nil
		},
	})
	return b
}

// SetMode set pin mode
func (b *PinBuilder) SetMode(mode PinMode) *PinBuilder {
	b.flow = append(b.flow, &Flow{
		index: 1,
		opt: func(p *Pin) error {
			cmd := fmt.Sprintf(direction, mode, p.Number)
			if err := exec.Command(cmd).Start(); err != nil {
				return err
			}
			p.Mode = mode
			return nil
		},
	})
	return b
}

func (b *PinBuilder) Build() (*Pin, error) {
	pin := &Pin{}
	sort.Slice(b.flow, func(i, j int) bool {
		return b.flow[i].index < b.flow[j].index
	})
	for _, f := range b.flow {
		if err := f.opt(pin); err != nil {
			return nil, err
		}
	}
	return pin, nil
}
