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

func (hpc *homePageCli) GetUserActions(session *Session) []Action {
	result := []Action{}
	result = append(result, Action{Name: "Search Games", ActionFunc: func() {
		hpc.router.Push(routes.GAMES_ROUTE, RouteArgs{})
	}})
	result = append(result, Action{Name: "View Cart", ActionFunc: func() {
		hpc.router.Push(routes.GAMES_ROUTE, RouteArgs{})
	}})
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

	// get user actions
	actions := hpc.GetUserActions(session)
	PromptUserForActions(actions, hpc.reader)

}
