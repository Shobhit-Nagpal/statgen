package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"statgen/src/server"
)

func copyFileContents(file fs.DirEntry, src, dest string) {
	fileContent, err := os.ReadFile(fmt.Sprintf("%s%s", src, file.Name()))
	if err != nil {
		fmt.Fprintf(os.Stderr, fmt.Sprintf("Error reading file %s: %s", file.Name(), err.Error()))
		os.Exit(1)
	}

	newFileName := fmt.Sprintf("%s%s", dest, file.Name())
	newFile, err := os.Create(newFileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, fmt.Sprintf("Error creating file %s: %s", newFileName, err.Error()))
		os.Exit(1)
	}

	_, err = newFile.Write(fileContent)
	if err != nil {
		fmt.Fprintf(os.Stderr, fmt.Sprintf("Error writing to file %s: %s", newFileName, err.Error()))
		os.Exit(1)
	}

}

func copyDirContents(dir fs.DirEntry, src, dest string) {

	err := os.Mkdir(fmt.Sprintf("%s%s/", dest, dir.Name()), 0750)
	if err != nil {
		fmt.Fprintf(os.Stderr, fmt.Sprintf("Error copying directory %s: %s", dir.Name(), err.Error()))
		os.Exit(1)
	}

	currentDir := fmt.Sprintf("%s%s/", src, dir.Name())

	dirContents, err := os.ReadDir(currentDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, fmt.Sprintf("Error copying directory %s: %s", currentDir, err.Error()))
		os.Exit(1)
	}

	for _, dirContent := range dirContents {
		if dirContent.IsDir() {
			copyDirContents(dirContent, currentDir, fmt.Sprintf("%s%s/", dest, dir.Name()))
		} else {
			copyFileContents(dirContent, currentDir, fmt.Sprintf("%s%s/", dest, dir.Name()))
		}
	}
}

func copyStatic() {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, fmt.Sprintf("Error getting working directory: %s", err.Error()))
		os.Exit(1)
	}

	publicDir := fmt.Sprintf("%s%s", currentDir, "/public/")
	staticDir := fmt.Sprintf("%s%s", currentDir, "/static/")

	//Remove contents in public directory
	publicDirContents, err := os.ReadDir(publicDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, fmt.Sprintf("Error reading directory %s's contents: %s", publicDir, err.Error()))
		os.Exit(1)
	}

	if len(publicDirContents) > 0 {
		for _, publicDirContent := range publicDirContents {
			if publicDirContent.IsDir() {
				err = os.RemoveAll(fmt.Sprintf("%s%s", publicDir, publicDirContent.Name()))
				if err != nil {
					fmt.Fprintf(os.Stderr, fmt.Sprintf("Error deleting directory %s's contents: %s", publicDir, err.Error()))
				}
			} else {
				err = os.Remove(fmt.Sprintf("%s%s", publicDir, publicDirContent.Name()))
				if err != nil {
					fmt.Fprintf(os.Stderr, fmt.Sprintf("Error deleting directory %s's contents: %s", publicDir, err.Error()))
				}
			}

		}
	}

	//Copy static dir contents

	dirContents, err := os.ReadDir(staticDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}

	for _, dirContent := range dirContents {
		if dirContent.IsDir() {
			copyDirContents(dirContent, staticDir, publicDir)
		} else {
			copyFileContents(dirContent, staticDir, publicDir)
		}
	}

}

func main() {
	dirPtr := flag.String("dir", ".", "Directory to serve files from")
	portPtr := flag.String("port", "8000", "Port to serve HTTP on")
	flag.Parse()

	fmt.Println("Starting up server...")

	copyStatic()
	server.Start(*dirPtr, *portPtr)
}
