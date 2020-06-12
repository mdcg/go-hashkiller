package cmd

import (
	"flag"
	"fmt"
)

func Greetings() {
	fmt.Println("Welcome to the mdcg_'s Hashkiller!")
	fmt.Print("To find out how to use it, type \"-h\" or \"--help\"\n\n")
}

func AvailableCommands() (*string, *string, *string) {
	wordlist_path := flag.String("wordlist", "wordlist.txt", "Path to the wordlist that will be used for analysis")
	encryption_type := flag.String("type", "md5", "Type of encryption to be performed.	")
	informed_hash := flag.String("hash", "", "Hash to be broken")
	flag.Parse()

	return wordlist_path, encryption_type, informed_hash
}
