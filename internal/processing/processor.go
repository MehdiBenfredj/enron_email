package processing

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type Result struct {
	from string
	to   []string
	err  error
}

func ProcessMail(path string) (r Result) {
	file, err := os.Open(path)
	if err != nil {
		r.err = fmt.Errorf("error opening file %s: %v", path, err)
		return r
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
			r.from = strings.Split(scanner.Text(), " ")[1]
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
		r.to = ExctractReceivers(to_str)
	}
	if err := scanner.Err(); err != nil {
		r.err = fmt.Errorf("error reading file %s: %s", path, err)
	}
	return r
}

func ExctractReceivers(to_str string) []string {
	return strings.Split(strings.Split(to_str, ": ")[1], ", ")
}

func Run(maildir, resultPath string) (map[string]map[string]int, error) {
	resultMap := make(map[string]map[string]int) // map k : sender string v : map : k : receiver v : nb of received msgs int
	err := filepath.WalkDir(maildir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			res := ProcessMail(path)
			if res.err != nil {
				return fmt.Errorf("error processing mail %s, %s", path, res.err)
			}
			if len(res.to) == 0 {
				return nil
			}
			if _, ok := resultMap[res.from]; !ok {
				resultMap[res.from] = make(map[string]int, 0)
			}
			FillReceivers(resultMap[res.from], res.to)
		} else {
			fmt.Printf("processing %s\n", d.Name())
		}

		err = WriteResults(resultMap, resultPath)

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("error walking directory: %s", err)
	}
	return resultMap, nil
}

func WriteResults(resultMap map[string]map[string]int, resultPath string) error {
	//write to file result path
	file, err := os.Create(resultPath + "/result.txt")
	if err != nil {
		return fmt.Errorf("error creating result file: %s", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for result := range resultMap {
		_, err := writer.WriteString(fmt.Sprintf("%s : %v\n", result, resultMap[result]))
		if err != nil {
			return fmt.Errorf("error writing to result file: %s", err)
		}
	}
	return nil
}

func FillReceivers(receiversMap map[string]int, receiversList []string) {
	for _, receiver := range receiversList {
		if _, ok := receiversMap[receiver]; !ok {
			receiversMap[receiver] = 1
		} else {
			receiversMap[receiver] = receiversMap[receiver] + 1
		}
	}
}

func RunParallel(maildir string, resultPath string) (map[string]map[string]int, error) {
	resultMap := make(map[string]map[string]int)
	notificationch := make(chan Result)
	wg := sync.WaitGroup{}
	err := filepath.WalkDir(maildir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			wg.Add(1)
			go func() error {
				res := ProcessMail(path)
				if res.err != nil {
					return fmt.Errorf("error processing mail %s, %s", path, res.err)
				}
				notificationch <- res
				wg.Done()
				return nil
			}()
		} else {
			fmt.Printf("processing %s\n", d.Name())
		}
		return nil
	})
	go func() {
		wg.Wait()
		close(notificationch)
	}()
	for notification := range notificationch {
		if _, ok := resultMap[notification.from]; !ok && len(notification.to) != 0 {
			resultMap[notification.from] = make(map[string]int, 0)
		}
		FillReceivers(resultMap[notification.from], notification.to)
	}

	if err != nil {
		return nil, fmt.Errorf("error walking directory: %s", err)
	}

	err = WriteResults(resultMap, resultPath)
	if err != nil {
		return nil, err
	}

	return resultMap, nil
}
