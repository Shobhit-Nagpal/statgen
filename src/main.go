package main

import (
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"statgen/src/block"
	"statgen/src/server"
	"strings"
)

func copyFileContents(file fs.DirEntry, src, dest string) {
	fileContent, err := os.ReadFile(fmt.Sprintf("%s%s", src, file.Name()))
	if err != nil {
		fmt.Fprintf(os.Stderr, fmt.Sprintf("Error reading file %s: %s\n\n", file.Name(), err.Error()))
		os.Exit(1)
	}

	newFileName := fmt.Sprintf("%s%s", dest, file.Name())
	newFile, err := os.Create(newFileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, fmt.Sprintf("Error creating file %s: %s\n\n", newFileName, err.Error()))
		os.Exit(1)
	}

	_, err = newFile.Write(fileContent)
	if err != nil {
		fmt.Fprintf(os.Stderr, fmt.Sprintf("Error writing to file %s: %s\n\n", newFileName, err.Error()))
		os.Exit(1)
	}

}

func copyDirContents(dir fs.DirEntry, src, dest string) {

	err := os.Mkdir(fmt.Sprintf("%s%s/", dest, dir.Name()), 0750)
	if err != nil {
		fmt.Fprintf(os.Stderr, fmt.Sprintf("Error copying directory %s: %s\n\n", dir.Name(), err.Error()))
		os.Exit(1)
	}

	currentDir := fmt.Sprintf("%s%s/", src, dir.Name())

	dirContents, err := os.ReadDir(currentDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, fmt.Sprintf("Error copying directory %s: %s\n\n", currentDir, err.Error()))
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
		fmt.Fprintf(os.Stderr, fmt.Sprintf("Error getting working directory: %s\n\n", err.Error()))
		os.Exit(1)
	}

	publicDir := fmt.Sprintf("%s%s", currentDir, "/public/")
	staticDir := fmt.Sprintf("%s%s", currentDir, "/static/")

	//Remove contents in public directory
	publicDirContents, err := os.ReadDir(publicDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, fmt.Sprintf("Error reading directory %s's contents: %s\n\n", publicDir, err.Error()))
		os.Exit(1)
	}

	if len(publicDirContents) > 0 {
		for _, publicDirContent := range publicDirContents {
			if publicDirContent.IsDir() {
				err = os.RemoveAll(fmt.Sprintf("%s%s", publicDir, publicDirContent.Name()))
				if err != nil {
					fmt.Fprintf(os.Stderr, fmt.Sprintf("Error deleting directory %s's contents: %s\n\n", publicDir, err.Error()))
				}
			} else {
				err = os.Remove(fmt.Sprintf("%s%s", publicDir, publicDirContent.Name()))
				if err != nil {
					fmt.Fprintf(os.Stderr, fmt.Sprintf("Error deleting directory %s's contents: %s\n\n", publicDir, err.Error()))
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

func extractTitle(markdown string) (string, error) {
  headingLine := strings.SplitN(markdown, "\n\n", 2)[0]
  heading := strings.SplitN(headingLine, " ", 2)

  headingNumber := len(heading[0])
  headingText := heading[1]

  if headingNumber != 1 {
    return "", errors.New("Page needs to begin with heading 1")
  }

  return headingText, nil
}

func generatePage(fromPath, templatePath, destPath string) {

  fmt.Fprintf(os.Stdout, "Generarting page from %s to %s using %s\n\n", fromPath, destPath, templatePath)

  template, err := os.ReadFile(templatePath)
  if err != nil {
    fmt.Fprintf(os.Stderr, fmt.Sprintf("Error reading template %s's contents: %s\n\n", templatePath, err.Error()))
  }

  md, err := os.ReadFile(fromPath)
  if err != nil {
    fmt.Fprintf(os.Stderr, fmt.Sprintf("Error reading markdown %s's contents: %s\n\n", fromPath, err.Error()))
  }

  htmlContent, err := block.MarkdownToHTMLNode(string(md))
  if err != nil {
    fmt.Fprintf(os.Stderr, fmt.Sprintf("Error converting markdown to HTML: %s\n\n", err.Error()))
  }

  pageTitle, err := extractTitle(string(md))
  if err != nil {
    fmt.Fprintf(os.Stderr, fmt.Sprintf("Error extracting title from markdown: %s\n\n", err.Error()))
  }

  html := strings.Replace(string(template), "{{ Title }}", pageTitle, 1)
  html = strings.Replace(html, "{{ Content }}", htmlContent, 1)
  
  destFile, err := os.Create(destPath)
  if err != nil {
    fmt.Fprintf(os.Stderr, fmt.Sprintf("Error creating %s: %s\n\n", destPath, err.Error()))
  }

  _, err = destFile.Write([]byte(html))
  if err != nil {
    fmt.Fprintf(os.Stderr, fmt.Sprintf("Error writing to %s: %s\n\n", destPath, err.Error()))
  }
}

func main() {
	dirPtr := flag.String("dir", ".", "Directory to serve files from")
	portPtr := flag.String("port", "8000", "Port to serve HTTP on")
	flag.Parse()

	fmt.Println("Starting up server...")

	copyStatic()
  generatePage("/home/shobhit/repos/statgen/content/index.md", "/home/shobhit/repos/statgen/template/template.html", "/home/shobhit/repos/statgen/public/index.html")
	server.Start(*dirPtr, *portPtr)
}
