package cli

import (
	"bufio"
	"fmt"
	"roc-gameshop-app/routes"
)

type homePageCli struct {
	router Router
	reader *bufio.Reader
}

func NewHomepageCli(router Router, reader *bufio.Reader) Cli {
	return &homePageCli{
		router: router,
		reader: reader,
	}
}

func (hpc *homePageCli) HandleRoute(args RouteArgs, session *Session) {
	// logic of game details page goes here
	// TODO
	fmt.Println("Welcome to ROC Gameshop")
	fmt.Println("")

	fmt.Println("Actions")
	fmt.Println("1. Search Games")
	fmt.Println("2. View Cart")
	fmt.Println("3. Login")
	fmt.Println("4. Register")
	fmt.Println("5. Exit")

	fmt.Println("")

	// temp for testing
	for {
		input, err := PromptUserForAction(hpc.reader)
		if err != nil {
			fmt.Printf("Invalid input: %s, please try again.\n", err)
			continue
		}
		switch input {
		case 1:
			// TODO
			return
		case 2:
			// TODO
			return
		case 3:
			// TODO
			return
		case 4:
			hpc.router.Push(routes.REGISTER_ROUTE, RouteArgs{})
			return
		default:
			// TODO
			return
		}
	}

}
