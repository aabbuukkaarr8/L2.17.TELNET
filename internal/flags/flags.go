package flags

import (
	"flag"
	"time"
)

var (
	Host    string
	Port    string
	Timeout time.Duration
)

func ParseFlags() {
	flag.DurationVar(&Timeout, "timeout", 10*time.Second, "connection timeout")

	flag.Parse()

	args := flag.Args()
	if len(args) >= 2 {
		Host = args[0]
		Port = args[1]
	}
}
