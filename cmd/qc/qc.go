package main

import (
	"fmt"
	"os"

	"github.com/yukidarake/qc"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}

	s := os.Args[1]

	fmt.Print(qc.Format(s))
}
