package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	// "runtime/debug"

	ignore "github.com/sabhiram/go-gitignore"
)

var maxDepth int
var showHidden bool
var directoriesOnly bool
var showVersion bool
var version string

func main() {
	flag.IntVar(&maxDepth, "L", 0, "Limit the depth of the tree (0 means no limit).")
	flag.BoolVar(&showHidden, "a", false, "Show all files, including hidden ones (starting with a dot).")
	flag.BoolVar(&directoriesOnly, "d", false, "Display directories only (excluding files).")
	flag.BoolVar(&showVersion, "v", false, "Print the current version.")

	flag.Parse()

	root := "."

	positionalArgs := flag.Args()
	if len(positionalArgs) > 0 {
		root = positionalArgs[0]
	}

	var ignoreObject *ignore.GitIgnore

	gitIgnorePath := filepath.Join(root, ".gitignore")
	if _, err := os.Stat(gitIgnorePath); err == nil {
		object, err := ignore.CompileIgnoreFile(gitIgnorePath)
		if err != nil {
			fmt.Printf("No se pudo leer .gitignore: %v\n", err)
		} else {
			ignoreObject = object
		}
	}

	if showVersion {
		fmt.Println("version: ", version)
		return
	}

	fmt.Println(root)
	printTree(root, root, "", ignoreObject, 0)
}

// printTree recorre los directorios recursivamente
func printTree(
	basePath string,
	currentPath string,
	prefix string,
	ignoreObj *ignore.GitIgnore,
	currentDepth int) {
	if maxDepth > 0 && currentDepth >= maxDepth {
		return
	}

	entries, err := os.ReadDir(currentPath)
	if err != nil {
		return
	}

	var validEntries []os.DirEntry

	for _, entry := range entries {
		if directoriesOnly && !entry.IsDir() {
			continue
		}

		if entry.Name() == ".git" {
			continue
		}

		if !showHidden && strings.HasPrefix(entry.Name(), ".") {
			continue
		}

		fullPath := filepath.Join(currentPath, entry.Name())
		relPath, _ := filepath.Rel(basePath, fullPath)
		if ignoreObj != nil && ignoreObj.MatchesPath(relPath) {
			continue
		}

		validEntries = append(validEntries, entry)
	}

	for i, entry := range validEntries {
		isLast := i == len(validEntries)-1
		connector := "├── "
		if isLast {
			connector = "└── "
		}

		fmt.Printf("%s%s%s\n", prefix, connector, entry.Name())

		if entry.IsDir() {
			newPrefix := prefix
			if isLast {
				newPrefix += "    "
			} else {
				newPrefix += "│   "
			}

			subPath := filepath.Join(currentPath, entry.Name())

			printTree(basePath, subPath, newPrefix, ignoreObj, currentDepth+1)
		}
	}
}
