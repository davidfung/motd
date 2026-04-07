package main

// $ go build -o motd

import (
	"fmt"

	"github.com/davidfung/motd"
)

func main() {
	fmt.Println(motd.Next())
}
