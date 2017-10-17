package main

import (
	"fmt"
	"io"
	"os"
)

//usage is a helpful usage string shown when
//the user doesn't supply enough arguments.
//The `...` allows line breaks within the string
const usage = `
USAGE:
	hmac sign signing-key < file-to-sign
	hmac verify signing-key signature < file-to-verify
`

//exit codes
//iota resets to 0 every time you start
//a const block, and increments by 1
//each time you use it
const (
	_                        = iota //ignore first usage (which is zero)
	exitCodeUsage                   // = 1
	exitCodeProcessing              // = 2
	exitCodeInvalidSignature        // = 3
)

//showUsage shows the usage string and exits
//with the code exitCodeUsage
func showUsage() {
	fmt.Println(usage)
	os.Exit(exitCodeUsage)
}

//sign returns a base64-encoded HMAC signature given a
//signingKey and a read stream. It returns an error if
//there was an error reading from the stream.
func sign(signingKey string, stream io.Reader) (string, error) {
	//TODO: implement this function according to the comments
	//HINTS:
	//https://drstearns.github.io/tutorials/sessions/#secdigitalsignatureswithhmac
	//https://golang.org/pkg/crypto/hmac/
	//https://golang.org/pkg/io/#Copy
	//https://golang.org/pkg/encoding/base64/
	return "", fmt.Errorf("TODO")
}

//verify returns true if the base64-encoded HMAC `signature`
//matches the bytes read from `stream`, or false if otherwise.
//If there is an error decoding the base64 signature, or reading
//from `stream`, this will return false and the error.
func verify(signingKey string, signature string, stream io.Reader) (bool, error) {
	//TODO: implement this function according to the comments
	//HINTS: (same as above plus the following)
	//https://golang.org/pkg/crypto/subtle/#ConstantTimeCompare
	return false, fmt.Errorf("TODO")
}

func main() {
	if len(os.Args) < 3 {
		showUsage()
	}

	//TODO: use the os.Args slice
	//to get the command and signing key
	//then switch on the command
	//and call the appropriate function above,
	//printing the return value or error

}
