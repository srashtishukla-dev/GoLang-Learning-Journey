package auth

import "fmt"

func LoginWithCredentials(username string, password string) {
	fmt.Println("Logging user using:", username, password)
}

func GetSession() string {
	return "session123"
}
