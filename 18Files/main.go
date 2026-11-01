/*
Creates a file.
Writes content to a file in append or overwrite mode

Takes arguments like python argparse. checks them case insensitive way
*/
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

const modeAppend string = "append"
const modeOverwrite string = "overwrite"

func main() {
	fmt.Println("Welcome to File management in Go")

	fileContentPtr := flag.String("file-content", "", "This is file content you want to pass")
	filePathPtr := flag.String("file-path", "", "This is the file path")
	modePtr := flag.String("mode", "", "Mode to append or overwrite the content")

	flag.Parse() // arguments are optional by default

	// 2. Dereference pointers into standard string variables
	fileContent := *fileContentPtr
	filePath := *filePathPtr
	mode := *modePtr

	if fileContent == "" {
		fileContent = "New thing arrives APPEND"
	}
	if filePath == "" {
		filePath = "./myText.txt"
	}
	if mode == "" {
		mode = "append"
	}

	createFile(filePath, fileContent, mode)
	readFile(filePath)

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}

func createFile(filePath string, fileContent string, mode ...string) {

	defaultMode := modeAppend
	if len(mode) > 0 {
		defaultMode = mode[0]
	}

	var file *os.File // In Go, any variable created inside an if block block using := only exists inside that block.

	//  compares string insensitively
	if strings.EqualFold(defaultMode, modeOverwrite) {
		// create the file
		file, err := os.Create(filePath)

		checkNilError(err)
		length, err := io.WriteString(file, fileContent) // overwrite

		checkNilError(err)
		fmt.Println("Length is ", length)

	}

	//  compares string insensitively
	if strings.EqualFold(defaultMode, modeAppend) {
		// Open the file in Append-Only mode.
		// O_WRONLY -->  write only
		// If the file doesn't exist, os.O_CREATE will create it.
		// 0644 sets standard read/write permissions.
		file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

		checkNilError(err)
		fileContent = "\n" + fileContent
		_, err = io.WriteString(file, fileContent)

		checkNilError(err)

	} else {
		panic("modes can only be 'append' or 'overwrite'")
	}
	defer file.Close()

}

func readFile(filePath string) {
	contentBytes, err := os.ReadFile(filePath)
	checkNilError(err)
	fmt.Println(contentBytes)
	fmt.Println(string(contentBytes))
}
