package cli

import (
	"bufio"
	"fmt"
	"roc-gameshop-app/handlers"

	"github.com/rodaine/table"
)

type overdueCli struct {
	rentalHandler handlers.RentalHandler
	router        Router
	reader        *bufio.Reader
}

func NewOverdueCli(router Router, reader *bufio.Reader, rentalHandler handlers.RentalHandler) Cli {
	return &overdueCli{
		router:        router,
		reader:        reader,
		rentalHandler: rentalHandler,
	}
}

func (oC *overdueCli) HandleRoute(args RouteArgs, session *Session) {
	fmt.Println("Rentals Overdue Report")
	rentals, err := oC.rentalHandler.GetOverdues()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("")
	fmt.Printf("There are %d rentals that already overdue\n", len(rentals))
	fmt.Println("")
	overduesTable := table.New("User Name", "Game Name", "Start Date", "End Date", "Status")
	for _, rental := range rentals {
		overduesTable.AddRow(rental.UserName, rental.GameName, rental.StartDate, rental.EndDate, rental.Status)
	}
	overduesTable.Print()

	fmt.Println("")
	actions := []Action{
		{
			Name: "Back",
			ActionFunc: func() {
				oC.router.Pop()
				return
			},
		},
	}

	PromptUserForActions(actions, oC.reader)
}
