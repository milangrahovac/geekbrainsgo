package main

import (
	"flag"
	"fmt"

	"github.com/milangrahovac/geekbrainsgo/homework9/pkg/confreader"
)

func main() {
	fmt.Println("homework 9 !!!")

	var filePath string
	flag.StringVar(&filePath, "f", "", "Please add file name.")
	flag.Parse()
	file := confreader.File{Path: filePath}

	if len(filePath) == 0 {
		fmt.Println("Error: You must add a file with flag f")
		fmt.Println("Please add a file with .json or .yaml extension.")
	} else if !file.IsFile() {
		fmt.Printf("Error: Cannot find file %s\n", filePath)
		fmt.Println("Please add a file with .json or .yaml extension.")

	} else {
		data := *file.OpenFile()
		fmt.Println("data from file:", filePath)
		fmt.Println("data:", data)
		fmt.Println()
		fmt.Println("School:", data.Name)
		fmt.Println("Website:", data.Website)
		fmt.Println("Course:", data.Course)
		fmt.Println("Keywords:", data.Keywords)

	}

}
