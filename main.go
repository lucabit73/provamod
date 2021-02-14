package main

import (
	"fmt"

	"github.com/lucabit73/provamod/pkg/shared"
)

func main() {

	fmt.Println("helle")
	a := shared.Bau{"luca", 42}
	a.Miao()
}
