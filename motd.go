package motd

import (
	_ "embed"
	"math/rand"
	"strings"
)

//go:embed motd.txt
var s string

func Next() string {
	ss := strings.Split(s, "\n")
	for {
		r := rand.Intn(len(ss))
		t := ss[r]
		if strings.TrimSpace(t) != "" {
			return t
		}
	}
}
