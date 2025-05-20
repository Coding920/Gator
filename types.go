package main

import (
	"github.com/Coding920/Gator/internal/config"
)

type state struct {
	cfg *config.Config
}

type command struct {
	name string
	args []string
}
