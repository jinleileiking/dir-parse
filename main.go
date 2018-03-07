package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func visit(path string, f os.FileInfo, err error) error {
	if f.IsDir() {
		fmt.Printf("%s\n", path)
	}
	return nil
}

func main() {
	flag.Parse()
	root := flag.Arg(0)
	err := filepath.Walk(root, visit)
	if err != nil {
		fmt.Printf("error\n")
	}
	return
}
