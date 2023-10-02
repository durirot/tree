package main

import (
	"fmt"
	"os"
	"strings"
)

func walkDir(location string, spacing string) error {
	files, err := os.ReadDir(location)
	if err != nil {
		return err
	}
	for _, file := range files {
		fmt.Println(strings.Join([]string{spacing, file.Name()}, ""))
		if file.IsDir() {
			err = walkDir(
				strings.Join([]string{location, file.Name()}, "/"),
				strings.Join([]string{"  ", spacing}, ""))
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func main() {
	fmt.Println("Hello world")
	err := walkDir(".", "")
	if err != nil {
		panic(err)
	}
}
