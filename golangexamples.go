// Package "golangexamples" contains GO Examples
package golangexamples


//Importing Essential Libraries
import (
    "strings"
	"github.com/ehteshamz/greetings"
)


// GoLangString is a global variable
var GoLangString = "Hello World!"


// Converts the String to Slice([]byte)
func String2Slice(str string) []byte {
	
	// Length of the Slice
	var length = len(str)
	
	// Slice to Return
	slice := make([]byte, length)
	
	// Loop to Convert the String to Slice
	for i := 0; i < length; i++ {
		slice[i] = str[i]
	}
	
	// Return the String
	return slice
}


// Returns the contents of the slice concatenated together and separated by a dash (-).
func ConcatSlice(sliceToConcat []byte) string {
	
	// String to Return
	var s string
	
	// Length of the Slice
	var length = len(sliceToConcat)
	
	// Loop to Convert the Slice to String
	for i := 0; i < length; i++ {
		if i == (length - 1) {
			s += string(sliceToConcat[i])
		} else {
			s += string(sliceToConcat[i]) + "-"
		}
	}
	
	// Return the String
	return s
}


// Encrypts the slice using the Caesar cypher, with the number provided as the shift.
func Encrypt(sliceToEncrypt []byte, ceaserCount int) string {

	// String to Return
	var s string
	
	// Length of the Slice
	var length = len(sliceToEncrypt)
	
	// String of Alphabets
	var alphas = "abcdefghijklmnopqrstuvwxyz"
	var alphas_length = len(alphas)
	
	// Loop to Encrypt the Slice
	for i := 0; i < length; i++ {
	
		// The Current Letter to Encrypt
		var c = strings.ToLower(string(sliceToEncrypt[i]))
		
		// Checking if the Current Letter is an Alphabet
		var b = false
		for j := 0; j < alphas_length; j++ {
			if c == string(alphas[j]) {
				b = true
			}
		}
		
		// Encrypting the Current Letter (If Alphabet)
		if b {
		
			// Checking for the Letter in Alphabets
			var idx int
			for j := 0; j < alphas_length; j++ {
				if c == string(alphas[j]) {
					idx = j
				}
			}
			
			// Encrypting with "ceaserCount"
			s += string(alphas[(idx + ceaserCount) % alphas_length])
			
		} else {
			s += c
		}
		
	}
	
	// Return the String
	return s
}


// Invokes the "greetings" package and returns the resulting string.
func EZGreetings(name string) string {
	return greetings.PrintGreetings(name)
}