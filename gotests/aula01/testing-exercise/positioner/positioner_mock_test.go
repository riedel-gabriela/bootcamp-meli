package positioner_test

// PositionerDefaultStubTest is a test suite for the default stub implementation of the Positioner interface.
import (
	"testdoubles/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPositionerMock(t *testing.T) {
	t.Run("GetLinearDistance_SamePosition", func(t *testing.T) {
		// given
		positioner := mocks.PositionerMock{} // mock do positioner
		position1 := &mocks.PositionMock{X: 1, Y: 2, Z: 3}
		position2 := &mocks.PositionMock{X: 1, Y: 2, Z: 3}

		// Set up mock expectation
		positioner.On("GetLinearDistance", position1, position2).Return(0.0)

		// when
		distance := positioner.GetLinearDistance(position1, position2)

		// then
		assert.Equal(t, 0.0, distance, "Expected linear distance to be 0.0 for same positions")
		positioner.AssertExpectations(t)
	})

	t.Run("GetLinearDistance_DifferentPositions", func(t *testing.T) {
		// given
		positioner := mocks.PositionerMock{}
		position1 := &mocks.PositionMock{X: 0, Y: 0, Z: 0}
		position2 := &mocks.PositionMock{X: 3, Y: 4, Z: 0}
		expectedDistance := 5.0 // sqrt(3^2 + 4^2) = 5

		// Set up mock expectation
		positioner.On("GetLinearDistance", position1, position2).Return(expectedDistance)

		// when
		distance := positioner.GetLinearDistance(position1, position2)

		// then
		assert.Equal(t, expectedDistance, distance, "Expected linear distance to be 5.0")
		positioner.AssertExpectations(t)
	})
}
