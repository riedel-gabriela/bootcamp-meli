package internal

import "context"

type TicketService interface {
	GetAll(ctx context.Context) (map[int]TicketAttributes, error)
	GetByCountry(ctx context.Context, country string) (map[int]TicketAttributes, error)
	GetTotalAmountTickets() (total int, err error)
}
