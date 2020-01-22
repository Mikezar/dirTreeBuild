package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Node struct {
	Name        string
	IsDirectory bool
	Size        int64
	Childs      []*Node
}

const (
	LineStart = "├"
	NewLine   = "│"
	LineEnd   = "└"
	Line      = "───"
)

func (parent *Node) buildTree(path string, includeFiles bool) error {
	var files, err = ioutil.ReadDir(path)

	if err != nil {
		return err
	}

	for _, f := range files {
		if !f.IsDir() && includeFiles || f.IsDir() {

			var childNode = Node{
				Name:        f.Name(),
				IsDirectory: f.IsDir(),
			}

			if !childNode.IsDirectory {
				childNode.Size = f.Size()
			}

			parent.Childs = append(parent.Childs, &childNode)

			if childNode.IsDirectory {
				childNode.buildTree(filepath.Join(path, childNode.Name), includeFiles)
			}
		}
	}

	return nil
}

func (parent *Node) printTree(output io.Writer, prefix string) {

	var count = len(parent.Childs) - 1

	for i, n := range parent.Childs {

		var line, newPrefix = n.buildLine(i == count, prefix)

		io.WriteString(output, line)

		if n.IsDirectory {
			n.printTree(output, newPrefix)
		}
	}
}

func (node *Node) buildLine(isLastNode bool, prefix string) (line string, newPrefix string) {

	var firstChar string
	var size string

	if isLastNode {
		firstChar = LineEnd

		if node.IsDirectory {
			newPrefix = prefix + "\t"
		}
	} else {
		firstChar = LineStart
		newPrefix = prefix + NewLine + "\t"
	}

	if !node.IsDirectory {
		size = " (empty)"

		if node.Size > 0 {
			size = fmt.Sprintf(" (%vb)", node.Size)
		}
	}

	line = prefix + firstChar + Line + node.Name + size + "\n"

	return
}

func dirTree(out io.Writer, path string, includeFiles bool) error {
	var mainNode = Node{
		Name:        path,
		IsDirectory: true,
	}

	err := mainNode.buildTree(path, includeFiles)

	if err != nil {
		return err
	}

	mainNode.printTree(out, "")

	return nil
}

func main() {
	out := os.Stdout

	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("Pass path and/or [-f] argurment")
	}

	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, os.Args[1], printFiles)
	if err != nil {
		panic(err.Error())
	}
}
