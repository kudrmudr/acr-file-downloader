package downloader

import (
	"math"
	"sync"
)

func NewMinPosition() MinPosition {
	return MinPosition{value: math.MaxUint64}
}

type MinPosition struct {
	sync.RWMutex
	value uint64
}

func (p *MinPosition) Get() uint64 {
	p.RLock()
	defer p.RUnlock()
	return p.value
}

func (p *MinPosition) Set(v uint64) {
	p.Lock()
	if v < p.value {
		p.value = v
	}
	p.Unlock()
}
