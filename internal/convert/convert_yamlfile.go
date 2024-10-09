package convert

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"strings"
)

type YamlFileData struct {
	FileName   string
	ResultFile string
}

func extendPrefix(prefix, key string) string {
	if prefix != "" {
		prefix += "_"
	}
	return prefix + key
}

func walkAndWrite(data interface{}, prefix string, file *os.File) {
	switch node := data.(type) {
	case map[string]interface{}:
		for key, value := range node {
			newPrefix := extendPrefix(prefix, key)
			walkAndWrite(value, newPrefix, file)
		}
	default:
		finalKey := strings.TrimRight(prefix, "_")
		line := fmt.Sprintf("export %s=%v\n", finalKey, node)
		_, err := file.WriteString(line)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			os.Exit(1)
		}
	}
}

func (d YamlFileData) Run() {
	if d.FileName == "" || d.ResultFile == "" {
		log.Println("入参异常")
		return
	}

	data, err := os.ReadFile(d.FileName)
	if err != nil {
		log.Fatalf("error reading YAML file: %v", err)
		return
	}

	var config interface{}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		fmt.Printf("Error parsing YAML file: %s\n", err)
		return
	}

	file, err := os.Create(d.ResultFile)
	if err != nil {
		fmt.Printf("Error creating file %s: %s\n", d.ResultFile, err)
		os.Exit(1)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	walkAndWrite(config, "", file)

}
