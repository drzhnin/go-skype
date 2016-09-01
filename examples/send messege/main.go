package main

import (
	"fmt"

	"github.com/andrewdruzhinin/go-skype/skype"
)

func main() {
	client := skype.NewClient("client_id", "client_secret")
	auth, _, err := client.Authorization.Authorize()
	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		return
	}
	client.Token = fmt.Sprintf("%s %s", auth.TokenType, auth.AccessToken)
	resp, err := client.Messege.Send("chat_id", "message/text", "Messege")
	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		return
	}
	fmt.Println(resp)
}
