package handlers

import (
	"net/http"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	"github.com/riedel-gabriela/desafio-goweb/internal"
)

func NewTicketDefautl(sv internal.TicketService) *TicketDefault {
	return &TicketDefault{sv: sv}
}

type TicketDefault struct {
	sv internal.TicketService
}

func (h *TicketDefault) GetTotalAmountTickets() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		total, err := h.sv.GetTotalAmountTickets()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"total": total,
		})
	}
}

func (h *TicketDefault) GetByCountry() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		country := chi.URLParam(r, "country")
		tickets, err := h.sv.GetByCountry(r.Context(), country)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if len(tickets) == 0 {
			http.Error(w, "No tickets found for the specified country", http.StatusNotFound)
			return
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"country": country,
			"total":   len(tickets),
			"tickets": tickets,
		})
	}
}

func (h *TicketDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tickets, err := h.sv.GetAll(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if len(tickets) == 0 {
			http.Error(w, "No tickets found", http.StatusNotFound)
			return
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"total":   len(tickets),
			"tickets": tickets,
		})
	}
}
