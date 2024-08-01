package cli

import (
	"bufio"
	"fmt"
	"roc-gameshop-app/handlers"

	"github.com/rodaine/table"
)

type salesReportCli struct {
	router             Router
	reader             *bufio.Reader
	salesReportHandler handlers.SalesReportHandler
}

func NewSalesReportCli(
	router Router,
	reader *bufio.Reader,
	salesReportHandler handlers.SalesReportHandler,
) Cli {
	return &salesReportCli{
		router:             router,
		reader:             reader,
		salesReportHandler: salesReportHandler,
	}
}

func (s *salesReportCli) GetUserActions(session *Session) []Action {
	result := []Action{}
	result = append(result, Action{Name: "Back", ActionFunc: func() {
		s.router.Pop()
	}})
	return result
}

func (s *salesReportCli) HandleRoute(args RouteArgs, session *Session) {

	fmt.Println("Sales Report")
	fmt.Println("")

	// Generate table
	salesReport, err := s.salesReportHandler.GetSalesReport()
	if err != nil {
		fmt.Printf("Error generating report: %v", err)
	} else {
		reportTable := table.New("No.", "Game", "Total Quantity Sold", "Total Revenue")
		for i, row := range salesReport {
			reportTable.AddRow(i+1, row.GameName, row.TotalQuantity, FormatAsCurrency(row.TotalSales)).WithPadding(2)
		}
		reportTable.Print()
	}
	fmt.Println("")
	// get user actions
	actions := s.GetUserActions(session)
	PromptUserForActions(actions, s.reader)
}
