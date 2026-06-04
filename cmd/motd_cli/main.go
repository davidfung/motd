// To build the command line tool manually:
// $ go build -o motd
//
// To install it as a Go tool:
// $ go install github.com/davidfung/motd/cmd/motd_cli@latest

package main

import (
	"fmt"

	"github.com/davidfung/motd"
)

func main() {
	fmt.Println(motd.Next())
}
