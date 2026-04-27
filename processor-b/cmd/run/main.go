package main

import (
	"context"

	orca "github.com/orca-telemetry/go"
)

var processor *orca.Processor

func init() {
	processor = orca.NewProcessor("buzz")
}

func main() {
	ctx := context.Background()
	processor.Start(ctx)
}
