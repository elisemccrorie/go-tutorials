package main

import (
	"fmt"
	"os"
	"crypto/sha1"
	"encoding/hex"
)

func main() {

	// The first argument
	// is always program name
	//myProgramName := os.Args[0]

	myString := os.Args[1]

	// it will display
	// the program name
//	fmt.Println(myProgramName)
	fmt.Println(myString)
	fmt.Println(MessageKey(myString))
}

// MessageKey returns the hash of a translatable text string, which is used in
// the Crowdin-Go Database as the message key, `msgkey`.
func MessageKey(message string) string {
	shaBytes := sha1.Sum([]byte(message))
	return hex.EncodeToString(shaBytes[:])
}
