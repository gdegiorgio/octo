package main

import (
	"github.com/gdegiorgio/octo/internal/commands/root"
)


func main(){
	root.NewRootCmd().Execute()
}
