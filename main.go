package main

import (
	"fmt"

	"github.com/emincanozcan/insider-assessment/config"
)

func main() {
	c, err := config.Load()
	if err != nil {
		panic("Missing environment variables!" + err.Error())
	}
	fmt.Println("ready!", c)
}
