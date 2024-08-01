package cli

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"roc-gameshop-app/entities"
	"runtime"
	"strconv"
	"strings"
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
	Pop() StackItem
	// starts the 'program'
	//
	//	will keep runnning until there is no more items in backstack
	Run(session *Session)
}

type StackItem struct {
	route string
	args  RouteArgs
}

type routerV1 struct {
	routeClis map[string]Cli
	backStack []StackItem
}

func NewRouter() Router {
	return &routerV1{
		routeClis: map[string]Cli{},
		backStack: []StackItem{},
	}
}

func newStackItem(route string, args RouteArgs) StackItem {
	return StackItem{
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
	fmt.Println("push backstack: ", r.backStack)
}

func (r *routerV1) Pop() StackItem {
	fmt.Println("backStack", r.backStack)
	// remove last element from stack
	topStackItem := r.backStack[len(r.backStack)-1]
	if len(r.backStack) > 0 {
		r.backStack = r.backStack[:len(r.backStack)-1]
	}
	return topStackItem
}

func (r *routerV1) Peek() *StackItem {
	if len(r.backStack) > 0 {
		return &r.backStack[len(r.backStack)-1]
	}
	return nil
}

func (r *routerV1) Run(session *Session) {
	// as long as there is something in back stack keep routing
	for len(r.backStack) > 0 {
		// get top of stack
		topStackItem := r.Peek()
		// fmt.Println(topStackItem)
		// get cli handler to handle route
		cli := r.routeClis[topStackItem.route]
		if cli == nil {
			// no cli not assigned to route
			// panic?
			panic(fmt.Sprintf("route '%s' has no cli assigned", topStackItem.route))
		}
		// clear the screen
		CallClear()

		// run cli
		cli.HandleRoute(topStackItem.args, session)
	}
}

// an action with an attached function to run
// if that action is selected by user
type Action struct {
	Name       string
	ActionFunc func()
}

// helper to prompt user action
func PromptUserForActionInput(reader *bufio.Reader) (int, error) {
	fmt.Print("What would you like to do? ")
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	input = strings.TrimSpace(input)

	// convert to int
	inputAsInt, err := strconv.Atoi(input)
	if err != nil {
		return 0, errors.New("please enter a valid number")
	}

	return inputAsInt, nil
}

// prints the given list of actions and asks
// user to choose one
// if valid choice
func PromptUserForActions(actions []Action, reader *bufio.Reader) {
	// prints actions
	for i, action := range actions {
		fmt.Printf("%d. %s\n", i+1, action.Name)
	}
	fmt.Println("")
	for {
		input, err := PromptUserForActionInput(reader)
		if err != nil {
			fmt.Printf("Invalid input: %s, please try again.\n", err)
			continue
		}
		if input > len(actions) {
			fmt.Printf("Invalid input, please try again.\n")
			continue
		}

		// get action
		selectedAction := actions[input-1]
		selectedAction.ActionFunc()
		return
	}
}

var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["darwin"] = func() {
		cmd := exec.Command("clear") //Mac example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}
