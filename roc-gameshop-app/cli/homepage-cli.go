package cli

import (
	"bufio"
	"fmt"
	"roc-gameshop-app/routes"
	"strings"
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

func (hpc *homePageCli) GetUserActions(session *Session) []Action {
	result := []Action{}
	result = append(result, Action{Name: "Search Games", ActionFunc: func() {
		//get game name to search
		fmt.Printf("Enter game name to search: ")
		name, err := hpc.reader.ReadString('\n')
		if err != nil {
			fmt.Println("error reading game name input")
		}
		name = strings.TrimSpace(name)
		hpc.router.Push(routes.GAMES_ROUTE, RouteArgs{"gameName": name})
	}})
	if session.CurrentUser != nil && session.CurrentUser.IsAdmin() {
		// admin actions
		result = append(result, Action{Name: "View Sales Report", ActionFunc: func() {
			hpc.router.Push(routes.SALES_REPORT_ROUTE, RouteArgs{})
		}})
		//
		result = append(result, Action{Name: "View Rentals Overdue", ActionFunc: func() {
			// TODO
			// hpc.router.Push(routes.RENTALS_OVERDUE_REPORT_ROUTE, RouteArgs{})
		}})
		//
		result = append(result, Action{Name: "View Reviews Report", ActionFunc: func() {
			// TODO
			// hpc.router.Push(routes.REVIEWS_REPORT_ROUTE, RouteArgs{})
		}})
	} else {
		// normal user actions
		result = append(result, Action{Name: "View Cart", ActionFunc: func() {
			hpc.router.Push(routes.CART_ROUTE, RouteArgs{})
		}})
	}
	// only append login/register if not logged in
	if session.CurrentUser == nil {
		result = append(result, Action{Name: "Login", ActionFunc: func() {
			hpc.router.Push(routes.LOGIN_ROUTE, RouteArgs{})
		}})
		result = append(result, Action{Name: "Register", ActionFunc: func() {
			hpc.router.Push(routes.REGISTER_ROUTE, RouteArgs{})
		}})
	} else {
		// allow logout
		result = append(result, Action{Name: "Logout", ActionFunc: func() {
			session.CurrentUser = nil
			hpc.router.Push(routes.HOME_PAGE_ROUTE, RouteArgs{})
		}})
	}
	result = append(result, Action{Name: "Exit", ActionFunc: func() {
		hpc.router.Pop()
	}})
	return result
}

func (hpc *homePageCli) HandleRoute(args RouteArgs, session *Session) {

	fmt.Println("Welcome to ROC Gameshop")
	fmt.Println("")
	if session.CurrentUser != nil {
		fmt.Printf("Welcome back, %s\n\n", session.CurrentUser.Name)
	}

	// get user actions
	actions := hpc.GetUserActions(session)
	PromptUserForActions(actions, hpc.reader)

}
