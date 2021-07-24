package main

import (
	"fmt"

	"github.com/tonygilkerson/gosimplepac/pkg"
)

func main() {
	fmt.Printf("hi\n")

	// Use simple pacakge
	name := pkg.GetName()
	fmt.Printf("pkg.GetName said: %s\n", name)

}
