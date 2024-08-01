package cli_test

import "roc-gameshop-app/cli"

type MockRouter struct {
	BackStack []cli.StackItem
	PopCalled int
}

// AddRouteCli implements cli.Router.
func (m *MockRouter) AddRouteCli(route string, cli cli.Cli) {
	panic("unimplemented")
}

// Pop implements cli.Router.
func (m *MockRouter) Pop() cli.StackItem {
	m.PopCalled += 1
	return cli.NewStackItem("", cli.RouteArgs{})
}

// Push implements cli.Router.
func (m *MockRouter) Push(route string, args cli.RouteArgs) {
	m.BackStack = append(m.BackStack, cli.NewStackItem(route, args))
}

// Run implements cli.Router.
func (m *MockRouter) Run(session *cli.Session) {
	panic("unimplemented")
}

func NewMockRouter() *MockRouter {
	return &MockRouter{}
}
