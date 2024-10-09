package common

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func GetFileContents(sqlFilename string) ([]string, error) {
	file, err := os.Open(sqlFilename)
	if err != nil {
		return nil, fmt.Errorf("无法打开文件: %v", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)

	var result []string
	var buffer strings.Builder

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		buffer.WriteString(line)
		buffer.WriteString("\n")

		if strings.HasSuffix(line, ";") {
			result = append(result, buffer.String())
			buffer.Reset()
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("读取文件出错: %v", err)
	}

	return result, nil
}
