package main

import (
	"fmt"

	"github.com/andrewdruzhinin/go-skype/skype"
)

func main() {
	client := skype.NewClient("your_client_id", "your_client_secret")
	auth, _, err := client.Authorization.Authorize()
	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		return
	}
	client.Token = fmt.Sprintf("%s %s", auth.TokenType, auth.AccessToken)
	fmt.Printf("Bot token: %s \n", client.Token)
}
