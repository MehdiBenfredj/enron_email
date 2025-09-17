package processing

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ProcessMail(path string) (from string, to []string, err error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", path, err)
		return from, to, err
	}
	defer file.Close()
	// Process the file here
	scanner := bufio.NewScanner(file)
	line := 1
	to_str := ""
	for scanner.Scan() {
		if line < 3 {
			line++
		} else if line == 3 {
			from = strings.Split(scanner.Text(), " ")[1]
			line++
		} else {
			if strings.Split(scanner.Text(), ":")[0] != "Subject" {
				if scanner.Text()[0] == '\t' {
					to_str += scanner.Text()[1:]
				} else {
					to_str += scanner.Text()
				}
				line++

			} else {
				break
			}
		}
	}
	//extract receivers
	if to_str != "" {
		to = ExctractReceivers(to_str)

	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file %s: %s \n", path, err)
	}
	return from, to, err
}

func ExctractReceivers(to_str string) []string {
	return strings.Split(strings.Split(to_str, ": ")[1], ", ")
}

func Run(maildir string) {
	r := make(map[string]map[string]int) // map k : sender string v : map : k : receiver v : nb of received msgs int
	err := filepath.Walk(maildir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			from, _, err := ProcessMail(path)
			if err != nil {
				fmt.Errorf("Error processing mail %s, %s", path, err)
			}
			_, ok := r[from]
			if !ok {
				//TODO
			}

		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking directory: %v\n", err)
	}
}
