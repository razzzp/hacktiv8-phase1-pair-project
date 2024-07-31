package cli

import (
	"bufio"
	"fmt"
	"roc-gameshop-app/entities"
	"roc-gameshop-app/handlers"
	"roc-gameshop-app/routes"
	"strings"
)

type registerCli struct {
	userHandler handlers.UserHandler
	router      Router
	reader      *bufio.Reader
}

func NewUserCli(router Router, reader *bufio.Reader, userHandler handlers.UserHandler) Cli {
	return &registerCli{
		router:      router,
		reader:      reader,
		userHandler: userHandler,
	}
}

func (uC *registerCli) HandleRoute(args RouteArgs, session *Session) {
	// logic of game details page goes here
	// TODO
	fmt.Println("Register User")
	fmt.Printf("Please enter your name: ")
	for {

		//get user name
		name, err := uC.reader.ReadString('\n')
		if err != nil {
			fmt.Println("error reading user name input")
			continue
		}
		name = strings.TrimSpace(name)
		//get user email
		email, err := uC.reader.ReadString('\n')
		if err != nil {
			fmt.Println("error reading user email input")
			continue
		}
		email = strings.TrimSpace(email)
		//get user role
		role, err := uC.reader.ReadString('\n')
		if err != nil {
			fmt.Println("error reading user role input")
			continue
		}
		role = strings.TrimSpace(role)
		//get user phoneNum
		phoneNumber, err := uC.reader.ReadString('\n')
		if err != nil {
			fmt.Println("error reading user phoneNumber input")
			continue
		}
		phoneNumber = strings.TrimSpace(phoneNumber)
		//get user salt
		salt, err := uC.reader.ReadString('\n')
		if err != nil {
			fmt.Println("error reading user salt input")
			continue
		}
		salt = strings.TrimSpace(salt)
		//get user password
		password, err := uC.reader.ReadString('\n')
		if err != nil {
			fmt.Println("error reading user password input")
			continue
		}
		password = strings.TrimSpace(password)

		instance := entities.User{
			Name:         name,
			Email:        email,
			Role:         role,
			PhoneNumber:  phoneNumber,
			Salt:         salt,
			PasswordHash: password,
		}
		err = uC.userHandler.Create(instance)
		if err != nil {
			continue
		}
		uC.router.Push(routes.HOME_PAGE_ROUTE, RouteArgs{})
		return
	}
	// for {
	// 	fmt.Print("What would you like to do? ")
	// 	input, _ := gDC.reader.ReadString('\n')
	// 	input = strings.TrimSpace(input)
	// 	if input == "1" {
	// 		// push to router and return to go to another route
	// 		gDC.router.Push(routes.GAME_DETAILS_ROUTE, RouteArgs{"gameId": "2"})
	// 		return
	// 	} else if input == "2" {
	// 		// pop and return to return to previous route
	// 		gDC.router.Pop()
	// 		return
	// 	} else {
	// 		fmt.Println("Invalid Action.")
	// 	}
	// }

}
