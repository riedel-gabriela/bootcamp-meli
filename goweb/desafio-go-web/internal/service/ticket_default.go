package service

import (
	"context"

	"github.com/riedel-gabriela/desafio-goweb/internal"
)

type TicketService struct {
	rp internal.TicketRepository
}

func NewTicketService(repo internal.TicketRepository) *TicketService {
	return &TicketService{rp: repo}
}

func (s *TicketService) GetTotalAmountTickets() (total int, err error) {
	return s.rp.GetTotalAmountTickets()
}

func (s *TicketService) GetAll(ctx context.Context) (map[int]internal.TicketAttributes, error) {
	return s.rp.Get(ctx)
}

func (s *TicketService) GetByCountry(ctx context.Context, country string) (map[int]internal.TicketAttributes, error) {
	return s.rp.GetTicketsByDestinationCountry(ctx, country)
}
