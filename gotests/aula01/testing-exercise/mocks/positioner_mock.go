package mocks

import "github.com/stretchr/testify/mock"

// Position represents a point in space (define fields as needed)
type PositionMock struct {
	X float64
	Y float64
	Z float64
}

type PositionerMock struct {
	mock.Mock
}

func (p *PositionerMock) GetLinearDistance(from, to *PositionMock) (linearDistance float64) {
	args := p.Called(from, to)
	if args.Get(0) == nil {
		return 0.0
	}
	return args.Get(0).(float64)
}
