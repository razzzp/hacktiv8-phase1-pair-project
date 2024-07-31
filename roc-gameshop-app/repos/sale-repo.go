package repos

import (
	"database/sql"
	"fmt"
	"roc-gameshop-app/entities"
)

type SaleRepo interface {
	GetAllSales() ([]entities.Sale, error)
	CreateSale(sale entities.Sale) error
	UpdateSale(id int, sale entities.Sale) error
	GetSaleById(id int) (*entities.Sale, error)
}

type saleRepo struct {
	db *sql.DB
}

// Create Sale
func (s *saleRepo) CreateSale(sale entities.Sale) error {
	query := `
	INSERT INTO sales (UserId, GameId, PriceId, StockId)
	VALUES (?,?,?,?)`

	_, err := s.db.Exec(query, sale.UserId, sale.GameId, sale.PriceId, sale.StockId)
	if err != nil {
		fmt.Println("Error executing create sale query")
		return err
	}
	fmt.Printf("Success creating sale for UserId %d and GameId %d\n", sale.UserId, sale.GameId)
	return nil
}

// Get All Sales
func (s *saleRepo) GetAllSales() ([]entities.Sale, error) {
	query := `SELECT * FROM sales`

	rows, err := s.db.Query(query)
	if err != nil {
		fmt.Println("Error executing get all sales query")
		return nil, err
	}
	defer rows.Close()

	sales := []entities.Sale{}
	for rows.Next() {
		sale := entities.Sale{}
		err := rows.Scan(&sale.SaleId, &sale.UserId, &sale.GameId, &sale.PriceId, &sale.StockId)
		if err != nil {
			fmt.Println("Error scanning returned sales data")
			return nil, err
		}
		sales = append(sales, sale)
	}

	return sales, nil
}

// Update Sale
func (s *saleRepo) UpdateSale(id int, sale entities.Sale) error {
	query := `
		UPDATE sales
		SET UserId = ?, GameId = ?, PriceId = ?, StockId = ?
		WHERE SaleId = ?
	`
	_, err := s.db.Exec(query, sale.UserId, sale.GameId, sale.PriceId, sale.StockId, id)
	if err != nil {
		fmt.Println("Error executing update sale query")
		return err
	}
	fmt.Println("Success updating sale")
	return nil
}

// Get Sale By ID
func (s *saleRepo) GetSaleById(id int) (*entities.Sale, error) {
	query := `
		SELECT * FROM sales WHERE SaleId = ?
	`

	row := s.db.QueryRow(query, id)
	sale := entities.Sale{}
	err := row.Scan(&sale.SaleId, &sale.UserId, &sale.GameId, &sale.PriceId, &sale.StockId)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No sale found with the given ID")
			return nil, nil
		}
		fmt.Println("Error executing get sale by id query")
		return nil, err
	}
	return &sale, nil
}

func NewSaleRepo(db *sql.DB) SaleRepo {
	return &saleRepo{db}
}
