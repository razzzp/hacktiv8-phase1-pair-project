package cli

import (
	"bufio"
	"fmt"
	"roc-gameshop-app/handlers"

	"github.com/rodaine/table"
)

type reviewReportCli struct {
	router              Router
	reader              *bufio.Reader
	reviewReportHandler handlers.ReviewHandler
}

func NewReviewReportCli(
	router Router,
	reader *bufio.Reader,
	reviewReportHandler handlers.ReviewHandler,
) Cli {
	return &reviewReportCli{
		router:              router,
		reader:              reader,
		reviewReportHandler: reviewReportHandler,
	}
}

func (r *reviewReportCli) GetUserActions(session *Session) []Action {
	result := []Action{}
	result = append(result, Action{Name: "Back", ActionFunc: func() {
		r.router.Pop()
	}})
	return result
}

func (r *reviewReportCli) HandleRoute(args RouteArgs, session *Session) {

	fmt.Println("Average Ratings Per Game Report")
	fmt.Println("")

	// Generate table
	ratings, err := r.reviewReportHandler.GetAvgRatings()
	if err != nil {
		fmt.Printf("Error generating report: %v", err)
	} else {
		reportTable := table.New("No.", "Game Name", "Average Ratings")
		for i, row := range ratings {
			reportTable.AddRow(i+1, row.GameName, row.AvgRating).WithPadding(2)
		}
		reportTable.Print()
	}
	fmt.Println("")
	// get user actions
	actions := r.GetUserActions(session)
	PromptUserForActions(actions, r.reader)
}
