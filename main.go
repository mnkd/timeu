package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	ExitCodeOK int = iota
	ExitCodeError
)

var (
	Version  string
	Revision string
)

var app App

func init() {
	var version, second bool
	flag.BoolVar(&version, "v", false, "print version.")
	flag.BoolVar(&second, "s", false, "print second.")
	flag.Parse()

	if version {
		fmt.Fprintln(os.Stdout, "Version:", Version)
		fmt.Fprintln(os.Stdout, "Revision:", Revision)
		os.Exit(ExitCodeOK)
	}

	config := NewConfig(second)
	app = NewApp(flag.Args(), config)
}

func main() {
	os.Exit(app.Run())
}
