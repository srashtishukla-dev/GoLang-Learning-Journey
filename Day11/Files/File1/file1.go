package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open(`Files\File1\example.txt`)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fileInfo, err := f.Stat()
	if err != nil {
		panic(err)
	}

	fmt.Println("File name:", fileInfo.Name())
	fmt.Println("File or folder:", fileInfo.IsDir())
	fmt.Println("file size:", fileInfo.Size())
	fmt.Println("file permission:", fileInfo.Mode())
}
