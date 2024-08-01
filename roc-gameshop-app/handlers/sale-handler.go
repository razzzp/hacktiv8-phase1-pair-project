package handlers

import (
	"fmt"
	"roc-gameshop-app/entities"
	"roc-gameshop-app/repos"
)

type SaleDto struct {
	SaleId        int
	GameId        int
	UserId        int
	SaleDate      string
	PurchasePrice string
	Quantity      string
}

type SaleHandler interface {
	GetAll() ([]*entities.Sale, error)
	GetById(id int) (*entities.Sale, error)
	Create(rental *entities.Sale) error
}

type saleHandler struct {
	saleRepo repos.SaleRepo
}

// // validate dto
// func (s *saleHandler) ValidateDto(saleDto *SaleDto) (*entities.Sale, error) {
// 	if saleDto == nil {
// 		return nil, errors.New("sale dto is nil")
// 	}
// 	result := entities.Sale{}
// 	asInt, err := strconv.Atoi(saleDto.GameId)
// }

// Create implements SaleHandler.
func (s *saleHandler) Create(sale *entities.Sale) error {
	return s.saleRepo.CreateSale(sale)
}

// GetAll implements SaleHandler.
func (s *saleHandler) GetAll() ([]*entities.Sale, error) {
	sales, err := s.saleRepo.GetAllSales()
	if err != nil {
		fmt.Println("error getting All Rentals")
		return nil, err
	}
	return sales, nil
}

// GetById implements SaleHandler.
func (s *saleHandler) GetById(id int) (*entities.Sale, error) {
	sale, err := s.saleRepo.GetSaleById(id)
	if err != nil {
		fmt.Println("error get a Rental")
		return nil, err
	}
	return sale, nil
}

func NewSaleHandler(saleRepo repos.SaleRepo) SaleHandler {
	return &saleHandler{
		saleRepo: saleRepo,
	}
}
