package lib

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func CleanTmpDirectory(dirPath string) error {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		err = os.RemoveAll(fmt.Sprintf("%s/%s", dirPath, file.Name()))
		if err != nil {
			return err
		}
	}

	return nil
}

func ExtractFileNames(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var fileNames []string
	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`\S+\.\w+$`)

	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindString(line)
		if matches != "" {
			fileNames = append(fileNames, matches)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return fileNames, nil
}
