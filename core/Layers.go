package core

import "math"

type Layers struct {
	Mask uint32
}

func NewLayers() *Layers {
	return &Layers{1}
}

func (l *Layers) Set(channel int) {
	l.Mask = 1 << uint32(channel)
}

func (l *Layers) Enable(channel int) {
	l.Mask |= 1 << uint32(channel)
}

func (l *Layers) EnableAll() {
	l.Mask = math.MaxUint32
}

func (l *Layers) Toggle(channel int) {
	l.Mask ^= 1 << uint32(channel)
}

func (l *Layers) Disable(channel int) {
	l.Mask &= ^(1 << uint32(channel))
}

func (l *Layers) DisableAll() {
	l.Mask = 0
}

func (l *Layers) Test(layers *Layers) bool {
	return (l.Mask & layers.Mask) != 0
}

func (l *Layers) IsEnabled(channel int) bool {
	return (l.Mask & (1 << uint32(channel))) != 0
}
