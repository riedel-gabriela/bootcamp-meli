package hunter

import (
	"testdoubles/mocks"
	"testdoubles/positioner"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHunterWhiteShark(t *testing.T) {
	t.Run("CreateWhiteShark", func(t *testing.T) {
		// given
		simulator := mocks.SimulatorMock{}

		// when
		shark := CreateWhiteShark(&simulator)

		// then
		assert.NotNil(t, shark, "Expected a White Shark instance, got nil")
		assert.IsType(t, &WhiteShark{}, shark)
	})

	t.Run("Hunt_Success", func(t *testing.T) {
		// given
		simulator := mocks.SimulatorMock{}
		simulator.On("CanCatch", mock.Anything, mock.Anything).Return(true)
		shark := CreateWhiteShark(&simulator)
		prey := mocks.PreyMock{}
		// Add these lines to set up PreyMock expectations
		prey.On("GetPosition").Return(&positioner.Position{X: 10, Y: 20})
		prey.On("GetSpeed").Return(5.0)

		// when
		result := shark.Hunt(&prey)

		// then
		assert.NoError(t, result)
		simulator.AssertExpectations(t)
		prey.AssertExpectations(t)
		simulator.AssertCalled(t, "CanCatch", mock.Anything, mock.Anything)
		prey.AssertCalled(t, "GetPosition")
		prey.AssertCalled(t, "GetSpeed")
	})

	t.Run("Hunt_Failure", func(t *testing.T) {
		// given
		simulator := mocks.SimulatorMock{}
		simulator.On("CanCatch", mock.Anything, mock.Anything).Return(false)
		shark := CreateWhiteShark(&simulator)
		prey := mocks.PreyMock{}
		prey.On("GetPosition").Return(&positioner.Position{X: 10, Y: 20})
		prey.On("GetSpeed").Return(5.0)
		// when
		result := shark.Hunt(&prey)

		// then
		assert.Error(t, result)
		// única forma que a aplicação retorna um erro é quando o caçador
		// não consegue pegar a presa. Caso a presa seja nula, o erro é nil
		assert.Equal(t, "white shark cannot catch the prey", result.Error())
		simulator.AssertExpectations(t)
		prey.AssertExpectations(t)
	})
}
