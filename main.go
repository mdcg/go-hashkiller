package main

import (
	"fmt"

	"github.com/mdcg/hashkiller/cmd"
)

func main() {
	cmd.Greetings()
	wordlist_path, encryption_type, informed_hash := cmd.AvailableCommands()
	encryption_func := cmd.ChooseEncryptionType(*encryption_type)
	result := cmd.StartDecryption(*wordlist_path, *informed_hash, encryption_func)
	fmt.Printf(result)
	// 2ee2ef1dbc555ce6a84533475715806c
}
