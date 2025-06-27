package service_test

import (
	"context"
	"testing"

	"github.com/riedel-gabriela/desafio-goweb/internal"
	"github.com/riedel-gabriela/desafio-goweb/internal/service"
	"github.com/stretchr/testify/require"
)

// Mock ajustado para implementar GetTotalAmountTickets
type ticketRepoMock struct {
	GetTotalAmountTicketsFunc func() (int, error)
}

func (m *ticketRepoMock) GetTotalAmountTickets() (int, error) {
	return m.GetTotalAmountTicketsFunc()
}

// Os outros métodos podem ser implementados se necessário
func (m *ticketRepoMock) Get(ctx context.Context) (map[int]internal.TicketAttributes, error) {
	return nil, nil
}
func (m *ticketRepoMock) GetTicketsByDestinationCountry(ctx context.Context, country string) (map[int]internal.TicketAttributes, error) {
	return nil, nil
}

func TestTicketService_GetTotalAmountTickets(t *testing.T) {
	t.Run("success to get total tickets", func(t *testing.T) {
		// arrange
		mockRepo := &ticketRepoMock{
			GetTotalAmountTicketsFunc: func() (int, error) {
				return 1, nil
			},
		}
		sv := service.NewTicketService(mockRepo)

		// act
		total, err := sv.GetTotalAmountTickets()

		// assert
		expectedTotal := 1
		require.NoError(t, err)
		require.Equal(t, expectedTotal, total)
	})
}
