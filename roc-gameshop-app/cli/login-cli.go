package cli

import (
	"bufio"
	"fmt"
	"roc-gameshop-app/handlers"
	"roc-gameshop-app/routes"
	"strings"
	"time"
)

type loginCli struct {
	authHandler handlers.AuthHandler
	router Router
	reader *bufio.Reader
}

func NewLoginCli(router Router, reader *bufio.Reader, authHandler handlers.AuthHandler) Cli {
	return &loginCli{
		router: router,
		reader: reader,
		authHandler: authHandler,
	}
}

func (aC *loginCli) HandleRoute(args RouteArgs, session *Session) {

	fmt.Println("Login User")
	
	loginLopp: for {
		//get user email
		fmt.Printf("Please enter your Email: ")
		email, err := aC.reader.ReadString('\n')
		if err != nil {
			fmt.Println("error reading user email input")
			continue
		}
		email = strings.TrimSpace(email)

		//get user password
		fmt.Printf("Please enter your Password: ")
		password, err := aC.reader.ReadString('\n')
		if err != nil {
			fmt.Println("error reading user password input")
			continue
		}
		password = strings.TrimSpace(password)

		user, err := aC.authHandler.Login(email, password)
		if err != nil {
			fmt.Println("Error login: ", err)
			//get user email
			fmt.Printf("Would you like to retry login? (y/n) ")
			res, err := aC.reader.ReadString('\n')
			if err != nil {
				fmt.Println("error reading user retry input")
				continue
			}
			res = strings.TrimSpace(res)
			switch res {
			case "y":
				continue
			case "n":
				break loginLopp
			default:
				continue
			}
		} else {
			session.CurrentUser = user
			fmt.Println("Login Successful")
			break loginLopp
		}
	}
	time.Sleep(time.Second)
	aC.router.Push(routes.HOME_PAGE_ROUTE, RouteArgs{})

}