package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var (
	interval          = flag.String("interval", "1s", "The interval duration in standard golang time.Duration format")
	number            = flag.Int("number", 5, "The number of times to iterate (must be at least 1)")
	alternateStdError = flag.Bool("alternate-stderror", false, "Alternate between stdout and stderr")
	exitCode          = flag.Int("exitcode", 0, "The exit code to exit with")
)

func main() {
	flag.Parse()
	if len(*interval) == 0 ||
		*number <= 0 {
		flag.Usage()
		os.Exit(2)
	}

	intervalDuration, err := time.ParseDuration(*interval)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < *number; i++ {
		var writer io.Writer = os.Stdout
		if *alternateStdError && i%2 == 1 {
			writer = os.Stderr
		}

		writer.Write([]byte(fmt.Sprintf("Dummy %d/%d (every %s)\n", i+1, *number, intervalDuration)))
		time.Sleep(intervalDuration)
	}

	os.Exit(*exitCode)
}
