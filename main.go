package main

import (
	"fmt"
	"github.com/Coding920/Gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		panic(err)
	}

	err = cfg.SetUser("Justin")
	if err != nil {
		panic(err)
	}

	cfg, err = config.Read()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", cfg)
}
