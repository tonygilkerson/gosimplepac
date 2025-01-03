package main

import (
	"fmt"

	"github.com/tonygilkerson/gosimplepac/pkg"
	"honnef.co/go/tools/simple"
)

func main() {
	fmt.Printf("hi\n")

	// Use simple pacakge
	name := pkg.GetName()
	fmt.Printf("pkg.GetName said: %s\n", name)


}
