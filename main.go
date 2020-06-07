package main

import (
	"fmt"

	"github.com/mdcg/hashkiller/cmd"
)

func main() {
	s := "mauro"
	cmd.Greetings()
	cmd.AvailableCommands()
	fmt.Println(cmd.EncryptToMd5(&s))
	fmt.Println(cmd.EncryptToSha1(&s))
	fmt.Println(cmd.EncryptToSha256(&s))
}
