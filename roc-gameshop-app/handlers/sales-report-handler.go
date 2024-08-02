package handlers

import (
	"database/sql"
	"fmt"
)

type SalesReportRow struct {
	GameName      string
	TotalQuantity int
	TotalSales    float64
}

type SalesReportHandler interface {
	GetSalesReport() ([]*SalesReportRow, error)
}

type salesReportHandler struct {
	db *sql.DB
}

func NewSalesReportHandler(db *sql.DB) SalesReportHandler {
	return &salesReportHandler{
		db: db,
	}
}

// GetSalesReport implements SalesReportHandler.
func (s *salesReportHandler) GetSalesReport() ([]*SalesReportRow, error) {
	query :=
		`SELECT
			g.Name AS Game,
			SUM(s.Quantity) AS QuantityPurchased,
			SUM(s.PurchasedPrice) AS TotalSales
		FROM Sales s
		INNER JOIN Games g ON s.GameId = g.GameId
		GROUP BY s.GameId
		ORDER BY TotalSales DESC;`

	rows, err := s.db.Query(query)
	if err != nil {
		fmt.Println("Error executing get all rentals query")
		return nil, err
	}
	result := []*SalesReportRow{}
	for rows.Next() {
		salesReportRow := SalesReportRow{}
		err := rows.Scan(&salesReportRow.GameName, &salesReportRow.TotalQuantity, &salesReportRow.TotalSales)
		if err != nil {
			fmt.Println("Error scanning returned rentals data")
			return nil, err
		}
		result = append(result, &salesReportRow)
	}

	return result, nil
}
