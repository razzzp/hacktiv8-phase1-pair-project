package cli

import (
	"bufio"
	"fmt"
	"roc-gameshop-app/entities"
	"roc-gameshop-app/handlers"
	"roc-gameshop-app/routes"
	"strings"
	"time"
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

	fmt.Println("Register User")

	for {
		//get user name
		fmt.Printf("Please enter your name: ")
		name, err := uC.reader.ReadString('\n')
		if err != nil {
			fmt.Println("error reading user name input")
			continue
		}
		name = strings.TrimSpace(name)

		//get user email
		fmt.Printf("Please enter your email: ")
		email, err := uC.reader.ReadString('\n')
		if err != nil {
			fmt.Println("error reading user email input")
			continue
		}
		email = strings.TrimSpace(email)
		//get user role
		fmt.Printf("Please enter your role: ")
		role, err := uC.reader.ReadString('\n')
		if err != nil {
			fmt.Println("error reading user role input")
			continue
		}
		role = strings.TrimSpace(role)
		//get user phoneNum
		fmt.Printf("Please enter your phone number: ")
		phoneNumber, err := uC.reader.ReadString('\n')
		if err != nil {
			fmt.Println("error reading user phoneNumber input")
			continue
		}
		phoneNumber = strings.TrimSpace(phoneNumber)

		//get user password
		fmt.Printf("Please enter your password: ")
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
			PasswordHash: password,
		}
		err = uC.userHandler.Create(instance)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("User '%s' successfully registered\n", email)
		// to let user see success msg :P
		time.Sleep(time.Second)
		uC.router.Push(routes.HOME_PAGE_ROUTE, RouteArgs{})
		return
	}
}
