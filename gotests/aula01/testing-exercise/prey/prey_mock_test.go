package prey

import (
	"testdoubles/mocks"
	"testdoubles/positioner"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPreyMock(t *testing.T) {
	t.Run("GetPosition", func(t *testing.T) {
		// given
		prey := &mocks.PreyMock{}
		expectedPosition := &positioner.Position{}
		prey.On("GetPosition").Return(expectedPosition)

		result := prey.GetPosition()

		assert.NotNil(t, result, "Expected a position, got nil")
	})

	t.Run("GetSpeed", func(t *testing.T) {
		// given
		prey := &mocks.PreyMock{}
		expectedSpeed := 10.0
		prey.On("GetSpeed").Return(expectedSpeed)

		// when
		result := prey.GetSpeed()

		// then
		assert.Equal(t, expectedSpeed, result, "Expected speed to be 10.0, got %f", result)
	})
}
