// Package "main"
package main


//Importing Essential Libraries
import (
    "fmt"
	"github.com/uzair0007/golangexamples"
)


// Main Function to Check the "golangexamples" Package
func Main() {
	s := golangexamples.String2Slice(golangexamples.GoLangString)
	fmt.Println(s)
}