package main

import (
	orca "github.com/orca-telemetry/go"
)

var processor *orca.Processor

func init() {
	processor = orca.NewProcessor("buzz")
	// TODO: Create all the algorithms
}

func main() {
	// TODO: regsiter all the algorithms
}
