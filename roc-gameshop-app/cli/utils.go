package cli

import (
	"fmt"
	"roc-gameshop-app/entities"
)

type RouteArgs map[string]string

// contains list of items user wants to checkout
type Cart struct {
	// TODO
}

// for storing session data
type Session struct {
	// returns nil if not logged in
	CurrentUser *entities.User

	// returns current cart
	CurrentCart *Cart
	// AddToCart()
}

type Cli interface {
	HandleRoute(args RouteArgs, session *Session)
}

// helper to ease navigating between pages
type Router interface {
	// assigns a cli to a route
	AddRouteCli(route string, cli Cli)
	// pushes a new route to the stack along with args
	//
	//	args can be filled with anything and will be
	//	used by the corresponding clis to handle the requests
	Push(route string, args RouteArgs)
	// Removes last element from stack
	Pop()
	// starts the 'program'
	//
	//	will keep runnning until there is no more items in backstack
	Run(session *Session)
}

type stackItem struct {
	route string
	args  RouteArgs
}

type routerV1 struct {
	routeClis map[string]Cli
	backStack []stackItem
}

func NewRouter() Router {
	return &routerV1{
		routeClis: map[string]Cli{},
		backStack: []stackItem{},
	}
}

func newStackItem(route string, args RouteArgs) stackItem {
	return stackItem{
		route: route,
		args:  args,
	}
}

func (r *routerV1) AddRouteCli(route string, cli Cli) {
	r.routeClis[route] = cli
}

func (r *routerV1) Push(route string, args RouteArgs) {
	// push to stack
	r.backStack = append(r.backStack, newStackItem(route, args))
}

func (r *routerV1) Pop() {
	// remove last element from stack
	if len(r.backStack) > 0 {
		r.backStack = r.backStack[:len(r.backStack)-1]
	}
}

func (r *routerV1) Run(session *Session) {
	// as long as there is something in back stack keep routing
	for len(r.backStack) > 0 {
		// get top of stack
		topStackItem := r.backStack[len(r.backStack)-1]
		fmt.Println(topStackItem)
		// get cli handler to handle route
		cli := r.routeClis[topStackItem.route]
		if cli == nil {
			// no cli not assigned to route
			// panic?
			panic(fmt.Sprintf("route '%s' has no cli assigned", topStackItem.route))
		}
		// TODO
		// clear the screen

		// run cli
		cli.HandleRoute(topStackItem.args, session)
	}
}
