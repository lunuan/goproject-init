package goprojectinit

import (
	"fmt"
	"os"
	"os/exec"
)

func ExecCommand(cmd string) error {
	c := exec.Command("bash", "-c", cmd)
	output, err := c.CombinedOutput()
	fmt.Println(string(output))
	return err
}

func NewProject(name string) {
	err := ExecCommand("go mod init " + name)
	if err != nil {
		fmt.Println("Error init go project")
	}
	dirs := []string{
		"api",
		"cmd",
		"configs",
		"deployments",
		"init",
		"internal",
		"pkg",
		"scripts",
		"test",
		"tools",
		"vendor",
		"web",
		"build",
	}
	for _, dir := range dirs {
		os.Mkdir(dir, 0755)
		fmt.Println("create new dir " + dir)
	}

	files := []string{
		"readme.md",
		"Makefile",
	}
	for _, filename := range files {
		file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			fmt.Printf("Error creating %s file\n", filename)
		}
		defer file.Close()
		fmt.Println("create new file " + filename)
	}

}
