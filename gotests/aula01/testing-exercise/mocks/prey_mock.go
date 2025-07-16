package mocks

import (
	"testdoubles/positioner"

	"github.com/stretchr/testify/mock"
)

type PreyMock struct {
	mock.Mock
}

// GetPosition implements prey.Prey.
func (p *PreyMock) GetPosition() (position *positioner.Position) {
	args := p.Called()
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*positioner.Position)
}

// GetSpeed implements prey.Prey.
func (p *PreyMock) GetSpeed() (speed float64) {
	args := p.Called()
	return args.Get(0).(float64)
}
