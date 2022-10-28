package console

import (
	"fmt"
	"log"

	"github.com/fatih/color"
)

type ConsoleType struct{}

var Console = new(ConsoleType)

func (*ConsoleType) Info(msg string) {
	color.Green(fmt.Sprintf("[Lxb Info] %s", msg))
	log.Printf("[Lxb Info] %s", msg)
}

func (*ConsoleType) Warn(msg string) {
	color.Yellow(fmt.Sprintf("[Lxb Warn] %s", msg))
	log.Printf("[Lxb Warn] %s", msg)
}

func (*ConsoleType) Error(msg string, err string) {
	color.Red(fmt.Sprintf("[Lxb Error] %s", msg))
	log.Printf("[Lxb Error] %s", msg)
	if err != "" {
		panic(err)
	}
}
