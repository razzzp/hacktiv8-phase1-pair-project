package cli

import (
	"bufio"
	"fmt"
	"roc-gameshop-app/routes"
)

type cartCli struct {
	router Router
	reader *bufio.Reader
}

func NewCartCli(router Router, reader *bufio.Reader) Cli {
	return &cartCli{
		router: router,
		reader: reader,
	}
}

func (cc *cartCli) checkout(session *Session) {
	// TODO
	if session.CurrentUser == nil {
		// have to login first
		cc.router.Push(routes.LOGIN_ROUTE, RouteArgs{})
		return
	}

}

func (cc *cartCli) removeItem(itemIndex int, session *Session) {
	// TODO
	if session.CurrentUser == nil {
		// have to login first
		cc.router.Push(routes.LOGIN_ROUTE, RouteArgs{})
		return
	}

}

func (cc *cartCli) GetUserActions(session *Session) []Action {
	result := []Action{}
	result = append(result, Action{Name: "Check Out", ActionFunc: func() {
		cc.checkout(session)
	}})
	result = append(result, Action{Name: "Remove Item", ActionFunc: func() {
		cc.router.Push(routes.GAMES_ROUTE, RouteArgs{})
	}})
	result = append(result, Action{Name: "Exit", ActionFunc: func() {
		cc.router.Pop()
	}})
	return result
}

func (cc *cartCli) HandleRoute(args RouteArgs, session *Session) {

	fmt.Println("Welcome to ROC Gameshop")
	fmt.Println("")

	// get user actions
	actions := cc.GetUserActions(session)
	PromptUserForActions(actions, cc.reader)

}
