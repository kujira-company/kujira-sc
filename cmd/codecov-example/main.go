package main

import (
	"github.com/kujira-company/kujira-sc/internal/cmd1"
	"github.com/kujira-company/kujira-sc/internal/cmd2"
)

func main() {
	_ = cmd1.Command1(1)
	_ = cmd1.Command1(2)
	_ = cmd2.Command2(1)
	_ = cmd2.Command2(2)
}
