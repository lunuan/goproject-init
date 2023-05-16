package main

import (
	"flag"

	"github.com/lunuan/goprojectinit/cmd/goprojectinit"
)

func main() {
	var project_name string
	flag.StringVar(&project_name, "name", "test", "project name")
	goprojectinit.NewProject(project_name)
}
