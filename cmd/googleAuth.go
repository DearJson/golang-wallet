package main

import (
	"fmt"
	"gfast/library"
)

func main() {
	fmt.Printf("%v", library.NewGoogleAuth().GetSecret())
}
