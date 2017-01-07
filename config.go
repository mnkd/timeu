package main

import "fmt"

type Config struct {
	Second bool
}

func NewConfig(second bool) Config {
	return Config{second}
}

func (c Config) String() string {
	return fmt.Sprintf("%T: {Second: %v}", c, c.Second)
}
