package main

import (
	"fmt"
	"log"

	"github.com/ewielguszewski/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal("Error reading config:", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)

	err = cfg.SetUser("ernest")

	cfg, err = config.Read()
	if err != nil {
		log.Fatal("Error reading config:", err)
	}
	fmt.Printf("Read config again: %+v\n", cfg)

}
