package cmd

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/briandowns/spinner"
)

var s = spinner.New(spinner.CharSets[14], 100*time.Millisecond)

func StartDecryption(wordlist_path, informed_hash string, encryption_func func(*string) string) string {
	start := time.Now()
	f, err := os.Open(wordlist_path)
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	s.Start()
	for scanner.Scan() {
		current_word := scanner.Text()
		encryptedWord := encryption_func(&current_word)
		if encryptedWord == informed_hash {
			s.Stop()
			return fmt.Sprintf("Decrypted hash! Value: %s - Took: %x milliseconds\n", current_word, time.Since(start).Milliseconds())
		}
	}
	s.Stop()
	return fmt.Sprintf("Not found. - Took: %x milliseconds\n", time.Since(start).Milliseconds())
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
