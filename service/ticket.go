package service

import (
	"encoding/json"
	"errors"
	"pengecekan-tiket-online/data"
	"pengecekan-tiket-online/dto"
	"strings"
)

type TicketService struct{}

// constructor
func NewTicketService() TicketService {
	return TicketService{}
}

// method untuk mendapatkan informasi tiket
func (ts *TicketService) TicketInfo(req dto.RequestDTO) (dto.ResponseDTO, error) {

	// Validasi input
	if strings.TrimSpace(req.Destination) == "" || strings.TrimSpace(req.Name) == "" {
		return dto.ResponseDTO{}, errors.New("destination dan name tidak boleh kosong")
	}

	// Parse data destinations dari JSON string
	var destinations map[string]int
	err := json.Unmarshal([]byte(data.Destinations), &destinations)
	if err != nil {
		return dto.ResponseDTO{}, errors.New("gagal membaca data destinasi")
	}

	// Cari harga berdasarkan destinasi (case-insensitive)
	var price int
	var foundDestination string
	for dest, p := range destinations {
		if strings.EqualFold(dest, req.Destination) {
			price = p
			foundDestination = dest
			break
		}
	}

	// Jika destinasi tidak ditemukan
	if foundDestination == "" {
		return dto.ResponseDTO{}, errors.New("destinasi tidak ditemukan")
	}

	// Buat response
	response := dto.ResponseDTO{
		Name:        req.Name,
		Destination: foundDestination,
		Price:       price,
	}

	return response, nil
}

// method untuk mendapatkan list semua destinasi yang tersedia
func (ts *TicketService) ListDestinations() (map[string]int, error) {

	// Parse data destinations dari JSON string
	var destinations map[string]int
	err := json.Unmarshal([]byte(data.Destinations), &destinations)
	if err != nil {
		return nil, errors.New("gagal membaca data destinasi")
	}

	return destinations, nil
}