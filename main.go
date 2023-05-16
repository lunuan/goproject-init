package main

import (
	"flag"
	"os"
	"strings"

	"github.com/lunuan/goprojectinit/cmd/goprojectinit"
)

func main() {
	wd, _ := os.Getwd()
	splited := strings.Split(wd, "/")
	current_dir := splited[len(splited)-1]

	var project_name string
	flag.StringVar(&project_name, "name", current_dir, "project name")
	flag.Parse()
	goprojectinit.NewProject(project_name)
}
