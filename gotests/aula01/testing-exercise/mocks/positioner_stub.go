package mocks

import "math"

// PositionerStub é um stub que implementa a interface Positioner com comportamento fixo
type PositionerStub struct {
	// Pode ter campos para controlar o comportamento se necessário
}

// GetLinearDistance calcula a distância euclidiana real entre duas posições
func (p *PositionerStub) GetLinearDistance(from, to *PositionMock) float64 {
	if from == nil || to == nil {
		return 0.0
	}

	dx := to.X - from.X
	dy := to.Y - from.Y
	dz := to.Z - from.Z

	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

// PositionerStubAlwaysZero é um stub que sempre retorna zero (para testes específicos)
type PositionerStubAlwaysZero struct{}

func (p *PositionerStubAlwaysZero) GetLinearDistance(from, to *PositionMock) float64 {
	return 0.0
}

// PositionerStubFixed é um stub que sempre retorna um valor fixo
type PositionerStubFixed struct {
	FixedDistance float64
}

func (p *PositionerStubFixed) GetLinearDistance(from, to *PositionMock) float64 {
	return p.FixedDistance
}
