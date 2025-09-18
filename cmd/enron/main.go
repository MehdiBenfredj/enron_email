package main

import (
	"fmt"
	"os"
	"time"

	"github.com/mehdibenfredj/enron_go/internal/processing"
)

func main() {
	// parse args here
	maildir := os.Args[1]
	resultPath := os.Args[2]

	// For testing purpose only

	start := time.Now()
	processing.RunParallel(maildir, resultPath)
	elapsed := time.Since(start)
	fmt.Printf("Execution took %s\n", elapsed)

}
