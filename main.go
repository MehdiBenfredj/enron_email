package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	emailMap := make(map[string]bool)
	err := filepath.Walk("./inbox", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				fmt.Printf("Error opening file %s: %v\n", path, err)
				return nil
			}
			defer file.Close()

			// Process the file here
			scanner := bufio.NewScanner(file)
			if scanner.Scan() {
				firstLine := scanner.Text()
				if emailMap[firstLine] {
					fmt.Printf("duplicate email id  : %s in %s \n", firstLine, path)
				}
				emailMap[firstLine] = true
			}
			if err := scanner.Err(); err != nil {
				fmt.Printf("Error reading file %s: %s \n", path, err)
			}
			fmt.Print(".")
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error walking directory: %v\n", err)
	}
}

//get first line
