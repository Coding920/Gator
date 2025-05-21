package main

import (
	"github.com/Coding920/Gator/internal/config"
	"github.com/Coding920/Gator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

type command struct {
	name string
	args []string
}
