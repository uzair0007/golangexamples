// Package "golangexamples" contains GO Examples
//package golangexamples
package main

import (
    "fmt"
    "strings"
	"github.com/ehteshamz/greetings"
)

// GoLangString is a global variable
//var GoLangSlice []byte = := []byte{'H', 'E', 'L', 'L', 'O', ' ', 'F', 'R', 'I', 'E', 'N', 'D', 'S', '!'}

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


func main() {
	stc := []byte{'F', 'R', 'I', 'E', 'N', 'D', 'S', '!'}
	fmt.Println(ConcatSlice(stc))
	
	//var s = "wkh dgydgfh zloo ehjlq dw wkuhh sp"
	//Text: If he had anything confidential to say.
	//Cipher: jg if ibe bozuijoh dpogjefoujbm up tbz.
	msg := []byte{'i', 'f', '@', 'h', 'e', '*', 'h', 'a', 'd', '-', 'a', 'n', 'y', 't', 'h', 'i', 'n', 'g', ' ', 'c', 'o', 'n', 'f', 'i', 'd', 'e', 'n', 't', 'i', 'a', 'l', ' ', 't', 'o', ' ', 's', 'a', 'y', '.'}
	fmt.Println(Encrypt(msg, 3))
	
	fmt.Println(EZGreetings("EZ"))
	
}