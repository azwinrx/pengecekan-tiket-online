package main

import (
    "pengecekan-tiket-online/dto"
    "pengecekan-tiket-online/handler"
    "pengecekan-tiket-online/service"
)



func main() {
    ticketService := service.NewTicketService()

    ticketHandler := handler.NewTicketHandler(ticketService)

    req := dto.RequestDTO{
        Name:        "Budi",
        Destination: "Jakarta",
    }

    ticketHandler.CheckTicket(req)
}