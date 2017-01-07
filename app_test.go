package main

import (
	"testing"
)

func isEqualToInt(v1, v2 int, t *testing.T) {
	if v1 != v2 {
		t.Error(v1, "not equal to", v2)
	}
}

func isEqualToString(v1, v2 string, t *testing.T) {
	if v1 != v2 {
		t.Error(v1, "not equal to", v2)
	}
}

func TestParseTime(t *testing.T) {
	app := App{}
	isEqualToInt(app.ParseTime("56"), 56, t)
	isEqualToInt(app.ParseTime("05:45"), 345, t)
	isEqualToInt(app.ParseTime("24:50:02"), 89402, t)
	isEqualToInt(app.ParseTime("1:24:50:02"), 0, t)
}

func TestFormatString(t *testing.T) {
	app := App{}
	isEqualToString(app.FormatString(56), "00:00:56", t)
	isEqualToString(app.FormatString(345), "00:05:45", t)
	isEqualToString(app.FormatString(89402), "24:50:02", t)

	app.Config.Second = true
	isEqualToString(app.FormatString(56), "56", t)
	isEqualToString(app.FormatString(345), "345", t)
	isEqualToString(app.FormatString(89402), "89402", t)
}
