package mocks

import (
	"testdoubles/simulator"

	"github.com/stretchr/testify/mock"
)

type SimulatorMock struct {
	mock.Mock
}

// CanCatch implements simulator.CatchSimulator.
func (s *SimulatorMock) CanCatch(hunter, prey *simulator.Subject) (canCatch bool) {
	args := s.Called(hunter, prey)
	return args.Bool(0)
}
