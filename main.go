package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ewielguszewski/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal("Error reading config:", err)
	}

	programState := state{
		cfg: &cfg,
	}

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)

	args := os.Args[:]

	if len(args) < 2 {
		log.Fatal("usage: cli <command> [args]")
		return
	}

	cmdName := args[1]
	cmdArgs := args[2:]

	err = cmds.run(&programState, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal("Error:", err)
		return
	}
	fmt.Println("Command executed successfully")
	fmt.Printf("Current config: %+v\n", programState.cfg)

}
