package internal

import "context"

type TicketRepository interface {
	Get(ctx context.Context) (t map[int]TicketAttributes, err error)
	GetTotalAmountTickets() (total int, err error)
	GetTicketsByDestinationCountry(ctx context.Context, country string) (t map[int]TicketAttributes, err error)
}
