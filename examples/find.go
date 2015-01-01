// This example presents a small utility not unlike UNIX 'find' but is always
// anchored to the current working directory and allows partial path matches.
package main

import (
	"fmt"
	"os"
	"strings"

	dtree "github.com/marcinwyszynski/directory_tree"
)

func matchRecursively(node *dtree.Node, query string) {
	if strings.Contains(node.FullPath, query) {
		fmt.Println(node.FullPath)
	}
	for _, child := range node.Children {
		matchRecursively(child, query)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s <query>\n", os.Args[0])
		os.Exit(1)
	}
	dir, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error getting working directory: %v\n", err)
		os.Exit(2)
	}
	root, err := dtree.NewTree(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error traversing the filesystem: %v\n", err)
		os.Exit(3)
	}
	matchRecursively(root, os.Args[1])
	os.Exit(0)
}
