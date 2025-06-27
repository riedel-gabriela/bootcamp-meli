package repository

import (
	"context"
	"sort"

	"github.com/riedel-gabriela/desafio-goweb/internal"
)

// NewRepositoryTicketMap creates a new repository for tickets in a map
func NewRepositoryTicketMap(db map[int]internal.TicketAttributes) *RepositoryTicketMap {
	defaultDb := make(map[int]internal.TicketAttributes)
	if db != nil {
		defaultDb = db
	}
	lastid := defineLastId(defaultDb)
	return &RepositoryTicketMap{
		db:     defaultDb,
		lastId: lastid,
	}
}

// RepositoryTicketMap implements the repository interface for tickets in a map
type RepositoryTicketMap struct {
	// db represents the database in a map
	// - key: id of the ticket
	// - value: ticket
	db map[int]internal.TicketAttributes

	// lastId represents the last id of the ticket
	lastId int
}

func defineLastId(db map[int]internal.TicketAttributes) (lastId int) {
	keys := make([]int, 0, len(db))
	for k := range db {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	if len(keys) == 0 {
		return 0
	}
	return keys[len(keys)-1]
}

func (r *RepositoryTicketMap) GetTotalAmountTickets() (int, error) {
	return len(r.db), nil
}

func (r *RepositoryTicketMap) Get(ctx context.Context) (t map[int]internal.TicketAttributes, err error) {
	t = make(map[int]internal.TicketAttributes, len(r.db))
	for k, v := range r.db {
		t[k] = v
	}
	return
}

func (r *RepositoryTicketMap) GetTicketsByDestinationCountry(ctx context.Context, country string) (t map[int]internal.TicketAttributes, err error) {
	t = make(map[int]internal.TicketAttributes)
	for k, v := range r.db {
		if v.Country == country {
			t[k] = v
		}
	}
	return
}
