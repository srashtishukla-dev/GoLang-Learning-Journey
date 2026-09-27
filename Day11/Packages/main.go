package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/srashtishukla-dev/podcast/Packages/auth"
	"github.com/srashtishukla-dev/podcast/Packages/user"
)

func main() {
	auth.LoginWithCredentials("srashtishukla", "secret")

	session := auth.GetSession()

	fmt.Println("session", session)

	u := user.User{
		Email: "user@email.com",
		Name:  "John Doe",
	}

	color.Red(u.Email)
	color.Green(u.Name)

}
