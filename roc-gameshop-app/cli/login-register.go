package cli

import (
	"bufio"
	"fmt"
	"roc-gameshop-app/routes"
)

type loginRegisterCli struct {
	router Router
	reader *bufio.Reader
}

func (l *loginRegisterCli) GetActions(session *Session) []Action {
	result := []Action{}

	result = append(result, Action{
		Name: "Login",
		ActionFunc: func() {
			l.router.Push(routes.LOGIN_ROUTE, RouteArgs{})
		},
	})

	result = append(result, Action{
		Name: "Register",
		ActionFunc: func() {
			l.router.Push(routes.REGISTER_ROUTE, RouteArgs{})
		},
	})

	return result
}

// HandleRoute implements Cli.
func (l *loginRegisterCli) HandleRoute(args RouteArgs, session *Session) {
	if session.CurrentUser != nil {
		// already logged in, pop
		l.router.Pop()
		return
	}

	fmt.Println("Login/Register")
	fmt.Println("")

	actions := l.GetActions(session)
	PromptUserForActions(actions, l.reader)
}

func NewLoginRegisterCli(
	router Router,
	reader *bufio.Reader,
) Cli {
	return &loginRegisterCli{
		router: router,
		reader: reader,
	}
}
