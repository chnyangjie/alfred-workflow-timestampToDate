package main

// Package is called aw
import (
	"fmt"
	"os"
	"strconv"
	"time"

	aw "github.com/deanishe/awgo"
)

// Workflow is the main API
var wf *aw.Workflow

func init() {
	// Create a new Workflow using default settings.
	// Critical settings are provided by Alfred via environment variables,
	// so this *will* die in flames if not run in an Alfred-like environment.
	wf = aw.New()
}

// Your workflow starts here
func run() {
	// Add a "Script Filter" result
	arg := os.Args[1]
	argInt, err := strconv.ParseInt(arg, 10, 64)
	if err != nil {
		wf.NewItem(fmt.Sprintf("Invalid input %s", arg))
	} else {
		if argInt > 9999999999 {
			argInt = argInt / 1000
		}
		dt := time.Unix(argInt, 0)
		s := dt.Format("2006-01-02 15:04:05")
		wf.NewItem(fmt.Sprintf(s)).Valid(true).Arg(s)
		s = dt.Format("2006-01-02T15:04:05")
		wf.NewItem(fmt.Sprintf(s)).Valid(true).Arg(s)
	}
	// Send results to Alfred
	wf.SendFeedback()
}

func main() {
	// Wrap your entry point with Run() to catch and log panics and
	// show an error in Alfred instead of silently dying
	wf.Run(run)
}
