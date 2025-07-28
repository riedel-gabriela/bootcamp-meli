package positioner_test

import (
	"testdoubles/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPositioner(t *testing.T) {
	t.Run("GetLinearDistance_RealCalculation", func(t *testing.T) {
		// given
		positioner := &mocks.PositionerStub{} // stub do positioner
		position1 := &mocks.PositionMock{X: 0, Y: 0, Z: 0}
		position2 := &mocks.PositionMock{X: 3, Y: 4, Z: 0}

		// when
		distance := positioner.GetLinearDistance(position1, position2)

		// then
		assert.Equal(t, 5.0, distance, "Expected linear distance to be 5.0 (3,4,0 triangle)")
	})

	t.Run("GetLinearDistance_SamePosition", func(t *testing.T) {
		// given
		positioner := &mocks.PositionerStub{}
		position1 := &mocks.PositionMock{X: 1, Y: 2, Z: 3}
		position2 := &mocks.PositionMock{X: 1, Y: 2, Z: 3}

		// when
		distance := positioner.GetLinearDistance(position1, position2)

		// then
		assert.Equal(t, 0.0, distance, "Expected linear distance to be 0.0 for same positions")
	})

	t.Run("GetLinearDistance_3DDistance", func(t *testing.T) {
		// given
		positioner := &mocks.PositionerStub{}
		position1 := &mocks.PositionMock{X: 1, Y: 1, Z: 1}
		position2 := &mocks.PositionMock{X: 4, Y: 5, Z: 1}

		// when
		distance := positioner.GetLinearDistance(position1, position2)

		// then
		expectedDistance := 5.0 // sqrt((4-1)² + (5-1)² + (1-1)²) = sqrt(9 + 16 + 0) = 5
		assert.Equal(t, expectedDistance, distance, "Expected linear distance to be 5.0")
	})

	t.Run("GetLinearDistance_NilPositions", func(t *testing.T) {
		// given
		positioner := &mocks.PositionerStub{}

		// when
		distance := positioner.GetLinearDistance(nil, nil)

		// then
		assert.Equal(t, 0.0, distance, "Expected linear distance to be 0.0 for nil positions")
	})

	t.Run("GetLinearDistance_AlwaysZeroStub", func(t *testing.T) {
		// given
		positioner := &mocks.PositionerStubAlwaysZero{}
		position1 := &mocks.PositionMock{X: 100, Y: 200, Z: 300}
		position2 := &mocks.PositionMock{X: 1, Y: 2, Z: 3}

		// when
		distance := positioner.GetLinearDistance(position1, position2)

		// then
		assert.Equal(t, 0.0, distance, "AlwaysZero stub should always return 0.0")
	})

	t.Run("GetLinearDistance_FixedValueStub", func(t *testing.T) {
		// given
		fixedValue := 42.5
		positioner := &mocks.PositionerStubFixed{FixedDistance: fixedValue}
		position1 := &mocks.PositionMock{X: 0, Y: 0, Z: 0}
		position2 := &mocks.PositionMock{X: 100, Y: 100, Z: 100}

		// when
		distance := positioner.GetLinearDistance(position1, position2)

		// then
		assert.Equal(t, fixedValue, distance, "Fixed stub should return the predefined value")
	})
}
