package main

import (
	"fmt"

	"github.com/andrewdruzhinin/go-skype/skype"
)

func main() {
	client := skype.NewClient("client_id", "client_secret")
	_, err := client.Authorization.Authorize()
	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		return
	}
	_, err = client.Messege.Send("сhat_id", "message/text", "HI")
	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		return
	}
}
