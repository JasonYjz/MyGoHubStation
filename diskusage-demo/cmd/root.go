package cmd

import (
	"flag"
)

type CmdEnv struct {
	Path    string
}

var C = &CmdEnv{}

func init() {
	flag.StringVar(&C.Path, "path", "/", "disk path")

	flag.Parse()
}
