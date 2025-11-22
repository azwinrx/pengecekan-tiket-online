package handler

import (
	"encoding/json"
	"fmt"
	"pengecekan-tiket-online/dto"
	"pengecekan-tiket-online/service"
)

type TicketHandler struct {
	TicketService service.TicketService
}

// constructor
func NewTicketHandler(ticketService service.TicketService) TicketHandler {
	return TicketHandler{
		TicketService: ticketService,
	}
}

// checking method
func (h *TicketHandler) CheckTicket(req dto.RequestDTO) {
	ticket, err := h.TicketService.TicketInfo(req)

	if err != nil {
		fmt.Println("Error, soalnya:", err.Error())
		return
	}
	fmt.Println("")
	fmt.Println("=== Harga Tiket ===")
	fmt.Printf("Penumpang  : %s\n", ticket.Name)
	fmt.Printf("Tujuan	   : %s\n", ticket.Destination)
	fmt.Printf("Harga	   : Rp %d\n", ticket.Price)
	fmt.Println("====================")

}

func (tickethandler *TicketHandler) ListTicket(){
	tickets, err := tickethandler.TicketService.ListDestinations()
	if err != nil {
		fmt.Println("Error, soalnya:", err.Error())
		return
	}

	data, err := json.MarshalIndent(tickets, "", "  ")
	if err != nil {
		fmt.Println("Error di marshalling:", err.Error())
		return
	}
	fmt.Println(string(data))
}