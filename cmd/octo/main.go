package main

import (
	"github.com/gdegiorgio/octo/internal/commands/octo"
)

// set via ldflags
var version = "development"

func main() {
	octo.Main(version)
}
