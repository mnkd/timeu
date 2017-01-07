package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type App struct {
	Args   []string
	Config Config
}

func (app App) Usage() {
	str := `Usage:
 timeu [-s|-v] [time]
 timeu [-s|-v] [time1] [+ | -] [time2]

Examples:
 $ timeu 56
 00:00:56

 $ timeu -s 42:56
 2576

 $ timeu 11:42:56 + 2:30:59
 14:13:55

 $ timeu -s 11:42:56 + 2:30:59
 51235

 $ timeu 11:42:56 - 2:30:59
 09:11:57
`
	fmt.Fprintln(os.Stderr, str)
}

func (app App) FormatString(t int) string {
	if app.Config.Second {
		return fmt.Sprintf("%d", t)
	}

	var h, m, s int
	h = t / 3600
	m = t / 60 % 60
	s = t % 60

	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}

func (app App) ParseTime(s string) int {
	array := strings.Split(s, ":")

	switch len(array) {
	case 1:
		i, _ := strconv.Atoi(array[0])
		return i
	case 2:
		m, _ := strconv.Atoi(array[0])
		s, _ := strconv.Atoi(array[1])
		return (m * 60) + s
	case 3:
		h, _ := strconv.Atoi(array[0])
		m, _ := strconv.Atoi(array[1])
		s, _ := strconv.Atoi(array[2])
		return (h * 60 * 60) + (m * 60) + s
	}

	fmt.Fprintf(os.Stderr, "Invalid time format: %v | array: %v\n", s, array)
	return 0
}

func (app App) Run() int {
	args := app.Args
	switch len(args) {
	case 0:
		app.Usage()
		return ExitCodeError
	case 1:
		v := app.ParseTime(args[0])
		fmt.Fprintln(os.Stdout, app.FormatString(v))
		return ExitCodeOK
	case 3:
		v1 := app.ParseTime(args[0])
		op := args[1]
		v2 := app.ParseTime(args[2])

		switch op {
		case "+":
			fmt.Fprintln(os.Stdout, app.FormatString(v1+v2))
		case "-":
			fmt.Fprintln(os.Stdout, app.FormatString(v1-v2))
		default:
			fmt.Fprintln(os.Stderr, "Unknown operator:", op)
			fmt.Fprintln(os.Stderr, "timeu supports + and -.")
			return ExitCodeError
		}
		return ExitCodeOK
	}

	app.Usage()
	return ExitCodeError
}

func NewApp(args []string, config Config) App {
	return App{Args: args, Config: config}
}
