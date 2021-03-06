// govmd creates a markdown document with the current date in the filename and
// pre-filled YAML title, author, and date template fields and opens the file
// in Vim.
package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	// fileName stores the generated file name
	var fileName string

	// yamlHeader stores the generated YAML header
	var yamlHeader string

	// Get the current time
	now := time.Now()

	// Build file name
	switch len(os.Args) {
	case 1:
		fileName += now.Format("2006-01-02-150405") + ".md"
	case 2:
		fileName += now.Format("2006-01-02") + "-" + os.Args[1] + ".md"
	default:
		fmt.Println("govmd: too many command line arguments")
		os.Exit(1)
	}

	// Build YAML header
	yamlHeader += "---\n"
	yamlHeader += "title: \"<++>\"\n"
	yamlHeader += "author: \"<++>\"\n"
	yamlHeader += "date: \"" + now.Format("January 02, 2006") + "\"\n"
	yamlHeader += "---\n\n"
	yamlHeader += "<++>\n"

	// Create file
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("govmd: %s", err)
	}

	defer f.Close()

	// Write file
	nbytes, err := f.WriteString(yamlHeader)
	if nbytes == 0 {
		fmt.Println("govmd: wrote 0 bytes!")
	}
	if err != nil {
		fmt.Printf("govmd: %s", err)
	}
	f.Sync()

	// Open file in Vim
	cmd := exec.Command("vim", fileName)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	vimErr := cmd.Run()
	if vimErr != nil {
		fmt.Printf("govmd: trying to run Vim: %s", vimErr)
	}

}
